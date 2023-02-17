package validators

import (
	"errors"
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"

	msgProto "github.com/0xPolygon/go-ibft/messages/proto"
)

var (
	ErrInvalidDelWeightBps      = errors.New("invalid delegation weight bps")
	ErrVotingPowerAlreadyExists = errors.New("voting power already exists in voting powers set")
)

// Represents VotingPower of a Validator
type votingPower struct {
	value      *big.Int
	valAddress *types.Address
}

// NewVotingPower Creates a new VotingPower and return its reference
func NewVotingPower(valAddress types.Address, stakedBalance big.Int, delegatedBalance big.Int, delegationWeightBps big.Int) (VotingPower, error) {
	v, err := calculateVotingPower(&stakedBalance, &delegatedBalance, &delegationWeightBps)
	if err != nil {
		return nil, err
	}

	return &votingPower{value: v, valAddress: &valAddress}, nil
}

func (v *votingPower) GetValue() big.Int {
	return *v.value
}

func (v *votingPower) ValAddress() types.Address {
	return *v.valAddress
}

// calculateVotingPower calculates the voting power of validator based on both staked and delegated amount
func calculateVotingPower(
	stakedAmount *big.Int,
	delegatedAmount *big.Int,
	delegationWeightBps *big.Int, // delegationWeightBps / 1000 = delegation weight exponent
) (*big.Int, error) {
	if delegationWeightBps.Cmp(big.NewInt(0)) == -1 || delegationWeightBps.Cmp(big.NewInt(1000)) == 1 {
		return nil, ErrInvalidDelWeightBps
	}

	denominator := big.NewInt(1000)
	numerator := delegatedAmount.Mul(delegatedAmount, delegationWeightBps)

	if (numerator.Cmp(denominator)) == -1 {
		return stakedAmount, nil
	}

	return stakedAmount.Add(stakedAmount, numerator.Div(numerator, denominator)), nil
}

// Represents set with Voting Power of Validators set
type votingPowers struct {
	totalVPower *big.Int
	values      []VotingPower
}

// NewVotingPowers creates a new VotingPowers struct
func NewVotingPowers(vpowers ...VotingPower) VotingPowers {
	values := make([]VotingPower, len(vpowers))
	totalVPower := big.NewInt(0)

	for idx, val := range vpowers {
		values[idx] = val

		valPower := val.GetValue()
		totalVPower.Add(totalVPower, &valPower)
	}

	return &votingPowers{
		values:      values,
		totalVPower: totalVPower,
	}
}

// Add adds a voting power into the collection
// Updates the value if already added
func (vv *votingPowers) Add(v VotingPower) error {
	if vv.Includes(v.ValAddress()) {
		return ErrVotingPowerAlreadyExists
	}

	vv.values = append(vv.values, v)
	vv.increaseVotingPower(v.GetValue())

	return nil
}

// At returns a validator's Voting Power at specified index in the collection
// check outside the function does the index exists
func (vv *votingPowers) At(index uint64) VotingPower {
	return vv.values[index]
}

// Index returns the index of the VotingPower whose validator's address matches with the given address
func (vv *votingPowers) Index(addr types.Address) int64 {
	for i, val := range vv.values {
		if val.ValAddress() == addr {
			return int64(i)
		}
	}

	return -1
}

// Includes return the bool indicating whether the voting power of validator
// whose address matches with the given address exists or not
func (vv *votingPowers) Includes(addr types.Address) bool {
	return vv.Index(addr) != -1
}

func (vv *votingPowers) GetTVotingPower() big.Int {
	return *vv.totalVPower
}

func (vv *votingPowers) GetVotingPower(val types.Address) (big.Int, error) {
	index := vv.Index(val)
	if index == -1 {
		return *big.NewInt(0), ErrVPowerNotFound
	}

	return vv.At(uint64(index)).GetValue(), nil
}

func (vv *votingPowers) increaseVotingPower(amount big.Int) {
	vv.totalVPower.Add(vv.totalVPower, &amount)
}

// calcMessagesPower
func (vv *votingPowers) CalcMessagesPower(messages []*msgProto.Message) (big.Int, error) {
	sum := big.NewInt(0)
	for _, msg := range messages {
		vPower, err := vv.GetVotingPower(types.BytesToAddress(msg.GetFrom()))
		if err != nil {
			return *big.NewInt(0), err
		}
		sum.Add(sum, &vPower)
	}

	return *sum, nil
}
