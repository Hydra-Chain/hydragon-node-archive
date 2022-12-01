package genesis

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/0xPolygon/polygon-edge/chain"
	"github.com/0xPolygon/polygon-edge/command"
	"github.com/0xPolygon/polygon-edge/command/helper"

	rootchain "github.com/0xPolygon/polygon-edge/command/rootchain/helper"
	"github.com/0xPolygon/polygon-edge/consensus/polybft"
	"github.com/0xPolygon/polygon-edge/consensus/polybft/bitmap"
	bls "github.com/0xPolygon/polygon-edge/consensus/polybft/signer"
	"github.com/0xPolygon/polygon-edge/contracts"
	"github.com/0xPolygon/polygon-edge/server"
	"github.com/0xPolygon/polygon-edge/types"
)

const (
	premineValidatorsFlag          = "premine-validators"
	polyBftValidatorPrefixPathFlag = "validator-prefix"
	smartContractsRootPathFlag     = "contracts-path"

	validatorSetSizeFlag = "validator-set-size"
	sprintSizeFlag       = "sprint-size"
	blockTimeFlag        = "block-time"
	validatorsFlag       = "polybft-validators"
	bridgeFlag           = "bridge"

	defaultEpochSize                  = uint64(10)
	defaultSprintSize                 = uint64(5)
	defaultValidatorSetSize           = 100
	defaultBlockTime                  = 2 * time.Second
	defaultPolyBftValidatorPrefixPath = "test-chain-"
	defaultBridge                     = false

	bootnodePortStart = 30301
)

func (p *genesisParams) generatePolyBFTConfig() (*chain.Chain, error) {
	validatorsInfo, err := ReadValidatorsByRegexp(path.Dir(p.genesisPath), p.polyBftValidatorPrefixPath)
	if err != nil {
		return nil, err
	}

	allocs, err := p.deployContracts()
	if err != nil {
		return nil, err
	}

	// use 1st account as governance address
	governanceAccount := validatorsInfo[0].Account
	polyBftConfig := &polybft.PolyBFTConfig{
		BlockTime:         p.blockTime,
		EpochSize:         p.epochSize,
		SprintSize:        p.sprintSize,
		ValidatorSetSize:  p.validatorSetSize,
		ValidatorSetAddr:  contracts.ValidatorSetContract,
		StateReceiverAddr: contracts.StateReceiverContract,
		Governance:        types.Address(governanceAccount.Ecdsa.Address()),
	}

	if p.bridgeEnabled {
		ip, err := rootchain.ReadRootchainIP()
		if err != nil {
			return nil, err
		}

		polyBftConfig.Bridge = &polybft.BridgeConfig{
			BridgeAddr:      rootchain.StateSenderAddress,
			CheckpointAddr:  rootchain.CheckpointManagerAddress,
			JSONRPCEndpoint: ip,
		}
	}

	chainConfig := &chain.Chain{
		Name: p.name,
		Params: &chain.Params{
			ChainID: int(p.chainID),
			Forks:   chain.AllForksEnabled,
			Engine: map[string]interface{}{
				string(server.PolyBFTConsensus): polyBftConfig,
			},
		},
		Bootnodes: p.bootnodes,
	}

	// set generic validators as bootnodes if needed
	if len(p.bootnodes) == 0 {
		for i, validator := range validatorsInfo {
			bnode := fmt.Sprintf("/ip4/%s/tcp/%d/p2p/%s", "127.0.0.1", bootnodePortStart+i, validator.NodeID)
			chainConfig.Bootnodes = append(chainConfig.Bootnodes, bnode)
		}
	}

	var (
		validatorPreminesMap map[types.Address]int
		premineInfos         []*premineInfo
	)

	if p.premineValidators != "" {
		validatorPreminesMap = make(map[types.Address]int, len(validatorsInfo))

		for i, vi := range validatorsInfo {
			premineInfo, err := parsePremineInfo(fmt.Sprintf("%s:%s",
				vi.Account.Ecdsa.Address().String(), p.premineValidators))
			if err != nil {
				return nil, err
			}

			premineInfos = append(premineInfos, premineInfo)
			validatorPreminesMap[premineInfo.address] = i
		}
	}

	if len(p.premine) > 0 {
		for _, premine := range p.premine {
			premineInfo, err := parsePremineInfo(premine)
			if err != nil {
				return nil, err
			}

			if i, ok := validatorPreminesMap[premineInfo.address]; ok {
				premineInfos[i] = premineInfo
			} else {
				premineInfos = append(premineInfos, premineInfo)
			}
		}
	}

	// premine accounts
	fillPremineMap(allocs, premineInfos)

	// set initial validator set
	genesisValidators, err := p.getGenesisValidators(validatorsInfo, allocs)
	if err != nil {
		return nil, err
	}

	polyBftConfig.InitialValidatorSet = genesisValidators

	pubKeys := make([]*bls.PublicKey, len(validatorsInfo))
	for i, validatorInfo := range validatorsInfo {
		pubKeys[i] = validatorInfo.Account.Bls.PublicKey()
	}

	genesisExtraData, err := generateExtraDataPolyBft(genesisValidators, pubKeys)
	if err != nil {
		return nil, err
	}

	// populate genesis parameters
	chainConfig.Genesis = &chain.Genesis{
		GasLimit:   p.blockGasLimit,
		Difficulty: 0,
		Alloc:      allocs,
		ExtraData:  genesisExtraData,
		GasUsed:    command.DefaultGenesisGasUsed,
		Mixhash:    polybft.PolyBFTMixDigest,
	}

	return chainConfig, nil
}

