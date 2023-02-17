package ibft

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_FaultyNodes(t *testing.T) {
	cases := []struct {
		Network, Faulty uint64
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{5, 1},
		{6, 1},
		{7, 2},
		{8, 2},
		{9, 2},
	}
	for _, c := range cases {
		pool := newTesterAccountPool(t, int(c.Network))
		vals := pool.ValidatorSet()
		assert.Equal(t, CalcMaxFaultyNodes(vals), int(c.Faulty))
	}
}

// TestNumValid checks if the quorum size is calculated
// correctly based on number of validators (network size).
func TestNumValid(t *testing.T) {
	cases := []struct {
		Network uint64
		Quorum  big.Int
	}{
		{1, *big.NewInt(15000)},
		{2, *big.NewInt(30000)},
		{3, *big.NewInt(45000)},
		{4, *big.NewInt(36840)},
		{5, *big.NewInt(46050)},
		{6, *big.NewInt(55260)},
		{7, *big.NewInt(64470)},
		{8, *big.NewInt(73680)},
		{9, *big.NewInt(82890)},
	}

	addAccounts := func(
		pool *testerAccountPool,
		numAccounts int,
	) {
		// add accounts
		for i := 0; i < numAccounts; i++ {
			pool.add(strconv.Itoa(i))
		}
	}

	for _, c := range cases {
		pool := newTesterAccountPool(t, int(c.Network))
		addAccounts(pool, int(c.Network))

		assert.Equal(t,
			c.Quorum,
			OptimalQuorumSize(pool.ValidatorSet(), pool.VPowers()),
		)
	}
}
