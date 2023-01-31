package ibft

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/0xPolygon/polygon-edge/validators"
)

func CalcMaxFaultyNodes(s validators.Validators) int {
	// N -> number of nodes in IBFT
	// F -> number of faulty nodes
	//
	// N = 3F + 1
	// => F = (N - 1) / 3
	//
	// IBFT tolerates 1 failure with 4 nodes
	// 4 = 3 * 1 + 1
	// To tolerate 2 failures, IBFT requires 7 nodes
	// 7 = 3 * 2 + 1
	// It should always take the floor of the result
	return (s.Len() - 1) / 3
}

type QuorumImplementation func(validators.Validators) big.Int

// LegacyQuorumSize returns the legacy quorum size for the given validator set
// H_MODIFY legacy is irelevant - we set the Optimal quorum size formula just in case
func LegacyQuorumSize(set validators.Validators) big.Int {
	tVotingPower := set.TotalVPower()
	//	if the number of validators is less than 4,
	//	then the entire set is required
	if CalcMaxFaultyNodes(set) == 0 {
		/*
			N: 1 -> Q: 1
			N: 2 -> Q: 2
			N: 3 -> Q: 3
		*/
		return tVotingPower
	}

	// (quorum optimal)	Q = ceil(2/3 * N)
	// H_MODIFY: qorum = 61.4% of total voting power (voting power * 614/1000)
	// H_MODIFY: assume that voting power would be always bigger than 15000
	// TODO: Add unit tests for quorum calc
	divisible := tVotingPower.Mul(&tVotingPower, big.NewInt(614))
	res := divisible.Div(divisible, big.NewInt(1000))
	return *res.Add(res, big.NewInt(1))
}

// OptimalQuorumSize returns the optimal quorum size for the given validator set
// H_MODIFY: We change the quorum calculation to be based on the staked balance.
// That way we anble the delegation functionality
func OptimalQuorumSize(set validators.Validators) big.Int {
	tVotingPower := set.TotalVPower()
	//	if the number of validators is less than 4,
	//	then the entire set is required
	if CalcMaxFaultyNodes(set) == 0 {
		/*
			N: 1 -> Q: 1
			N: 2 -> Q: 2
			N: 3 -> Q: 3
		*/
		return tVotingPower
	}

	// (quorum optimal)	Q = ceil(2/3 * N)
	// H_MODIFY: qorum = 61.4% of total voting power (voting power * 614/1000)
	// H_MODIFY: assume that voting power would be always bigger than 15000
	// TODO: Add unit tests for quorum calc
	divisible := tVotingPower.Mul(&tVotingPower, big.NewInt(614))
	res := divisible.Div(divisible, big.NewInt(1000))
	return *res.Add(res, big.NewInt(1))
}

func CalcProposer(
	validators validators.Validators,
	round uint64,
	lastProposer types.Address,
) validators.Validator {
	var seed uint64

	if lastProposer == types.ZeroAddress {
		seed = round
	} else {
		offset := int64(0)

		if index := validators.Index(lastProposer); index != -1 {
			offset = index
		}

		seed = uint64(offset) + round + 1
	}

	pick := seed % uint64(validators.Len())

	return validators.At(pick)
}
