package cli

import (
	"fmt"

	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetCmdWallet queries information about a wallet
func GetCmdWallet(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "wallet [bech32]",
		Short: "Get wallet info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/wallet/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not find wallet - %s \n", string(addr))
				return nil
			}

			var out multisigservice.MultisigWallet
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
