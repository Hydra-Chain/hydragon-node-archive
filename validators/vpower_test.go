package validators

import (
	"math/big"
	"testing"

	msgProto "github.com/0xPolygon/go-ibft/messages/proto"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/stretchr/testify/assert"
)

func newTestVotingPower(addr types.Address) VotingPower {
	vPower, err := NewVotingPower(addr, *big.NewInt(15000), *big.NewInt(0), *big.NewInt(850))
	if err != nil {
		panic("newTestVotingPower must always return a proper VotingPower")
	}

	return vPower
}

func testBigInt(str string) *big.Int {
	amount, success := new(big.Int).SetString(str, 10)
	if !success {
		panic("it must succeed")
	}

	return amount
}

func Test_NewVotingPower(t *testing.T) {
	t.Parallel()

	t.Run("should return error if voting power calculation is unsuccessful", func(t *testing.T) {
		res, err := NewVotingPower(addr1, *testStakedBalance, *testDelegatedBalance, *big.NewInt(1001))

		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("should successfuly create VotingPower", func(t *testing.T) {
		vpower, err := NewVotingPower(addr1, *testStakedBalance, *testDelegatedBalance, *testdelegationWeightBps)

		assert.Equal(
			t,
			&votingPower{value: big.NewInt(15000), valAddress: &addr1},
			vpower,
		)
		assert.Nil(t, err)
	})
}

func Test_calculateVotingPower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		stakedBalance       *big.Int
		delegatedBalance    *big.Int
		delegationWeightBps *big.Int
		expectedRes         *big.Int
		expectedErr         error
	}{
		{
			name:                "should return error when too small delegationWeightBps",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(0),
			delegationWeightBps: big.NewInt(-1),
			expectedRes:         nil,
			expectedErr:         ErrInvalidDelWeightBps,
		},
		{
			name:                "should return error when too big delegationWeightBps",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(0),
			delegationWeightBps: big.NewInt(1001),
			expectedRes:         nil,
			expectedErr:         ErrInvalidDelWeightBps,
		},
		{
			name:                "should return only staked amount when delegated amount too small",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(1),
			delegationWeightBps: big.NewInt(850),
			expectedRes:         big.NewInt(15000),
			expectedErr:         nil,
		},
		{
			name:                "should return only staked amount when delegationWeightBps is 0",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(100000),
			delegationWeightBps: big.NewInt(0),
			expectedRes:         big.NewInt(15000),
			expectedErr:         nil,
		},
		{
			name:                "should properly calculate 1",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(1000),
			delegationWeightBps: big.NewInt(1),
			expectedRes:         big.NewInt(15001),
			expectedErr:         nil,
		},
		{
			name:                "should properly calculate 2",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(500),
			delegationWeightBps: big.NewInt(1000),
			expectedRes:         big.NewInt(15500),
			expectedErr:         nil,
		},
		{
			name:                "should properly calculate 3",
			stakedBalance:       big.NewInt(15000),
			delegatedBalance:    big.NewInt(15000),
			delegationWeightBps: big.NewInt(850),
			expectedRes:         big.NewInt(27750),
			expectedErr:         nil,
		},
		{
			name:                "should properly calculate when amount is very big",
			stakedBalance:       testBigInt("9000000000000000000000000000000"),
			delegatedBalance:    testBigInt("9000000000000000000000000000000"),
			delegationWeightBps: big.NewInt(850),
			expectedRes:         testBigInt("16650000000000000000000000000000"),
			expectedErr:         nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := calculateVotingPower(test.stakedBalance, test.delegatedBalance, test.delegationWeightBps)

			assert.Equal(t, test.expectedRes, res)
			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func Test_NewVotingPowers(t *testing.T) {
	t.Parallel()

	t.Run("should return error if voting power calculation is unsuccessful", func(t *testing.T) {
		res, err := NewVotingPower(addr1, *testStakedBalance, *testDelegatedBalance, *big.NewInt(1001))

		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("should successfuly create VotingPowers with 2 elements", func(t *testing.T) {
		vPower, _ := NewVotingPower(addr1, *testStakedBalance, *testDelegatedBalance, *testdelegationWeightBps)
		vPowerTwo, _ := NewVotingPower(addr1, *testStakedBalance, *testDelegatedBalance, *testdelegationWeightBps)
		vPowers := NewVotingPowers(vPower, vPowerTwo)

		assert.Equal(
			t,
			&votingPowers{values: []VotingPower{vPower, vPowerTwo}, totalVPower: big.NewInt(30000)},
			vPowers,
		)
	})

	t.Run("should successfuly create VotingPowers with 0 elements", func(t *testing.T) {
		vPowers := NewVotingPowers()

		assert.Equal(
			t,
			&votingPowers{values: []VotingPower{}, totalVPower: big.NewInt(0)},
			vPowers,
		)
	})
}

func Test_VotingPowersAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		vPowers             VotingPowers
		newVPower           VotingPower
		expectedErr         error
		expectedVPowers     VotingPowers
		expectedTotalVPower *big.Int
	}{
		{
			name: "should return error in case of duplicated validator",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			newVPower:   newTestVotingPower(addr1),
			expectedErr: ErrVotingPowerAlreadyExists,
			expectedVPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			expectedTotalVPower: big.NewInt(15000),
		},
		{
			name: "should add Voting Power and increase total voting power",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			newVPower:   newTestVotingPower(addr2),
			expectedErr: nil,
			expectedVPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			expectedTotalVPower: big.NewInt(30000),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.ErrorIs(
				t,
				test.expectedErr,
				test.vPowers.Add(test.newVPower),
			)

			assert.Equal(
				t,
				test.expectedVPowers,
				test.vPowers,
			)

			assert.Equal(
				t,
				*test.expectedTotalVPower,
				test.vPowers.GetTVotingPower(),
			)
		})
	}
}

func Test_votingPowersIncreaseVotingPower(t *testing.T) {
	t.Parallel()

	vPower := &votingPowers{totalVPower: big.NewInt(0), values: []VotingPower{}}
	vPower.increaseVotingPower(*big.NewInt(500))

	assert.Equal(
		t,
		big.NewInt(500),
		vPower.totalVPower,
	)
}

func Test_VotingPowersAt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		vPowers        VotingPowers
		index          uint64
		expectedVPower VotingPower
	}{
		{
			name: "should return proper voting power 1",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			index:          0,
			expectedVPower: newTestVotingPower(addr1),
		},
		{
			name: "should return proper voting power 2",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			index:          1,
			expectedVPower: newTestVotingPower(addr2),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				test.expectedVPower,
				test.vPowers.At(test.index),
			)
		})
	}
}

