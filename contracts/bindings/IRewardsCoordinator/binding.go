// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIRewardsCoordinator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRewardsCoordinatorDistributionRoot is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorDistributionRoot struct {
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Disabled                       bool
}

// IRewardsCoordinatorEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IRewardsCoordinatorOperatorDirectedRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorOperatorDirectedRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	OperatorRewards          []IRewardsCoordinatorOperatorReward
	StartTimestamp           uint32
	Duration                 uint32
	Description              string
}

// IRewardsCoordinatorOperatorReward is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorOperatorReward struct {
	Operator common.Address
	Amount   *big.Int
}

// IRewardsCoordinatorRewardsMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IRewardsCoordinatorEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IRewardsCoordinatorTokenTreeMerkleLeaf
}

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IRewardsCoordinatorTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// ContractIRewardsCoordinatorMetaData contains all meta data concerning the ContractIRewardsCoordinator contract.
var ContractIRewardsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmission\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultOperatorSplitBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsMerkleClaim[]\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultOperatorSplit\",\"inputs\":[{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultOperatorSplitBipsSet\",\"inputs\":[{\"name\":\"oldDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedAVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPISplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ContractIRewardsCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIRewardsCoordinatorMetaData.ABI instead.
var ContractIRewardsCoordinatorABI = ContractIRewardsCoordinatorMetaData.ABI

// ContractIRewardsCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractIRewardsCoordinatorMethods interface {
	ContractIRewardsCoordinatorCalls
	ContractIRewardsCoordinatorTransacts
	ContractIRewardsCoordinatorFilters
}

// ContractIRewardsCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIRewardsCoordinatorCalls interface {
	CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error)

	GENESISREWARDSTIMESTAMP(opts *bind.CallOpts) (uint32, error)

	MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error)

	MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error)

	MAXREWARDSDURATION(opts *bind.CallOpts) (uint32, error)

	ActivationDelay(opts *bind.CallOpts) (uint32, error)

	CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error)

	CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error)

	CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error)

	ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error)

	CumulativeClaimed(opts *bind.CallOpts, claimer common.Address, token common.Address) (*big.Int, error)

	CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error)

	DefaultOperatorSplitBips(opts *bind.CallOpts) (uint16, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error)

	GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error)

	GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorDistributionRoot, error)

	GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error)

	GetOperatorAVSSplit(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error)

	GetOperatorPISplit(opts *bind.CallOpts, operator common.Address) (uint16, error)

	GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error)

	RewardsUpdater(opts *bind.CallOpts) (common.Address, error)
}

// ContractIRewardsCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIRewardsCoordinatorTransacts interface {
	CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error)

	CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error)

	CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error)

	CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmission []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error)

	DisableRoot(opts *bind.TransactOpts, rootIndex uint32) (*types.Transaction, error)

	ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error)

	ProcessClaims(opts *bind.TransactOpts, claims []IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error)

	SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error)

	SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error)

	SetDefaultOperatorSplit(opts *bind.TransactOpts, split uint16) (*types.Transaction, error)

	SetOperatorAVSSplit(opts *bind.TransactOpts, operator common.Address, avs common.Address, split uint16) (*types.Transaction, error)

	SetOperatorPISplit(opts *bind.TransactOpts, operator common.Address, split uint16) (*types.Transaction, error)

	SetRewardsForAllSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error)

	SetRewardsUpdater(opts *bind.TransactOpts, _rewardsUpdater common.Address) (*types.Transaction, error)

	SubmitRoot(opts *bind.TransactOpts, root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error)
}

// ContractIRewardsCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIRewardsCoordinatorFilters interface {
	FilterAVSRewardsSubmissionCreated(opts *bind.FilterOpts, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator, error)
	WatchAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorAVSRewardsSubmissionCreated, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseAVSRewardsSubmissionCreated(log types.Log) (*ContractIRewardsCoordinatorAVSRewardsSubmissionCreated, error)

	FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractIRewardsCoordinatorActivationDelaySetIterator, error)
	WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorActivationDelaySet) (event.Subscription, error)
	ParseActivationDelaySet(log types.Log) (*ContractIRewardsCoordinatorActivationDelaySet, error)

	FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*ContractIRewardsCoordinatorClaimerForSetIterator, error)
	WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error)
	ParseClaimerForSet(log types.Log) (*ContractIRewardsCoordinatorClaimerForSet, error)

	FilterDefaultOperatorSplitBipsSet(opts *bind.FilterOpts) (*ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator, error)
	WatchDefaultOperatorSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet) (event.Subscription, error)
	ParseDefaultOperatorSplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet, error)

	FilterDistributionRootDisabled(opts *bind.FilterOpts, rootIndex []uint32) (*ContractIRewardsCoordinatorDistributionRootDisabledIterator, error)
	WatchDistributionRootDisabled(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDistributionRootDisabled, rootIndex []uint32) (event.Subscription, error)
	ParseDistributionRootDisabled(log types.Log) (*ContractIRewardsCoordinatorDistributionRootDisabled, error)

	FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (*ContractIRewardsCoordinatorDistributionRootSubmittedIterator, error)
	WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (event.Subscription, error)
	ParseDistributionRootSubmitted(log types.Log) (*ContractIRewardsCoordinatorDistributionRootSubmitted, error)

	FilterOperatorAVSSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address, avs []common.Address) (*ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator, error)
	WatchOperatorAVSSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorAVSSplitBipsSet, caller []common.Address, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSSplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorOperatorAVSSplitBipsSet, error)

	FilterOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator, error)
	WatchOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseOperatorDirectedAVSRewardsSubmissionCreated(log types.Log) (*ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, error)

	FilterOperatorPISplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator, error)
	WatchOperatorPISplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorPISplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error)
	ParseOperatorPISplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorOperatorPISplitBipsSet, error)

	FilterRewardsClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*ContractIRewardsCoordinatorRewardsClaimedIterator, error)
	WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error)
	ParseRewardsClaimed(log types.Log) (*ContractIRewardsCoordinatorRewardsClaimed, error)

	FilterRewardsForAllSubmitterSet(opts *bind.FilterOpts, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (*ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator, error)
	WatchRewardsForAllSubmitterSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsForAllSubmitterSet, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error)
	ParseRewardsForAllSubmitterSet(log types.Log) (*ContractIRewardsCoordinatorRewardsForAllSubmitterSet, error)

	FilterRewardsSubmissionForAllCreated(opts *bind.FilterOpts, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator, error)
	WatchRewardsSubmissionForAllCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsSubmissionForAllCreated, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseRewardsSubmissionForAllCreated(log types.Log) (*ContractIRewardsCoordinatorRewardsSubmissionForAllCreated, error)

	FilterRewardsSubmissionForAllEarnersCreated(opts *bind.FilterOpts, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator, error)
	WatchRewardsSubmissionForAllEarnersCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseRewardsSubmissionForAllEarnersCreated(log types.Log) (*ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, error)

	FilterRewardsUpdaterSet(opts *bind.FilterOpts, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (*ContractIRewardsCoordinatorRewardsUpdaterSetIterator, error)
	WatchRewardsUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsUpdaterSet, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (event.Subscription, error)
	ParseRewardsUpdaterSet(log types.Log) (*ContractIRewardsCoordinatorRewardsUpdaterSet, error)
}

// ContractIRewardsCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractIRewardsCoordinator struct {
	ContractIRewardsCoordinatorCaller     // Read-only binding to the contract
	ContractIRewardsCoordinatorTransactor // Write-only binding to the contract
	ContractIRewardsCoordinatorFilterer   // Log filterer for contract events
}

// ContractIRewardsCoordinator implements the ContractIRewardsCoordinatorMethods interface.
var _ ContractIRewardsCoordinatorMethods = (*ContractIRewardsCoordinator)(nil)

// ContractIRewardsCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIRewardsCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRewardsCoordinatorCaller implements the ContractIRewardsCoordinatorCalls interface.
var _ ContractIRewardsCoordinatorCalls = (*ContractIRewardsCoordinatorCaller)(nil)

// ContractIRewardsCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIRewardsCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRewardsCoordinatorTransactor implements the ContractIRewardsCoordinatorTransacts interface.
var _ ContractIRewardsCoordinatorTransacts = (*ContractIRewardsCoordinatorTransactor)(nil)

// ContractIRewardsCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIRewardsCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRewardsCoordinatorFilterer implements the ContractIRewardsCoordinatorFilters interface.
var _ ContractIRewardsCoordinatorFilters = (*ContractIRewardsCoordinatorFilterer)(nil)

// ContractIRewardsCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIRewardsCoordinatorSession struct {
	Contract     *ContractIRewardsCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIRewardsCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIRewardsCoordinatorCallerSession struct {
	Contract *ContractIRewardsCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractIRewardsCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIRewardsCoordinatorTransactorSession struct {
	Contract     *ContractIRewardsCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIRewardsCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIRewardsCoordinatorRaw struct {
	Contract *ContractIRewardsCoordinator // Generic contract binding to access the raw methods on
}

// ContractIRewardsCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIRewardsCoordinatorCallerRaw struct {
	Contract *ContractIRewardsCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIRewardsCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIRewardsCoordinatorTransactorRaw struct {
	Contract *ContractIRewardsCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIRewardsCoordinator creates a new instance of ContractIRewardsCoordinator, bound to a specific deployed contract.
func NewContractIRewardsCoordinator(address common.Address, backend bind.ContractBackend) (*ContractIRewardsCoordinator, error) {
	contract, err := bindContractIRewardsCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinator{ContractIRewardsCoordinatorCaller: ContractIRewardsCoordinatorCaller{contract: contract}, ContractIRewardsCoordinatorTransactor: ContractIRewardsCoordinatorTransactor{contract: contract}, ContractIRewardsCoordinatorFilterer: ContractIRewardsCoordinatorFilterer{contract: contract}}, nil
}

// NewContractIRewardsCoordinatorCaller creates a new read-only instance of ContractIRewardsCoordinator, bound to a specific deployed contract.
func NewContractIRewardsCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractIRewardsCoordinatorCaller, error) {
	contract, err := bindContractIRewardsCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorCaller{contract: contract}, nil
}

// NewContractIRewardsCoordinatorTransactor creates a new write-only instance of ContractIRewardsCoordinator, bound to a specific deployed contract.
func NewContractIRewardsCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIRewardsCoordinatorTransactor, error) {
	contract, err := bindContractIRewardsCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorTransactor{contract: contract}, nil
}

// NewContractIRewardsCoordinatorFilterer creates a new log filterer instance of ContractIRewardsCoordinator, bound to a specific deployed contract.
func NewContractIRewardsCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIRewardsCoordinatorFilterer, error) {
	contract, err := bindContractIRewardsCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorFilterer{contract: contract}, nil
}

