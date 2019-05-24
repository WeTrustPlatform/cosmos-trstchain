package app

import (
	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// setUpRouter registers all the module's routes
func (app *TRSTChainApp) setUpRouter() {
	app.Router().
		AddRoute("bank", bank.NewHandler(app.coinKeeper)).
		AddRoute("multisigservice", multisigservice.NewHandler(app.multisigServiceKeeper))

	app.QueryRouter().
		AddRoute("account", auth.NewQuerier(app.accountKeeper))
}
