package polybft

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/hashicorp/go-hclog"
)

// ValidatorSet interface of the current validator set
type ValidatorSet interface {
	// Includes check if given address is among the current validator set
	Includes(address types.Address) bool

	// Len returns the size of the validator set
	Len() int

	// Accounts returns the list of the ValidatorMetadata
	Accounts() AccountSet

	// checks if submitted signers have reached quorum
	HasQuorum(signers map[types.Address]struct{}) bool
}

type validatorSet struct {
	// validators represents current list of validators
	validators AccountSet

	// votingPowerMap represents voting powers per validator address
	votingPowerMap map[types.Address]*big.Int

	// totalVotingPower is sum of active validator set
	totalVotingPower *big.Int

	// quorumSize is 2/3 super-majority of totalVotingPower
	quorumSize *big.Int

	// logger instance
	logger hclog.Logger
}

// NewValidatorSet creates a new validator set.
func NewValidatorSet(valz AccountSet, logger hclog.Logger) *validatorSet {
	totalVotingPower := valz.GetTotalVotingPower()
	quorumSize := getQuorumSize(totalVotingPower)

	votingPowerMap := make(map[types.Address]*big.Int, len(valz))
	for _, val := range valz {
		votingPowerMap[val.Address] = val.VotingPower
	}

	return &validatorSet{
		validators:       valz,
		votingPowerMap:   votingPowerMap,
		totalVotingPower: totalVotingPower,
		quorumSize:       quorumSize,
		logger:           logger.Named("validator_set"),
	}
}

// HasQuorum determines if there is quorum of enough signers reached,
// based on its voting power and quorum size
func (vs validatorSet) HasQuorum(signers map[types.Address]struct{}) bool {
	aggregateVotingPower := big.NewInt(0)
	valsCount := 0

	for address := range signers {
		if votingPower := vs.votingPowerMap[address]; votingPower != nil {
			_ = aggregateVotingPower.Add(aggregateVotingPower, votingPower)

			valsCount++
		}
	}

	hasQuorum := false
	quorumSize := vs.quorumSize
	if valsCount < 4 {
		quorumSize = vs.totalVotingPower
	}

	hasQuorum = aggregateVotingPower.Cmp(quorumSize) >= 0

	vs.logger.Debug("HasQuorum",
		"signers", len(signers),
		"signers voting power", aggregateVotingPower,
		"quorum size", quorumSize,
		"hasQuorum", hasQuorum)

	return hasQuorum
}

func (vs validatorSet) Accounts() AccountSet {
	return vs.validators
}

func (vs validatorSet) Includes(address types.Address) bool {
	return vs.validators.ContainsAddress(address)
}

func (vs validatorSet) Len() int {
	return vs.validators.Len()
}

// getQuorumSize calculates quorum size as 2/3 super-majority of provided total voting power
// H_MODIFY: Qourum size is 614/1000 = 61.4% of total voting power
func getQuorumSize(totalVotingPower *big.Int) *big.Int {
	quorum := new(big.Int)
	// In case for some reason voting power goes below too small value, we set quorum to total voting power
	if totalVotingPower.Cmp(big.NewInt(10)) == -1 {
		quorum = quorum.Add(quorum, totalVotingPower)
	} else {
		quorum = quorum.Mul(totalVotingPower, big.NewInt(614)).Div(quorum, big.NewInt(1000)).Add(quorum, big.NewInt(1))
	}

	return quorum
}