func Test_VotingPowersIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		vPowers  VotingPowers
		addr     types.Address
		expected int64
	}{
		{
			name: "success",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			addr:     addr2,
			expected: 1,
		},
		{
			name: "not found",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			addr:     types.StringToAddress("fake"),
			expected: -1,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				test.expected,
				test.vPowers.Index(test.addr),
			)
		})
	}
}

func Test_VotingPowersIncludes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		vPowers  VotingPowers
		addr     types.Address
		expected bool
	}{
		{
			name: "success",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			addr:     addr1,
			expected: true,
		},
		{
			name: "not found",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			addr:     types.StringToAddress("fake"),
			expected: false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				test.expected,
				test.vPowers.Includes(test.addr),
			)
		})
	}
}

func Test_VotingPowersGetTVotingPower(t *testing.T) {
	t.Parallel()

	vPower := &votingPowers{totalVPower: big.NewInt(500), values: []VotingPower{}}
	vPower.GetTVotingPower()

	assert.Equal(
		t,
		*big.NewInt(500),
		vPower.GetTVotingPower(),
	)
}

func Test_VotingPowersGetVotingPower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		vPowers     VotingPowers
		addr        types.Address
		expectedRes big.Int
		expectedErr error
	}{
		{
			name: "should return error in case voting power is not found",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			addr:        addr2,
			expectedRes: *big.NewInt(0),
			expectedErr: ErrVPowerNotFound,
		},
		{
			name: "should return the voting power",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			addr:        addr2,
			expectedRes: *big.NewInt(15000),
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := test.vPowers.GetVotingPower(test.addr)

			assert.Equal(
				t,
				test.expectedRes,
				res,
			)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}

func Test_VotingPowersCalcMessagesPower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		vPowers     VotingPowers
		msgs        []*msgProto.Message
		expectedRes big.Int
		expectedErr error
	}{
		{
			name: "should return error when cannot get voting power of a specific validator",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
			),
			msgs:        []*msgProto.Message{{From: addr1[:]}, {From: addr2[:]}},
			expectedRes: *big.NewInt(0),
			expectedErr: ErrVPowerNotFound,
		},
		{
			name: "should return the total voting power of all messages' signers",
			vPowers: NewVotingPowers(
				newTestVotingPower(addr1),
				newTestVotingPower(addr2),
			),
			msgs:        []*msgProto.Message{{From: addr1[:]}, {From: addr2[:]}},
			expectedRes: *big.NewInt(30000),
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := test.vPowers.CalcMessagesPower(test.msgs)

			assert.Equal(
				t,
				test.expectedRes,
				res,
			)

			assert.ErrorIs(t, test.expectedErr, err)
		})
	}
}
