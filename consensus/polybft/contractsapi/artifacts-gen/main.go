package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/0xPolygon/polygon-edge/consensus/polybft/contractsapi/artifact"
	"github.com/0xPolygon/polygon-edge/helper/common"
)

const (
	extension = ".sol"
)

func main() {
	_, filename, _, _ := runtime.Caller(0) //nolint: dogsled
	currentPath := path.Dir(filename)
	scpath := path.Join(currentPath, "../../../../core-contracts/artifacts/contracts/")

	readContracts := []struct {
		Path string
		Name string
	}{
		{
			"child/L2StateSender.sol",
			"L2StateSender",
		},
		{
			"child/StateReceiver.sol",
			"StateReceiver",
		},
		{
			"child/System.sol",
			"System",
		},
		{
			"common/BLS.sol",
			"BLS",
		},
		{
			"common/BN256G2.sol",
			"BN256G2",
		},
		{
			"common/Merkle.sol",
			"Merkle",
		},
		{
			"root/CheckpointManager.sol",
			"CheckpointManager",
		},
		{
			"root/ExitHelper.sol",
			"ExitHelper",
		},
		{
			"root/StateSender.sol",
			"StateSender",
		},
		{
			"child/ChildValidatorSet.sol",
			"ChildValidatorSet",
		},
		{
			"child/LiquidityToken.sol",
			"LiquidityToken",
		},
		{
			"libs/GenesisProxy.sol",
			"GenesisProxy",
		},
		{
			"../@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol",
			"TransparentUpgradeableProxy",
		},
	}

	str := `// This is auto-generated file. DO NOT EDIT.
package contractsapi

`

	for _, v := range readContracts {
		artifactBytes, err := artifact.ReadArtifactData(scpath, v.Path, getContractName(v.Path))
		if err != nil {
			log.Fatal(err)
		}

		dst := &bytes.Buffer{}
		if err = json.Compact(dst, artifactBytes); err != nil {
			log.Fatal(err)
		}

		str += fmt.Sprintf("var %sArtifact string = `%s`\n", v.Name, dst.String())
	}

	output, err := format.Source([]byte(str))
	if err != nil {
		fmt.Println(str)
		log.Fatal(err)
	}

	if err = common.SaveFileSafe(currentPath+"/../gen_sc_data.go", output, 0600); err != nil {
		log.Fatal(err)
	}
}

// getContractName extracts smart contract name from provided path
func getContractName(path string) string {
	pathSegments := strings.Split(path, string([]rune{os.PathSeparator}))
	nameSegment := pathSegments[len(pathSegments)-1]

	return strings.Split(nameSegment, extension)[0]
}
