package client

import (
	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

// NewModuleClient creates new instances of MultisigService module
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	multisigserviceTxCmd := &cobra.Command{
		Use:   "multisigservice",
		Short: "Multisigservice transactions subcommands",
	}

	multisigserviceTxCmd.AddCommand(client.PostCommands(
		cli.GetCmdCreateWallet(mc.cdc),
	)...)

	return multisigserviceTxCmd
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {

	// Group nameservice queries under a subcommand
	multisigserviceQueryCmd := &cobra.Command{
		Use:   "multisigservice",
		Short: "Querying commands for the multisigservice module",
	}

	multisigserviceQueryCmd.AddCommand(client.GetCommands(
		cli.GetCmdWallet(mc.storeKey, mc.cdc),
	)...)

	return multisigserviceQueryCmd
}
