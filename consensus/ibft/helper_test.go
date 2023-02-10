package ibft

import (
	"crypto/ecdsa"
	"math/big"
	"strconv"
	"testing"

	"github.com/0xPolygon/polygon-edge/crypto"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/0xPolygon/polygon-edge/validators"
)

type testerAccount struct {
	alias string
	priv  *ecdsa.PrivateKey
}

func (t *testerAccount) Address() types.Address {
	return crypto.PubKeyToAddress(&t.priv.PublicKey)
}

type testerAccountPool struct {
	t        *testing.T
	accounts []*testerAccount
}

func newTesterAccountPool(t *testing.T, num ...int) *testerAccountPool {
	t.Helper()

	pool := &testerAccountPool{
		t:        t,
		accounts: []*testerAccount{},
	}

	if len(num) == 1 {
		for i := 0; i < num[0]; i++ {
			key, _ := crypto.GenerateECDSAKey()

			pool.accounts = append(pool.accounts, &testerAccount{
				alias: strconv.Itoa(i),
				priv:  key,
			})
		}
	}

	return pool
}

func (ap *testerAccountPool) add(accounts ...string) {
	ap.t.Helper()

	for _, account := range accounts {
		if acct := ap.get(account); acct != nil {
			continue
		}

		priv, err := crypto.GenerateECDSAKey()
		if err != nil {
			panic("BUG: Failed to generate crypto key")
		}

		ap.accounts = append(ap.accounts, &testerAccount{
			alias: account,
			priv:  priv,
		})
	}
}

func (ap *testerAccountPool) get(name string) *testerAccount {
	ap.t.Helper()

	for _, i := range ap.accounts {
		if i.alias == name {
			return i
		}
	}

	return nil
}

func (ap *testerAccountPool) ValidatorSet() validators.Validators {
	ap.t.Helper()

	v := validators.NewECDSAValidatorSet()
	for _, i := range ap.accounts {
		_ = v.Add(&validators.ECDSAValidator{
			Address: i.Address(),
		})
	}

	return v
}

func (ap *testerAccountPool) VPowers() validators.VotingPowers {
	ap.t.Helper()

	v := validators.NewVotingPowers()
	for _, i := range ap.accounts {
		vPower, err := validators.NewVotingPower(i.Address(), *big.NewInt(15000), *big.NewInt(0), *big.NewInt(85))
		if err != nil {
			panic(err)
		}

		_ = v.Add(vPower)
	}

	return v
}
