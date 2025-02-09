package wallet

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/fireblocks"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var _ Wallet = (*fireblocksWallet)(nil)

var (
	// ErrNotYetBroadcasted indicates that the transaction has not been broadcasted yet.
	// This can happen if the transaction is still being processed by Fireblocks and has not been broadcasted to the
	// blockchain yet.
	ErrNotYetBroadcasted = errors.New("transaction not yet broadcasted")
	// ErrReceiptNotYetAvailable indicates that the transaction has been broadcasted but has not been confirmed onchain
	// yet.
	ErrReceiptNotYetAvailable = errors.New("transaction receipt not yet available")
	ErrTransactionFailed      = errors.New("transaction failed")
)

type ethClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	bind.ContractBackend
}

type fireblocksWallet struct {
	// mu protects access to nonceToTxID and txIDToNonce which can be
	// accessed concurrently by SendTransaction and GetTransactionReceipt
	mu sync.Mutex

	fireblocksClient fireblocks.Client
	ethClient        ethClient
	vaultAccountName string
	logger           logging.Logger
	chainID          *big.Int

	// nonceToTx keeps track of the transaction ID for each nonce
	// this is used to retrieve the transaction hash for a given nonce
	// when a replacement transaction is submitted.
	nonceToTxID map[uint64]TxID
	txIDToNonce map[TxID]uint64

	// caches
	account              *fireblocks.VaultAccount
	whitelistedContracts map[common.Address]*fireblocks.WhitelistedContract
	whitelistedAccounts  map[common.Address]*fireblocks.WhitelistedAccount
}

func NewFireblocksWallet(
	fireblocksClient fireblocks.Client,
	ethClient ethClient,
	vaultAccountName string,
	logger logging.Logger,
) (Wallet, error) {
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID: %w", err)
	}
	logger.Debug("Creating new Fireblocks wallet for chain", "chainID", chainID)
	return &fireblocksWallet{
		fireblocksClient: fireblocksClient,
		ethClient:        ethClient,
		vaultAccountName: vaultAccountName,
		logger:           logger,
		chainID:          chainID,

		nonceToTxID: make(map[uint64]TxID),
		txIDToNonce: make(map[TxID]uint64),

		// caches
		account:              nil,
		whitelistedContracts: make(map[common.Address]*fireblocks.WhitelistedContract),
		whitelistedAccounts:  make(map[common.Address]*fireblocks.WhitelistedAccount),
	}, nil
}

func (t *fireblocksWallet) getAccount(ctx context.Context) (*fireblocks.VaultAccount, error) {
	if t.account == nil {
		accounts, err := t.fireblocksClient.ListVaultAccounts(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing vault accounts: %w", err)
		}
		for i, a := range accounts {
			if a.Name == t.vaultAccountName {
				t.account = &accounts[i]
				break
			}
		}
	}
	return t.account, nil
}

func (f *fireblocksWallet) getWhitelistedAccount(
	ctx context.Context,
	address common.Address,
) (*fireblocks.WhitelistedAccount, error) {
	assetID, ok := fireblocks.AssetIDByChain[f.chainID.Uint64()]
	if !ok {
		return nil, fmt.Errorf("unsupported chain %d", f.chainID.Uint64())
	}
	whitelistedAccount, ok := f.whitelistedAccounts[address]
	if !ok {
		accounts, err := f.fireblocksClient.ListExternalWallets(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing external wallets: %w", err)
		}
		for i, a := range accounts {
			for _, asset := range a.Assets {
				if asset.Address == address && asset.Status == "APPROVED" && asset.ID == assetID {
					f.whitelistedAccounts[address] = &accounts[i]
					whitelistedAccount = &accounts[i]
					return whitelistedAccount, nil
				}
			}
		}
	}

	if whitelistedAccount == nil {
		return nil, fmt.Errorf("account %s not found in whitelisted accounts", address.Hex())
	}
	return whitelistedAccount, nil
}

