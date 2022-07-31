package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/osmosis-labs/osmosis/v10/x/alias/types"
)

// Flag names and values
const (
	FlagAs = "as"
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

	cmd.AddCommand(
		NewExecCmd(),
	)

	return cmd
}

// NewExecCmd broadcast MsgExec
func NewExecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [msgType] [msgJson]",
		Short: "Execute a list of messages as another user",
		Args:  cobra.ExactArgs(2), // ToDo: Better arg definition?
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			sender := clientCtx.GetFromAddress()

			as, err := cmd.Flags().GetString(FlagAs)
			if err != nil {
				return err
			}

			msg := types.NewMsgExec(sender.String(), args[1], args[0], as)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagAs, "", "The addr to execute the message as")
	return cmd
}
