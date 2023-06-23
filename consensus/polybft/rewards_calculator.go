package polybft

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/hashicorp/go-hclog"
)

var (
	denominator = big.NewInt(10000)
)

type RewardsCalculator interface {
	GetMaxReward(block *types.Header) (*big.Int, error)
}

type rewardsCalculator struct {
	logger     hclog.Logger
	blockchain blockchainBackend
}

func NewRewardsCalculator(logger hclog.Logger, blockchain blockchainBackend) RewardsCalculator {
	return &rewardsCalculator{
		logger:     logger,
		blockchain: blockchain,
	}
}

func (r *rewardsCalculator) GetMaxReward(block *types.Header) (*big.Int, error) {
	stakedBalance, err := r.getStakedBalance(block)
	if err != nil {
		return nil, err
	}

	baseReward, err := r.getMaxBaseReward(block)
	if err != nil {
		return nil, err
	}

	vestingBonus, err := r.getVestingBonus(52)
	if err != nil {
		return nil, err
	}

	rsiBonus, err := r.getMaxRSIBonus()
	if err != nil {
		return nil, err
	}

	macroFactor, err := r.getMaxMacroFactor()
	if err != nil {
		return nil, err
	}

	reward := calcMaxReward(stakedBalance, baseReward.Numerator, vestingBonus, rsiBonus, macroFactor)

	return reward, nil
}

func (r *rewardsCalculator) getStakedBalance(block *types.Header) (*big.Int, error) {
	provider, err := r.blockchain.GetStateProviderForBlock(block)
	if err != nil {
		return nil, err
	}

	systemState := r.blockchain.GetSystemState(provider)
	reward, err := systemState.GetStakedBalance()
	if err != nil {
		return nil, err
	}

	return reward, nil
}

func (r *rewardsCalculator) getMaxBaseReward(block *types.Header) (*BigNumDecimal, error) {
	provider, err := r.blockchain.GetStateProviderForBlock(block)
	if err != nil {
		return nil, err
	}

	systemState := r.blockchain.GetSystemState(provider)
	reward, err := systemState.GetBaseReward()
	if err != nil {
		return nil, err
	}

	return reward, nil
}

func (r *rewardsCalculator) getVestingBonus(vestingWeeks uint64) (*big.Int, error) {
	numerator := big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(int64(vestingWeeks)))

	return numerator, nil
}

func (r *rewardsCalculator) getMaxRSIBonus() (*big.Int, error) {
	numerator := big.NewInt(15000)

	return numerator, nil
}

func (r *rewardsCalculator) getMaxMacroFactor() (*big.Int, error) {
	numerator := big.NewInt(10000)

	return numerator, nil
}

func calcMaxReward(staked *big.Int, base *big.Int, vesting *big.Int, rsi *big.Int, macro *big.Int) *big.Int {
	res := big.NewInt(0)
	denSum := big.NewInt(0).Mul(denominator, denominator)
	denSum.Mul(denSum, denominator)

	return res.Div(res.Mul(staked, res.Mul(macro, res.Mul(rsi, res.Add(base, vesting)))), denSum)
}
