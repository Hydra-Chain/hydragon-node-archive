package validators

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xPolygon/polygon-edge/types"
)

var (
	ErrInvalidBLSValidatorFormat = errors.New("invalid validator format, expected [Validator Address]:[BLS Public Key]:[Staked Amount]")
)

// NewValidatorFromType instantiates a validator by specified type
func NewValidatorFromType(t ValidatorType) (Validator, error) {
	switch t {
	case ECDSAValidatorType:
		return new(ECDSAValidator), nil
	case BLSValidatorType:
		return new(BLSValidator), nil
	}

	return nil, ErrInvalidValidatorType
}

// NewValidatorSetFromType instantiates a validators by specified type
func NewValidatorSetFromType(t ValidatorType) Validators {
	switch t {
	case ECDSAValidatorType:
		return NewECDSAValidatorSet()
	case BLSValidatorType:
		return NewBLSValidatorSet()
	}

	return nil
}

// NewECDSAValidatorSet creates Validator Set for ECDSAValidator with initialized validators
func NewECDSAValidatorSet(ecdsaValidators ...*ECDSAValidator) Validators {
	validators := make([]Validator, len(ecdsaValidators))

	for idx, val := range ecdsaValidators {
		validators[idx] = Validator(val)
	}

	return &Set{
		ValidatorType: ECDSAValidatorType,
		Validators:    validators,
	}
}

// NewBLSValidatorSet creates Validator Set for BLSValidator with initialized validators
func NewBLSValidatorSet(blsValidators ...*BLSValidator) Validators {
	validators := make([]Validator, len(blsValidators))
	votingPower := big.NewInt(0)
	for idx, val := range blsValidators {
		validators[idx] = Validator(val)
		votingPower.Add(votingPower, &val.VotingPower)
	}

	return &Set{
		ValidatorType:    BLSValidatorType,
		Validators:       validators,
		TotalVotingPower: *votingPower,
	}
}

// ParseValidator parses a validator represented in string
func ParseValidator(validatorType ValidatorType, validator string) (Validator, error) {
	switch validatorType {
	case ECDSAValidatorType:
		return ParseECDSAValidator(validator), nil
	case BLSValidatorType:
		return ParseBLSValidator(validator)
	default:
		// shouldn't reach here
		return nil, fmt.Errorf("invalid validator type: %s", validatorType)
	}
}

// ParseValidator parses an array of validator represented in string
func ParseValidators(validatorType ValidatorType, rawValidators []string) (Validators, error) {
	set := NewValidatorSetFromType(validatorType)
	if set == nil {
		return nil, fmt.Errorf("invalid validator type: %s", validatorType)
	}

	for _, s := range rawValidators {
		validator, err := ParseValidator(validatorType, s)
		if err != nil {
			return nil, err
		}

		if err := set.Add(validator); err != nil {
			return nil, err
		}
	}

	return set, nil
}

// ParseBLSValidator parses ECDSAValidator represented in string
func ParseECDSAValidator(validator string) *ECDSAValidator {
	return &ECDSAValidator{
		Address: types.StringToAddress(validator),
	}
}

// ParseBLSValidator parses BLSValidator represented in string
func ParseBLSValidator(validator string) (*BLSValidator, error) {
	subValues := strings.Split(validator, ":")

	if len(subValues) != 3 {
		return nil, ErrInvalidBLSValidatorFormat
	}

	addrBytes, err := hex.DecodeString(strings.TrimPrefix(subValues[0], "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	pubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(subValues[1], "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse BLS Public Key: %w", err)
	}

	staked := new(big.Int)
	staked, succ := staked.SetString(subValues[2], 10)
	if !succ {
		return nil, fmt.Errorf("failed to parse Staked Amount: %w", err)
	}

	return &BLSValidator{
		Address:      types.BytesToAddress(addrBytes),
		BLSPublicKey: pubKeyBytes,
		VotingPower:  *staked,
	}, nil
}

// calculateVotingPower calculates the voting power of validator based on both staked and delegated amount
func CalculateVotingPower(
	stakedAmount *big.Int,
	delegatedAmount *big.Int,
	delegationWeightBps *big.Int, // delegationWeightBps / 1000 = delegation weight exponent
) (*big.Int, error) {
	if delegationWeightBps.Cmp(big.NewInt(0)) == -1 || delegationWeightBps.Cmp(big.NewInt(1000)) == 1 {
		return nil, errors.New("invalid delegation weight bps")
	}

	denominator := big.NewInt(1000)
	numerator := delegatedAmount.Mul(delegatedAmount, delegationWeightBps)

	if (numerator.Cmp(denominator)) == -1 {
		return stakedAmount, nil
	}

	return stakedAmount.Add(stakedAmount, numerator.Div(numerator, denominator)), nil
}
