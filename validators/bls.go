package validators

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygon/polygon-edge/helper/hex"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/umbracle/fastrlp"
)

var (
	ErrInvalidTypeAssert = errors.New("invalid type assert")
)

type BLSValidatorPublicKey []byte

// String returns a public key in hex
func (k BLSValidatorPublicKey) String() string {
	return hex.EncodeToHex(k[:])
}

// MarshalText implements encoding.TextMarshaler
func (k BLSValidatorPublicKey) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

// UnmarshalText parses an BLS Public Key in hex
func (k *BLSValidatorPublicKey) UnmarshalText(input []byte) error {
	kk, err := hex.DecodeHex(string(input))
	if err != nil {
		return err
	}

	*k = kk

	return nil
}

// BLSValidator is a validator using BLS signing algorithm
type BLSValidator struct {
	Address      types.Address
	BLSPublicKey BLSValidatorPublicKey
	votingPower  big.Int
}

// NewBLSValidator is a constructor of BLSValidator
func NewBLSValidator(addr types.Address, blsPubkey []byte, stakedBalance big.Int) *BLSValidator {
	return &BLSValidator{
		Address:      addr,
		BLSPublicKey: blsPubkey,
		votingPower:  stakedBalance,
	}
}

// Type returns the ValidatorType of BLSValidator
func (v *BLSValidator) Type() ValidatorType {
	return BLSValidatorType
}

// String returns string representation of BLSValidator
// Format => [Address]:[BLSPublicKey]
func (v *BLSValidator) String() string {
	return fmt.Sprintf(
		"%s:%s (Staked Balance: %s)",
		v.Address.String(),
		hex.EncodeToHex(v.BLSPublicKey),
		v.votingPower.Text(10),
	)
}

// Addr returns the validator address
func (v *BLSValidator) Addr() types.Address {
	return v.Address
}

// Copy returns copy of BLS Validator
func (v *BLSValidator) Copy() Validator {
	pubkey := make([]byte, len(v.BLSPublicKey))
	copy(pubkey, v.BLSPublicKey)

	return &BLSValidator{
		Address:      v.Address,
		BLSPublicKey: pubkey,
		votingPower:  v.votingPower,
	}
}

// Equal checks the given validator matches with its data
func (v *BLSValidator) Equal(vr Validator) bool {
	vv, ok := vr.(*BLSValidator)
	if !ok {
		return false
	}

	return v.Address == vv.Address && bytes.Equal(v.BLSPublicKey, vv.BLSPublicKey) && v.votingPower.Cmp(&vv.votingPower) == 0
}

// MarshalRLPWith is a RLP Marshaller
// Didn't add Voting Power in the RLP format because we don't want to modify the extraData of the block header
// It is going to work properly only with "contract store" configurations
func (v *BLSValidator) MarshalRLPWith(arena *fastrlp.Arena) *fastrlp.Value {
	vv := arena.NewArray()

	vv.Set(arena.NewBytes(v.Address.Bytes()))
	vv.Set(arena.NewCopyBytes(v.BLSPublicKey))

	return vv
}

// UnmarshalRLPFrom is a RLP Unmarshaller
// Didn't add Voting Power in the unmarshaling because we don't include it in the extraData of the block header
// It is going to work properly only with "contract store" configurations
func (v *BLSValidator) UnmarshalRLPFrom(p *fastrlp.Parser, val *fastrlp.Value) error {
	elems, err := val.GetElems()
	if err != nil {
		return err
	}

	if len(elems) < 2 {
		return fmt.Errorf("incorrect number of elements to decode BLSValidator, expected 2 but found %d", len(elems))
	}

	if err := elems[0].GetAddr(v.Address[:]); err != nil {
		return fmt.Errorf("failed to decode Address: %w", err)
	}

	if v.BLSPublicKey, err = elems[1].GetBytes(v.BLSPublicKey); err != nil {
		return fmt.Errorf("failed to decode BLSPublicKey: %w", err)
	}

	return nil
}

// Bytes returns bytes of BLSValidator in RLP encode
func (v *BLSValidator) Bytes() []byte {
	return types.MarshalRLPTo(v.MarshalRLPWith, nil)
}

// SetFromBytes parses given bytes in RLP encode and map to its fields
func (v *BLSValidator) SetFromBytes(input []byte) error {
	return types.UnmarshalRlp(v.UnmarshalRLPFrom, input)
}

// Addr returns the validator Voting Power (Staked Balance)
func (v *BLSValidator) VotingPower() big.Int {
	return v.votingPower
}
