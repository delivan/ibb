package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdCreateNft())
	cmd.AddCommand(CmdUpdateNft())
	cmd.AddCommand(CmdDeleteNft())

	cmd.AddCommand(CmdCreateRepay())
	cmd.AddCommand(CmdUpdateRepay())
	cmd.AddCommand(CmdDeleteRepay())

	cmd.AddCommand(CmdCreateWithdraw())
	cmd.AddCommand(CmdUpdateWithdraw())
	cmd.AddCommand(CmdDeleteWithdraw())

	cmd.AddCommand(CmdCreateUser())
	cmd.AddCommand(CmdDeleteUser())

	cmd.AddCommand(CmdCreateBorrow())
	cmd.AddCommand(CmdUpdateBorrow())
	cmd.AddCommand(CmdDeleteBorrow())

	cmd.AddCommand(CmdCreateDeposit())
	cmd.AddCommand(CmdUpdateDeposit())
	cmd.AddCommand(CmdDeleteDeposit())

	cmd.AddCommand(CmdCreatePool())
	cmd.AddCommand(CmdDeletePool())
	cmd.AddCommand(CmdCreateAllPool())

	return cmd
}
