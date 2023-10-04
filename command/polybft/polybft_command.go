package polybft

import (
	// H_MODIFY: Registration module is moved to sidechain
	"github.com/0xPolygon/polygon-edge/command/sidechain/registration"
	"github.com/0xPolygon/polygon-edge/command/sidechain/whitelist"

	"github.com/0xPolygon/polygon-edge/command/sidechain/rewards"
	"github.com/0xPolygon/polygon-edge/command/sidechain/unstaking"
	sidechainWithdraw "github.com/0xPolygon/polygon-edge/command/sidechain/withdraw"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "polybft",
		Short: "Polybft command",
	}

	// Hydra modification: modify sidechain commands and remove rootchain commands
	polybftCmd.AddCommand(
		// sidechain (validator set) command to unstake on child chain
		unstaking.GetCommand(),
		// sidechain (validator set) command to withdraw stake on child chain
		sidechainWithdraw.GetCommand(),
		// sidechain (reward pool) command to withdraw pending rewards
		rewards.GetCommand(),
		// sidechain (validator set) command to register validator
		registration.GetCommand(),
		// sidechain (validator set) command to whitelist validators
		whitelist.GetCommand(),
	)

	return polybftCmd
}