// bindContractIRewardsCoordinator binds a generic wrapper to an already deployed contract.
func bindContractIRewardsCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIRewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIRewardsCoordinator.Contract.ContractIRewardsCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ContractIRewardsCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ContractIRewardsCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIRewardsCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_ContractIRewardsCoordinator.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_ContractIRewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GENESISREWARDSTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "GENESIS_REWARDS_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_ContractIRewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) MAXFUTURELENGTH() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXFUTURELENGTH(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXFUTURELENGTH(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) MAXREWARDSDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "MAX_REWARDS_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) MAXREWARDSDURATION() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXREWARDSDURATION(&_ContractIRewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) MAXREWARDSDURATION() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.MAXREWARDSDURATION(&_ContractIRewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) ActivationDelay() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.ActivationDelay(&_ContractIRewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.ActivationDelay(&_ContractIRewardsCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.CalculateEarnerLeafHash(&_ContractIRewardsCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.CalculateEarnerLeafHash(&_ContractIRewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.CalculateTokenLeafHash(&_ContractIRewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.CalculateTokenLeafHash(&_ContractIRewardsCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CheckClaim(claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	return _ContractIRewardsCoordinator.Contract.CheckClaim(&_ContractIRewardsCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CheckClaim(claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	return _ContractIRewardsCoordinator.Contract.CheckClaim(&_ContractIRewardsCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "claimerFor", earner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractIRewardsCoordinator.Contract.ClaimerFor(&_ContractIRewardsCoordinator.CallOpts, earner)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractIRewardsCoordinator.Contract.ClaimerFor(&_ContractIRewardsCoordinator.CallOpts, earner)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, claimer common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "cumulativeClaimed", claimer, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIRewardsCoordinator.Contract.CumulativeClaimed(&_ContractIRewardsCoordinator.CallOpts, claimer, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address claimer, address token) view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CumulativeClaimed(claimer common.Address, token common.Address) (*big.Int, error) {
	return _ContractIRewardsCoordinator.Contract.CumulativeClaimed(&_ContractIRewardsCoordinator.CallOpts, claimer, token)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "currRewardsCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_ContractIRewardsCoordinator.CallOpts)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_ContractIRewardsCoordinator.CallOpts)
}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) DefaultOperatorSplitBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "defaultOperatorSplitBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) DefaultOperatorSplitBips() (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.DefaultOperatorSplitBips(&_ContractIRewardsCoordinator.CallOpts)
}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) DefaultOperatorSplitBips() (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.DefaultOperatorSplitBips(&_ContractIRewardsCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) DomainSeparator() ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.DomainSeparator(&_ContractIRewardsCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractIRewardsCoordinator.Contract.DomainSeparator(&_ContractIRewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getCurrentClaimableDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_ContractIRewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_ContractIRewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getCurrentDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetCurrentDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetCurrentDistributionRoot(&_ContractIRewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetCurrentDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetCurrentDistributionRoot(&_ContractIRewardsCoordinator.CallOpts)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getDistributionRootAtIndex", index)

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetDistributionRootAtIndex(&_ContractIRewardsCoordinator.CallOpts, index)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	return _ContractIRewardsCoordinator.Contract.GetDistributionRootAtIndex(&_ContractIRewardsCoordinator.CallOpts, index)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getDistributionRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetDistributionRootsLength() (*big.Int, error) {
	return _ContractIRewardsCoordinator.Contract.GetDistributionRootsLength(&_ContractIRewardsCoordinator.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetDistributionRootsLength() (*big.Int, error) {
	return _ContractIRewardsCoordinator.Contract.GetDistributionRootsLength(&_ContractIRewardsCoordinator.CallOpts)
}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetOperatorAVSSplit(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getOperatorAVSSplit", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.GetOperatorAVSSplit(&_ContractIRewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.GetOperatorAVSSplit(&_ContractIRewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetOperatorPISplit(opts *bind.CallOpts, operator common.Address) (uint16, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getOperatorPISplit", operator)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.GetOperatorPISplit(&_ContractIRewardsCoordinator.CallOpts, operator)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _ContractIRewardsCoordinator.Contract.GetOperatorPISplit(&_ContractIRewardsCoordinator.CallOpts, operator)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.GetRootIndexFromHash(&_ContractIRewardsCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractIRewardsCoordinator.Contract.GetRootIndexFromHash(&_ContractIRewardsCoordinator.CallOpts, rootHash)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCaller) RewardsUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIRewardsCoordinator.contract.Call(opts, &out, "rewardsUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) RewardsUpdater() (common.Address, error) {
	return _ContractIRewardsCoordinator.Contract.RewardsUpdater(&_ContractIRewardsCoordinator.CallOpts)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorCallerSession) RewardsUpdater() (common.Address, error) {
	return _ContractIRewardsCoordinator.Contract.RewardsUpdater(&_ContractIRewardsCoordinator.CallOpts)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "createOperatorDirectedAVSRewardsSubmission", avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractIRewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractIRewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "createRewardsForAllEarners", rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateRewardsForAllEarners(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateRewardsForAllEarners(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmission) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmission []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "createRewardsForAllSubmission", rewardsSubmission)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmission) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) CreateRewardsForAllSubmission(rewardsSubmission []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmission)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmission) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) CreateRewardsForAllSubmission(rewardsSubmission []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_ContractIRewardsCoordinator.TransactOpts, rewardsSubmission)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) DisableRoot(opts *bind.TransactOpts, rootIndex uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "disableRoot", rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.DisableRoot(&_ContractIRewardsCoordinator.TransactOpts, rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.DisableRoot(&_ContractIRewardsCoordinator.TransactOpts, rootIndex)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) ProcessClaim(claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ProcessClaim(&_ContractIRewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) ProcessClaim(claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ProcessClaim(&_ContractIRewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) ProcessClaims(opts *bind.TransactOpts, claims []IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "processClaims", claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) ProcessClaims(claims []IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ProcessClaims(&_ContractIRewardsCoordinator.TransactOpts, claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) ProcessClaims(claims []IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.ProcessClaims(&_ContractIRewardsCoordinator.TransactOpts, claims, recipient)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetActivationDelay(&_ContractIRewardsCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetActivationDelay(&_ContractIRewardsCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetClaimerFor(&_ContractIRewardsCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetClaimerFor(&_ContractIRewardsCoordinator.TransactOpts, claimer)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetDefaultOperatorSplit(opts *bind.TransactOpts, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setDefaultOperatorSplit", split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetDefaultOperatorSplit(&_ContractIRewardsCoordinator.TransactOpts, split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetDefaultOperatorSplit(&_ContractIRewardsCoordinator.TransactOpts, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetOperatorAVSSplit(opts *bind.TransactOpts, operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setOperatorAVSSplit", operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetOperatorAVSSplit(&_ContractIRewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetOperatorAVSSplit(&_ContractIRewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetOperatorPISplit(opts *bind.TransactOpts, operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setOperatorPISplit", operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetOperatorPISplit(&_ContractIRewardsCoordinator.TransactOpts, operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetOperatorPISplit(&_ContractIRewardsCoordinator.TransactOpts, operator, split)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetRewardsForAllSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setRewardsForAllSubmitter", _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_ContractIRewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_ContractIRewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SetRewardsUpdater(opts *bind.TransactOpts, _rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "setRewardsUpdater", _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetRewardsUpdater(&_ContractIRewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SetRewardsUpdater(&_ContractIRewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.contract.Transact(opts, "submitRoot", root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SubmitRoot(&_ContractIRewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorTransactorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractIRewardsCoordinator.Contract.SubmitRoot(&_ContractIRewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator is returned from FilterAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for AVSRewardsSubmissionCreated events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator struct {
	Event *ContractIRewardsCoordinatorAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorAVSRewardsSubmissionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorAVSRewardsSubmissionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorAVSRewardsSubmissionCreated represents a AVSRewardsSubmissionCreated event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorAVSRewardsSubmissionCreated struct {
	Avs                   common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterAVSRewardsSubmissionCreated(opts *bind.FilterOpts, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorAVSRewardsSubmissionCreatedIterator{contract: _ContractIRewardsCoordinator.contract, event: "AVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorAVSRewardsSubmissionCreated, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorAVSRewardsSubmissionCreated)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseAVSRewardsSubmissionCreated(log types.Log) (*ContractIRewardsCoordinatorAVSRewardsSubmissionCreated, error) {
	event := new(ContractIRewardsCoordinatorAVSRewardsSubmissionCreated)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorActivationDelaySetIterator struct {
	Event *ContractIRewardsCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorActivationDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorActivationDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractIRewardsCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorActivationDelaySetIterator{contract: _ContractIRewardsCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorActivationDelaySet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*ContractIRewardsCoordinatorActivationDelaySet, error) {
	event := new(ContractIRewardsCoordinatorActivationDelaySet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorClaimerForSetIterator struct {
	Event *ContractIRewardsCoordinatorClaimerForSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorClaimerForSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorClaimerForSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorClaimerForSet represents a ClaimerForSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*ContractIRewardsCoordinatorClaimerForSetIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorClaimerForSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorClaimerForSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimerForSet is a log parse operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*ContractIRewardsCoordinatorClaimerForSet, error) {
	event := new(ContractIRewardsCoordinatorClaimerForSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator is returned from FilterDefaultOperatorSplitBipsSet and is used to iterate over the raw logs and unpacked data for DefaultOperatorSplitBipsSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator struct {
	Event *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet represents a DefaultOperatorSplitBipsSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet struct {
	OldDefaultOperatorSplitBips uint16
	NewDefaultOperatorSplitBips uint16
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorSplitBipsSet is a free log retrieval operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterDefaultOperatorSplitBipsSet(opts *bind.FilterOpts) (*ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator, error) {

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorDefaultOperatorSplitBipsSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "DefaultOperatorSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorSplitBipsSet is a free log subscription operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchDefaultOperatorSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet) (event.Subscription, error) {

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultOperatorSplitBipsSet is a log parse operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseDefaultOperatorSplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet, error) {
	event := new(ContractIRewardsCoordinatorDefaultOperatorSplitBipsSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorDistributionRootDisabledIterator is returned from FilterDistributionRootDisabled and is used to iterate over the raw logs and unpacked data for DistributionRootDisabled events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDistributionRootDisabledIterator struct {
	Event *ContractIRewardsCoordinatorDistributionRootDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorDistributionRootDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorDistributionRootDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorDistributionRootDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorDistributionRootDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorDistributionRootDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorDistributionRootDisabled represents a DistributionRootDisabled event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDistributionRootDisabled struct {
	RootIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootDisabled is a free log retrieval operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterDistributionRootDisabled(opts *bind.FilterOpts, rootIndex []uint32) (*ContractIRewardsCoordinatorDistributionRootDisabledIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorDistributionRootDisabledIterator{contract: _ContractIRewardsCoordinator.contract, event: "DistributionRootDisabled", logs: logs, sub: sub}, nil
}

// WatchDistributionRootDisabled is a free log subscription operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchDistributionRootDisabled(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDistributionRootDisabled, rootIndex []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorDistributionRootDisabled)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributionRootDisabled is a log parse operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseDistributionRootDisabled(log types.Log) (*ContractIRewardsCoordinatorDistributionRootDisabled, error) {
	event := new(ContractIRewardsCoordinatorDistributionRootDisabled)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDistributionRootSubmittedIterator struct {
	Event *ContractIRewardsCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorDistributionRootSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorDistributionRootSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (*ContractIRewardsCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorDistributionRootSubmittedIterator{contract: _ContractIRewardsCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorDistributionRootSubmitted)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributionRootSubmitted is a log parse operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*ContractIRewardsCoordinatorDistributionRootSubmitted, error) {
	event := new(ContractIRewardsCoordinatorDistributionRootSubmitted)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator is returned from FilterOperatorAVSSplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorAVSSplitBipsSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator struct {
	Event *ContractIRewardsCoordinatorOperatorAVSSplitBipsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorOperatorAVSSplitBipsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorOperatorAVSSplitBipsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorOperatorAVSSplitBipsSet represents a OperatorAVSSplitBipsSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorAVSSplitBipsSet struct {
	Caller                  common.Address
	Operator                common.Address
	Avs                     common.Address
	ActivatedAt             uint32
	OldOperatorAVSSplitBips uint16
	NewOperatorAVSSplitBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSSplitBipsSet is a free log retrieval operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterOperatorAVSSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address, avs []common.Address) (*ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorOperatorAVSSplitBipsSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "OperatorAVSSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSSplitBipsSet is a free log subscription operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchOperatorAVSSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorAVSSplitBipsSet, caller []common.Address, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorOperatorAVSSplitBipsSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorAVSSplitBipsSet is a log parse operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseOperatorAVSSplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorOperatorAVSSplitBipsSet, error) {
	event := new(ContractIRewardsCoordinatorOperatorAVSSplitBipsSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator is returned from FilterOperatorDirectedAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for OperatorDirectedAVSRewardsSubmissionCreated events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator struct {
	Event *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated represents a OperatorDirectedAVSRewardsSubmissionCreated event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated struct {
	Caller                                common.Address
	Avs                                   common.Address
	OperatorDirectedRewardsSubmissionHash [32]byte
	SubmissionNonce                       *big.Int
	OperatorDirectedRewardsSubmission     IRewardsCoordinatorOperatorDirectedRewardsSubmission
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterOperatorDirectedAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator{contract: _ContractIRewardsCoordinator.contract, event: "OperatorDirectedAVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorDirectedAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDirectedAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseOperatorDirectedAVSRewardsSubmissionCreated(log types.Log) (*ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, error) {
	event := new(ContractIRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator is returned from FilterOperatorPISplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorPISplitBipsSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator struct {
	Event *ContractIRewardsCoordinatorOperatorPISplitBipsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorOperatorPISplitBipsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorOperatorPISplitBipsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorOperatorPISplitBipsSet represents a OperatorPISplitBipsSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorOperatorPISplitBipsSet struct {
	Caller                 common.Address
	Operator               common.Address
	ActivatedAt            uint32
	OldOperatorPISplitBips uint16
	NewOperatorPISplitBips uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOperatorPISplitBipsSet is a free log retrieval operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterOperatorPISplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorOperatorPISplitBipsSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "OperatorPISplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorPISplitBipsSet is a free log subscription operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchOperatorPISplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorOperatorPISplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorOperatorPISplitBipsSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorPISplitBipsSet is a log parse operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseOperatorPISplitBipsSet(log types.Log) (*ContractIRewardsCoordinatorOperatorPISplitBipsSet, error) {
	event := new(ContractIRewardsCoordinatorOperatorPISplitBipsSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsClaimedIterator struct {
	Event *ContractIRewardsCoordinatorRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorRewardsClaimed represents a RewardsClaimed event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsClaimed struct {
	Root          [32]byte
	Earner        common.Address
	Claimer       common.Address
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*ContractIRewardsCoordinatorRewardsClaimedIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorRewardsClaimedIterator{contract: _ContractIRewardsCoordinator.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorRewardsClaimed)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseRewardsClaimed(log types.Log) (*ContractIRewardsCoordinatorRewardsClaimed, error) {
	event := new(ContractIRewardsCoordinatorRewardsClaimed)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator is returned from FilterRewardsForAllSubmitterSet and is used to iterate over the raw logs and unpacked data for RewardsForAllSubmitterSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator struct {
	Event *ContractIRewardsCoordinatorRewardsForAllSubmitterSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorRewardsForAllSubmitterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorRewardsForAllSubmitterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorRewardsForAllSubmitterSet represents a RewardsForAllSubmitterSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsForAllSubmitterSet struct {
	RewardsForAllSubmitter common.Address
	OldValue               bool
	NewValue               bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRewardsForAllSubmitterSet is a free log retrieval operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterRewardsForAllSubmitterSet(opts *bind.FilterOpts, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (*ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorRewardsForAllSubmitterSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "RewardsForAllSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsForAllSubmitterSet is a free log subscription operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchRewardsForAllSubmitterSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsForAllSubmitterSet, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorRewardsForAllSubmitterSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsForAllSubmitterSet is a log parse operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseRewardsForAllSubmitterSet(log types.Log) (*ContractIRewardsCoordinatorRewardsForAllSubmitterSet, error) {
	event := new(ContractIRewardsCoordinatorRewardsForAllSubmitterSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator is returned from FilterRewardsSubmissionForAllCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllCreated events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator struct {
	Event *ContractIRewardsCoordinatorRewardsSubmissionForAllCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorRewardsSubmissionForAllCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorRewardsSubmissionForAllCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorRewardsSubmissionForAllCreated represents a RewardsSubmissionForAllCreated event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsSubmissionForAllCreated struct {
	Submitter             common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllCreated is a free log retrieval operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterRewardsSubmissionForAllCreated(opts *bind.FilterOpts, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorRewardsSubmissionForAllCreatedIterator{contract: _ContractIRewardsCoordinator.contract, event: "RewardsSubmissionForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllCreated is a free log subscription operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchRewardsSubmissionForAllCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsSubmissionForAllCreated, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorRewardsSubmissionForAllCreated)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsSubmissionForAllCreated is a log parse operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseRewardsSubmissionForAllCreated(log types.Log) (*ContractIRewardsCoordinatorRewardsSubmissionForAllCreated, error) {
	event := new(ContractIRewardsCoordinatorRewardsSubmissionForAllCreated)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator is returned from FilterRewardsSubmissionForAllEarnersCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllEarnersCreated events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator struct {
	Event *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated represents a RewardsSubmissionForAllEarnersCreated event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated struct {
	TokenHopper           common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllEarnersCreated is a free log retrieval operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterRewardsSubmissionForAllEarnersCreated(opts *bind.FilterOpts, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator{contract: _ContractIRewardsCoordinator.contract, event: "RewardsSubmissionForAllEarnersCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllEarnersCreated is a free log subscription operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchRewardsSubmissionForAllEarnersCreated(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsSubmissionForAllEarnersCreated is a log parse operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseRewardsSubmissionForAllEarnersCreated(log types.Log) (*ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, error) {
	event := new(ContractIRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRewardsCoordinatorRewardsUpdaterSetIterator is returned from FilterRewardsUpdaterSet and is used to iterate over the raw logs and unpacked data for RewardsUpdaterSet events raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsUpdaterSetIterator struct {
	Event *ContractIRewardsCoordinatorRewardsUpdaterSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIRewardsCoordinatorRewardsUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRewardsCoordinatorRewardsUpdaterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIRewardsCoordinatorRewardsUpdaterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIRewardsCoordinatorRewardsUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRewardsCoordinatorRewardsUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRewardsCoordinatorRewardsUpdaterSet represents a RewardsUpdaterSet event raised by the ContractIRewardsCoordinator contract.
type ContractIRewardsCoordinatorRewardsUpdaterSet struct {
	OldRewardsUpdater common.Address
	NewRewardsUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsUpdaterSet is a free log retrieval operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) FilterRewardsUpdaterSet(opts *bind.FilterOpts, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (*ContractIRewardsCoordinatorRewardsUpdaterSetIterator, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.FilterLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRewardsCoordinatorRewardsUpdaterSetIterator{contract: _ContractIRewardsCoordinator.contract, event: "RewardsUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsUpdaterSet is a free log subscription operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) WatchRewardsUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractIRewardsCoordinatorRewardsUpdaterSet, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (event.Subscription, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _ContractIRewardsCoordinator.contract.WatchLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRewardsCoordinatorRewardsUpdaterSet)
				if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsUpdaterSet is a log parse operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractIRewardsCoordinator *ContractIRewardsCoordinatorFilterer) ParseRewardsUpdaterSet(log types.Log) (*ContractIRewardsCoordinatorRewardsUpdaterSet, error) {
	event := new(ContractIRewardsCoordinatorRewardsUpdaterSet)
	if err := _ContractIRewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
