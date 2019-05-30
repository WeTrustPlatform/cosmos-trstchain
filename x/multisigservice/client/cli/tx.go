package cli

import (
	"errors"
	"strings"

	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdCreateWallet is the CLI command for creating a new wallet
func GetCmdCreateWallet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-wallet [required signatures] [owner[,owner]...]",
		Short: "Create a new multisig wallet",
		Args:  cobra.ExactArgs(2),
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

			requiredSignatures, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return errors.New("Cannot parse requiredSignatures to sdk.Int")
			}

			msg := multisigservice.NewMsgCreateWallet(
				cliCtx.GetFromAddress(),
				addresses,
				requiredSignatures,
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

// ParseAddresses returns array of AccAddresses from the CLI arg
func ParseAddresses(input string) ([]sdk.AccAddress, error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return nil, errors.New("Owners array cannot be empty")
	}

	addrStrs := strings.Split(input, ",")
	accAddresses := make([]sdk.AccAddress, len(addrStrs))
	for i, addrStr := range addrStrs {
		accAddr, err := types.AccAddressFromBech32(addrStr)
		if err != nil {
			panic(err)
		}
		accAddresses[i] = accAddr
	}

	return accAddresses, nil
}