func (p *genesisParams) getGenesisValidators(validators []GenesisTarget,
	allocs map[types.Address]*chain.GenesisAccount) ([]*polybft.Validator, error) {
	result := make([]*polybft.Validator, 0)

	if len(p.validators) > 0 {
		for _, validator := range p.validators {
			parts := strings.Split(validator, ":")
			if len(parts) != 2 || len(parts[0]) != 32 || len(parts[1]) < 2 {
				continue
			}

			addr := types.StringToAddress(parts[0])

			balance, err := chain.GetGenesisAccountBalance(addr, allocs)
			if err != nil {
				return nil, err
			}

			result = append(result, &polybft.Validator{
				Address: addr,
				BlsKey:  parts[1],
				Balance: balance,
			})
		}
	} else {
		for _, validator := range validators {
			pubKeyMarshalled := validator.Account.Bls.PublicKey().Marshal()
			addr := types.Address(validator.Account.Ecdsa.Address())

			balance, err := chain.GetGenesisAccountBalance(addr, allocs)
			if err != nil {
				return nil, err
			}

			result = append(result, &polybft.Validator{
				Address: addr,
				BlsKey:  hex.EncodeToString(pubKeyMarshalled),
				Balance: balance,
			})
		}
	}

	return result, nil
}

func (p *genesisParams) generatePolyBftGenesis() error {
	config, err := params.generatePolyBFTConfig()
	if err != nil {
		return err
	}

	return helper.WriteGenesisConfigToDisk(config, params.genesisPath)
}

func (p *genesisParams) deployContracts() (map[types.Address]*chain.GenesisAccount, error) {
	genesisContracts := []struct {
		name         string
		relativePath string
		address      types.Address
	}{
		{
			// Validator contract
			name:         "ChildValidatorSet",
			relativePath: "child/ChildValidatorSet.sol",
			address:      contracts.ValidatorSetContract,
		},
		{
			// State receiver contract
			name:         "StateReceiver",
			relativePath: "child/StateReceiver.sol",
			address:      contracts.StateReceiverContract,
		},
		{
			// Native Token contract (Matic ERC-20)
			name:         "MRC20",
			relativePath: "child/MRC20.sol",
			address:      contracts.NativeTokenContract,
		},
		{
			// BLS contract
			name:         "BLS",
			relativePath: "common/BLS.sol",
			address:      contracts.BLSContract,
		},
		{
			// Merkle contract
			name:         "Merkle",
			relativePath: "common/Merkle.sol",
			address:      contracts.MerkleContract,
		},
	}

	allocations := make(map[types.Address]*chain.GenesisAccount, len(genesisContracts))

	for _, contract := range genesisContracts {
		artifact, err := polybft.ReadArtifact(p.smartContractsRootPath, contract.relativePath, contract.name)
		if err != nil {
			return nil, err
		}

		allocations[contract.address] = &chain.GenesisAccount{
			Balance: big.NewInt(0),
			Code:    artifact.DeployedBytecode,
		}
	}

	return allocations, nil
}

// generateExtraDataPolyBft populates Extra with specific fields required for polybft consensus protocol
func generateExtraDataPolyBft(validators []*polybft.Validator, publicKeys []*bls.PublicKey) ([]byte, error) {
	if len(validators) != len(publicKeys) {
		return nil, fmt.Errorf("expected same length for genesis validators and BLS public keys")
	}

	delta := &polybft.ValidatorSetDelta{
		Added:   make(polybft.AccountSet, len(validators)),
		Removed: bitmap.Bitmap{},
	}

	for i, validator := range validators {
		delta.Added[i] = &polybft.ValidatorMetadata{
			Address:     validator.Address,
			BlsKey:      publicKeys[i],
			VotingPower: chain.ConvertWeiToTokensAmount(validator.Balance).Uint64(),
		}
	}

	// Order validators based on its addresses
	sort.Slice(delta.Added, func(i, j int) bool {
		return bytes.Compare(delta.Added[i].Address[:], delta.Added[j].Address[:]) < 0
	})

	extra := polybft.Extra{Validators: delta, Checkpoint: &polybft.CheckpointData{}}

	return append(make([]byte, polybft.ExtraVanity), extra.MarshalRLPTo(nil)...), nil
}