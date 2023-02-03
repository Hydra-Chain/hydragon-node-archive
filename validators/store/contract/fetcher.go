package contract

import (
	"fmt"
	"math/big"

	"github.com/0xPolygon/polygon-edge/contracts/staking"
	"github.com/0xPolygon/polygon-edge/crypto"
	"github.com/0xPolygon/polygon-edge/state"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/0xPolygon/polygon-edge/validators"
)

// FetchValidators fetches validators from a contract switched by validator type
func FetchValidators(
	validatorType validators.ValidatorType,
	transition *state.Transition,
	from types.Address,
) (validators.Validators, validators.VotingPowers, error) {
	switch validatorType {
	case validators.ECDSAValidatorType:
		return FetchECDSAValidators(transition, from)
	case validators.BLSValidatorType:
		return FetchBLSValidators(transition, from)
	}

	return nil, nil, fmt.Errorf("unsupported validator type: %s", validatorType)
}

// FetchECDSAValidators queries a contract for validator addresses and returns ECDSAValidators
// TODO: Modify based on contract implementation
func FetchECDSAValidators(
	transition *state.Transition,
	from types.Address,
) (validators.Validators, validators.VotingPowers, error) {
	valAddrs, err := staking.QueryValidators(transition, from)
	if err != nil {
		return nil, nil, err
	}

	ecdsaValidators := validators.NewECDSAValidatorSet()
	votingPowers := validators.NewVotingPowers()
	for _, addr := range valAddrs {
		if err := ecdsaValidators.Add(validators.NewECDSAValidator(addr)); err != nil {
			return nil, nil, err
		}

		if err := votingPowers.Add(validators.NewVotingPower(addr, *big.NewInt(0), *big.NewInt(0))); err != nil {
			return ecdsaValidators, nil, err
		}
	}

	return ecdsaValidators, votingPowers, nil
}

// FetchBLSValidators queries a contract for validator addresses & BLS Public Keys and returns ECDSAValidators
// TODO: Modify based on contract implementation
func FetchBLSValidators(
	transition *state.Transition,
	from types.Address,
) (validators.Validators, validators.VotingPowers, error) {
	valAddrs, err := staking.QueryValidators(transition, from)
	if err != nil {
		return nil, nil, err
	}

	blsPublicKeys, err := staking.QueryBLSPublicKeys(transition, from)
	if err != nil {
		return nil, nil, err
	}

	blsValidators := validators.NewBLSValidatorSet()
	votingPowers := validators.NewVotingPowers()

	for idx := range valAddrs {
		// ignore the validator whose BLS Key is not set
		// because BLS validator needs to have both Address and BLS Public Key set
		// in the contract
		if _, err := crypto.UnmarshalBLSPublicKey(blsPublicKeys[idx]); err != nil {
			continue
		}

		if err := blsValidators.Add(validators.NewBLSValidator(
			valAddrs[idx],
			blsPublicKeys[idx],
		)); err != nil {
			return nil, nil, err
		}

		if err := votingPowers.Add(validators.NewVotingPower(valAddrs[idx], *big.NewInt(0), *big.NewInt(0))); err != nil {
			return blsValidators, nil, err
		}
	}

	return blsValidators, votingPowers, nil
}
