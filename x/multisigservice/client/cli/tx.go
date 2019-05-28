package cli

import (
	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdCreateWallet is the CLI command for creating a new wallet
func GetCmdCreateWallet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-wallet [required signatures] [owner[,owner]...]",
		Short: "Create a new multisig wallet",
		Args:  cobra.ExactValidArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			addresses, err := ParseAddresses(args[1])
			if err != nil {
				return err
			}

			msg := multisigservice.NewMsgCreateWallet(
				cliCtx.GetFromAddress(),
				addresses,
				args[0],
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}
