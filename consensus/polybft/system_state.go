package polybft

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xPolygon/polygon-edge/consensus/polybft/contractsapi"
	bls "github.com/0xPolygon/polygon-edge/consensus/polybft/signer"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
)

// VotingPowerExponent is a data transfer object which holds voting power exponent used to balance the voting power
type VotingPowerExponent struct {
	Numerator   *big.Int
	Denominator *big.Int
}

// ValidatorInfo is data transfer object which holds validator information,
// provided by smart contract
type ValidatorInfo struct {
	Address             ethgo.Address
	Stake               *big.Int
	WithdrawableRewards *big.Int
	IsActive            bool
	IsWhitelisted       bool
}

// SystemState is an interface to interact with the consensus system contracts in the chain
type SystemState interface {
	// GetEpoch retrieves current epoch number from the smart contract
	GetEpoch() (uint64, error)
	// GetNextCommittedIndex retrieves next committed bridge state sync index
	GetNextCommittedIndex() (uint64, error)
	// GetStakeOnValidatorSet retrieves stake of given validator on ValidatorSet contract
	GetStakeOnValidatorSet(validatorAddr types.Address) (*big.Int, error)
	// GetVotingPowerExponent retrieves voting power exponent from the ChildValidatorSet smart contract
	GetVotingPowerExponent() (exponent *VotingPowerExponent, err error)
	// GetValidatorBlsKey retrieves validator BLS public key from the ChildValidatorSet smart contract
	GetValidatorBlsKey(addr types.Address) (*bls.PublicKey, error)
}

var _ SystemState = &SystemStateImpl{}

// SystemStateImpl is implementation of SystemState interface
type SystemStateImpl struct {
	validatorContract       *contract.Contract
	sidechainBridgeContract *contract.Contract
}

// NewSystemState initializes new instance of systemState which abstracts smart contracts functions
func NewSystemState(valSetAddr types.Address, stateRcvAddr types.Address, provider contract.Provider) *SystemStateImpl {
	// H_MODIFY: Use ChildValidatorSet abi
	s := &SystemStateImpl{}
	s.validatorContract = contract.NewContract(
		ethgo.Address(valSetAddr),
		contractsapi.ChildValidatorSet.Abi, contract.WithProvider(provider),
	)

	s.sidechainBridgeContract = contract.NewContract(
		ethgo.Address(stateRcvAddr),
		contractsapi.StateReceiver.Abi,
		contract.WithProvider(provider),
	)

	return s
}

// H_MODIFY: Get validator stake from childValidatorSet contract using the getValidatorTotalStake function
// TODO: getValidatorTotalStake is a temporary solution and must be removed
// (check the func in the contract for more info)
// GetStakeOnValidatorSet retrieves stake of given validator on ValidatorSet contract
func (s *SystemStateImpl) GetStakeOnValidatorSet(validatorAddr types.Address) (*big.Int, error) {
	rawResult, err := s.validatorContract.Call("getValidatorTotalStake", ethgo.Latest, validatorAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to call getValidator function: %w", err)
	}

	stake, isOk := rawResult["stake"].(*big.Int)
	if !isOk {
		return nil, fmt.Errorf("failed to decode stake")
	}

	// in case stake is 0 validator is not active no mather is there a delegated balance
	if stake.Cmp(big.NewInt(0)) == 0 {
		return bigZero, nil
	}

	totalStake, isOk := rawResult["totalStake"].(*big.Int)
	if !isOk {
		return nil, fmt.Errorf("failed to decode totalStake")
	}

	return totalStake, nil
}

// GetEpoch retrieves current epoch number from the smart contract
func (s *SystemStateImpl) GetEpoch() (uint64, error) {
	// check how system tx is executed and why epoch data is not changed
	rawResult, err := s.validatorContract.Call("currentEpochId", ethgo.Latest)
	if err != nil {
		return 0, err
	}

	epochNumber, isOk := rawResult["0"].(*big.Int)

	fmt.Println("HEERE IS IT epochNumber", epochNumber)

	if !isOk {
		return 0, fmt.Errorf("failed to decode epoch")
	}

	return epochNumber.Uint64(), nil
}

// H: add a function to fetch the voting power exponent
func (s *SystemStateImpl) GetVotingPowerExponent() (exponent *VotingPowerExponent, err error) {
	rawOutput, err := s.validatorContract.Call("getExponent", ethgo.Latest)
	if err != nil {
		return nil, err
	}

	expNumerator, ok := rawOutput["numerator"].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to decode voting power exponent numerator")
	}

	expDenominator, ok := rawOutput["denominator"].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to decode voting power exponent denominator")
	}

	return &VotingPowerExponent{Numerator: expNumerator, Denominator: expDenominator}, nil
}

// H: add a function to fetch the validator bls key
func (s *SystemStateImpl) GetValidatorBlsKey(addr types.Address) (*bls.PublicKey, error) {
	rawOutput, err := s.validatorContract.Call("getValidator", ethgo.Latest, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to call getValidator function: %w", err)
	}

	rawKey, ok := rawOutput["blsKey"].([4]*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to decode blskey")
	}

	blsKey, err := bls.UnmarshalPublicKeyFromBigInt(rawKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal BLS public key: %w", err)
	}

	return blsKey, nil
}

// GetNextCommittedIndex retrieves next committed bridge state sync index
func (s *SystemStateImpl) GetNextCommittedIndex() (uint64, error) {
	rawResult, err := s.sidechainBridgeContract.Call("lastCommittedId", ethgo.Latest)
	if err != nil {
		return 0, err
	}

	nextCommittedIndex, isOk := rawResult["0"].(*big.Int)
	if !isOk {
		return 0, fmt.Errorf("failed to decode next committed index")
	}

	return nextCommittedIndex.Uint64() + 1, nil
}

func buildLogsFromReceipts(entry []*types.Receipt, header *types.Header) []*types.Log {
	var logs []*types.Log

	for _, taskReceipt := range entry {
		for _, taskLog := range taskReceipt.Logs {
			log := new(types.Log)
			*log = *taskLog

			data := map[string]interface{}{
				"Hash":   header.Hash,
				"Number": header.Number,
			}
			log.Data, _ = json.Marshal(&data)
			logs = append(logs, log)
		}
	}

	return logs
}
