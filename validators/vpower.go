package validators

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"

	msgProto "github.com/0xPolygon/go-ibft/messages/proto"
)

// Represents VotingPower of a Validator
type votingPower struct {
	value      *big.Int
	valAddress *types.Address
}

func NewVotingPower(valAddress types.Address, stakedBalance big.Int, delegatedBalance big.Int) VotingPower {
	return &votingPower{value: big.NewInt(1), valAddress: &valAddress}
}

func (v *votingPower) GetValue() big.Int {
	return *v.value
}

func (v *votingPower) VallAddress() types.Address {
	return *v.valAddress
}

// Represents set with Voting Power of Validators set
type votingPowers struct {
	totalVPower *big.Int
	values      []VotingPower
}

// NewVotingPowers creates an empty VotingPowers struct
func NewVotingPowers() VotingPowers {
	return &votingPowers{}
}

// Add adds a voting power into the collection
// Updates the value if already added
func (vv *votingPowers) Add(v VotingPower) error {
	if vv.Includes(v.VallAddress()) {
		return ErrValidatorAlreadyExists
	}

	vv.values = append(vv.values, v)

	return nil
}

// At returns a validator's Voting Power at specified index in the collection
func (vv *votingPowers) At(index uint64) VotingPower {
	return vv.values[index]
}

// Index returns the index of the VotingPower whose validator's address matches with the given address
func (vv *votingPowers) Index(addr types.Address) int64 {
	for i, val := range vv.values {
		if val.VallAddress() == addr {
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
		vpower, err := vv.GetVotingPower(types.BytesToAddress(msg.From))
		if err != nil {
			return *big.NewInt(0), err
		}
		sum.Add(sum, &vpower)
	}

	return *sum, nil
}