func (t *fireblocksWallet) getWhitelistedContract(
	ctx context.Context,
	address common.Address,
) (*fireblocks.WhitelistedContract, error) {
	assetID, ok := fireblocks.AssetIDByChain[t.chainID.Uint64()]
	if !ok {
		return nil, fmt.Errorf("unsupported chain %d", t.chainID.Uint64())
	}
	contract, ok := t.whitelistedContracts[address]
	if !ok {
		contracts, err := t.fireblocksClient.ListContracts(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing contracts: %w", err)
		}
		for i_c, c := range contracts {
			for _, a := range c.Assets {
				if a.Address == address && a.Status == "APPROVED" && a.ID == assetID {
					t.whitelistedContracts[address] = &contracts[i_c]
					contract = &contracts[i_c]
					return contract, nil
				}
			}
		}
	}

	if contract == nil {
		return nil, fmt.Errorf("contract %s not found in whitelisted contracts", address.Hex())
	}
	return contract, nil
}

func (t *fireblocksWallet) SendTransaction(ctx context.Context, tx *types.Transaction) (TxID, error) {
	assetID, ok := fireblocks.AssetIDByChain[t.chainID.Uint64()]
	if !ok {
		return "", fmt.Errorf("unsupported chain %d", t.chainID.Uint64())
	}
	account, err := t.getAccount(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting account: %w", err)
	}
	foundAsset := false
	for _, a := range account.Assets {
		if a.ID == assetID {
			if a.Available == "0" {
				return "", errors.New("insufficient funds")
			}
			foundAsset = true
			break
		}
	}
	if !foundAsset {
		return "", fmt.Errorf("asset %s not found in account %s", assetID, t.vaultAccountName)
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	// if the nonce is already in the map, it means that the transaction was already submitted
	// we need to get the replacement transaction hash and use it as the replaceTxByHash parameter
	replaceTxByHash := ""
	nonce := tx.Nonce()
	if txID, ok := t.nonceToTxID[nonce]; ok {
		fireblockTx, err := t.fireblocksClient.GetTransaction(ctx, txID)
		if err != nil {
			return "", fmt.Errorf("error getting fireblocks transaction %s: %w", txID, err)
		}
		if fireblockTx.TxHash != "" {
			replaceTxByHash = fireblockTx.TxHash
		}
	}

	gasLimit := ""
	if tx.Gas() > 0 {
		gasLimit = strconv.FormatUint(tx.Gas(), 10)
	}

	// if the gas fees are specified in the transaction, use them.
	// Otherwise, use the default "HIGH" gas price estimated by Fireblocks
	maxFee := ""
	priorityFee := ""
	gasPrice := ""
	feeLevel := fireblocks.FeeLevel("")
	if tx.GasFeeCap().Cmp(big.NewInt(0)) > 0 && tx.GasTipCap().Cmp(big.NewInt(0)) > 0 {
		maxFee = weiToGwei(tx.GasFeeCap()).String()
		priorityFee = weiToGwei(tx.GasTipCap()).String()
	} else if tx.GasPrice().Cmp(big.NewInt(0)) > 0 {
		gasPrice = weiToGwei(tx.GasPrice()).String()
	} else {
		feeLevel = fireblocks.FeeLevelHigh
	}

	var res *fireblocks.TransactionResponse
	if len(tx.Data()) == 0 && tx.Value().Cmp(big.NewInt(0)) > 0 {
		targetAccount, clientErr := t.getWhitelistedAccount(ctx, *tx.To())
		if clientErr != nil {
			return "", fmt.Errorf("error getting whitelisted account %s: %w", tx.To().Hex(), clientErr)
		}
		req := fireblocks.NewTransferRequest(
			"", // externalTxID
			assetID,
			account.ID,                      // source account ID
			targetAccount.ID,                // destination account ID
			weiToEther(tx.Value()).String(), // amount in ETH
			replaceTxByHash,                 // replaceTxByHash
			gasPrice,
			gasLimit,
			maxFee,
			priorityFee,
			feeLevel,
		)
		res, err = t.fireblocksClient.Transfer(ctx, req)
	} else if len(tx.Data()) > 0 {
		contract, clientErr := t.getWhitelistedContract(ctx, *tx.To())
		if clientErr != nil {
			return "", fmt.Errorf("error getting whitelisted contract %s: %w", tx.To().Hex(), clientErr)
		}
		req := fireblocks.NewContractCallRequest(
			"", // externalTxID
			assetID,
			account.ID,                      // source account ID
			contract.ID,                     // destination account ID
			weiToEther(tx.Value()).String(), // amount
			hexutil.Encode(tx.Data()),       // calldata
			replaceTxByHash,                 // replaceTxByHash
			gasPrice,
			gasLimit,
			maxFee,
			priorityFee,
			feeLevel,
		)
		res, err = t.fireblocksClient.ContractCall(ctx, req)
	} else {
		return "", errors.New("transaction has no value and no data")
	}

	if err != nil {
		return "", fmt.Errorf("error sending a transaction %s: %w", tx.To().Hex(), err)
	}
	t.nonceToTxID[nonce] = res.ID
	t.txIDToNonce[res.ID] = nonce
	t.logger.Debug("Fireblocks contract call complete", "txID", res.ID, "status", res.Status)

	return res.ID, nil
}

func (t *fireblocksWallet) CancelTransactionBroadcast(ctx context.Context, txID TxID) (bool, error) {
	return t.fireblocksClient.CancelTransaction(ctx, string(txID))
}

func (t *fireblocksWallet) GetTransactionReceipt(ctx context.Context, txID TxID) (*types.Receipt, error) {
	fireblockTx, err := t.fireblocksClient.GetTransaction(ctx, txID)
	if err != nil {
		return nil, fmt.Errorf("error getting fireblocks transaction %s: %w", txID, err)
	}
	if fireblockTx.Status == fireblocks.Completed {
		txHash := common.HexToHash(fireblockTx.TxHash)
		receipt, err := t.ethClient.TransactionReceipt(ctx, txHash)
		if err == nil {
			t.mu.Lock()
			defer t.mu.Unlock()
			if nonce, ok := t.txIDToNonce[txID]; ok {
				delete(t.nonceToTxID, nonce)
				delete(t.txIDToNonce, txID)
			}

			return receipt, nil
		}
		if errors.Is(err, ethereum.NotFound) {
			return nil, fmt.Errorf("%w: for txID %s", ErrReceiptNotYetAvailable, txID)
		} else {
			return nil, fmt.Errorf("Transaction receipt retrieval failed: %w", err)
		}
	} else if fireblockTx.Status == fireblocks.Failed ||
		fireblockTx.Status == fireblocks.Rejected ||
		fireblockTx.Status == fireblocks.Cancelled ||
		fireblockTx.Status == fireblocks.Blocked {
		return nil, fmt.Errorf("%w: the Fireblocks transaction %s has been %s", ErrTransactionFailed, txID, fireblockTx.Status)
	} else if fireblockTx.Status == fireblocks.Submitted ||
		fireblockTx.Status == fireblocks.PendingScreening ||
		fireblockTx.Status == fireblocks.PendingAuthorization ||
		fireblockTx.Status == fireblocks.Queued ||
		fireblockTx.Status == fireblocks.PendingSignature ||
		fireblockTx.Status == fireblocks.PendingEmailApproval ||
		fireblockTx.Status == fireblocks.Pending3rdParty ||
		fireblockTx.Status == fireblocks.Broadcasting {
		return nil, fmt.Errorf("%w: the Fireblocks transaction %s is in status %s", ErrNotYetBroadcasted, txID, fireblockTx.Status)
	}

	return nil, fmt.Errorf(
		"%w: the Fireblocks transaction %s is in status %s",
		ErrReceiptNotYetAvailable,
		txID,
		fireblockTx.Status,
	)
}

func (f *fireblocksWallet) SenderAddress(ctx context.Context) (common.Address, error) {
	account, err := f.getAccount(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("error getting account: %w", err)
	}
	addresses, err := f.fireblocksClient.GetAssetAddresses(
		ctx,
		account.ID,
		fireblocks.AssetIDByChain[f.chainID.Uint64()],
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("error getting asset addresses: %w", err)
	}
	if len(addresses) == 0 {
		return common.Address{}, errors.New("no addresses found")
	}
	return common.HexToAddress(addresses[0].Address), nil
}

func weiToGwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.GWei))
}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
