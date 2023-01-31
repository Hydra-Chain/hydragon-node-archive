package validators

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECDSAValidatorsMarshalJSON(t *testing.T) {
	t.Parallel()

	validators := &Set{
		ValidatorType: ECDSAValidatorType,
		Validators: []Validator{
			&ECDSAValidator{addr1},
			&ECDSAValidator{addr2},
		},
	}

	res, err := json.Marshal(validators)

	assert.NoError(t, err)

	assert.JSONEq(
		t,
		fmt.Sprintf(
			`[
				{
					"Address": "%s"
				},
				{
					"Address": "%s"
				}
			]`,
			addr1.String(),
			addr2.String(),
		),
		string(res),
	)
}

func TestECDSAValidatorsUnmarshalJSON(t *testing.T) {
	t.Parallel()

	inputStr := fmt.Sprintf(
		`[
			{
				"Address": "%s"
			},
			{
				"Address": "%s"
			}
		]`,
		addr1.String(),
		addr2.String(),
	)

	validators := NewECDSAValidatorSet()

	assert.NoError(
		t,
		json.Unmarshal([]byte(inputStr), validators),
	)

	assert.Equal(
		t,
		&Set{
			ValidatorType: ECDSAValidatorType,
			Validators: []Validator{
				&ECDSAValidator{addr1},
				&ECDSAValidator{addr2},
			},
		},
		validators,
	)
}

func TestBLSValidatorsMarshalJSON(t *testing.T) {
	t.Parallel()

	validators := &Set{
		ValidatorType: BLSValidatorType,
		Validators: []Validator{
			&BLSValidator{addr1, testBLSPubKey1, *OneHydraBig},
			&BLSValidator{addr2, testBLSPubKey2, *OneHydraBig},
		},
	}

	res, err := json.Marshal(validators)

	assert.NoError(t, err)

	assert.JSONEq(
		t,
		fmt.Sprintf(
			`[
				{
					"Address": "%s",
					"BLSPublicKey": "%s",
					"VotingPower": %s
				},
				{
					"Address": "%s",
					"BLSPublicKey": "%s",
					"VotingPower":%s
				}
			]`,
			addr1,
			testBLSPubKey1,
			"1000000000000000000",
			addr2,
			testBLSPubKey2,
			"1000000000000000000",
		),
		string(res),
	)
}

func TestBLSValidatorsUnmarshalJSON(t *testing.T) {
	t.Parallel()

	inputStr := fmt.Sprintf(
		`[
			{
				"Address": "%s",
				"BLSPublicKey": "%s",
				"VotingPower": %s
			},
			{
				"Address": "%s",
				"BLSPublicKey": "%s",
				"VotingPower":%s
			}
		]`,
		addr1,
		testBLSPubKey1,
		"1000000000000000000",
		addr2,
		testBLSPubKey2,
		"1000000000000000000",
	)

	validators := NewBLSValidatorSet()

	assert.NoError(
		t,
		json.Unmarshal([]byte(inputStr), validators),
	)

	assert.Equal(
		t,
		&Set{
			ValidatorType: BLSValidatorType,
			Validators: []Validator{
				&BLSValidator{addr1, testBLSPubKey1, *OneHydraBig},
				&BLSValidator{addr2, testBLSPubKey2, *OneHydraBig},
			},
			TotalVotingPower: *big.NewInt(0).Add(OneHydraBig, OneHydraBig),
		},
		validators,
	)
}
