package app

import (
	"github.com/WeTrustPlatform/cosmos-trstchain/x/multisigservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "trstchain"
)

// TRSTChainApp defines all the parameters needed for the app
type TRSTChainApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// sdk keepers

	// AccountKeeper handles address -> account lookups
	keyAccount    *sdk.KVStoreKey
	accountKeeper auth.AccountKeeper

	// FeeCollectionKeeper collects transaction fees and renders them to the fee distribution module
	keyFeeCollection    *sdk.KVStoreKey
	feeCollectionKeeper auth.FeeCollectionKeeper

	// ParamsKeeper handles parameter storage for the application
	keyParams    *sdk.KVStoreKey
	tkeyParams   *sdk.TransientStoreKey
	paramsKeeper params.Keeper

	// CoinKeeper allows you perform sdk.Coins interactions
	coinKeeper bank.Keeper

	// app-specific keepers

	keyMain *sdk.KVStoreKey

	// MultisigServiceKeeper stores multsig wallets info
	keyMultisigService    *sdk.KVStoreKey
	multisigServiceKeeper multisigservice.Keeper
}

// NewTRSTChainApp creates new instances of TRSTChain
func NewTRSTChainApp(logger log.Logger, db dbm.DB) *TRSTChainApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &TRSTChainApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:            sdk.NewKVStoreKey("main"),
		keyMultisigService: sdk.NewKVStoreKey("multisig_service"),

		keyAccount:       sdk.NewKVStoreKey("account"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),
		keyParams:        sdk.NewKVStoreKey("params"),
		tkeyParams:       sdk.NewTransientStoreKey("transient_params"),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	app.coinKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	app.multisigServiceKeeper = multisigservice.NewKeeper(
		app.cdc,
		app.coinKeeper,
		app.accountKeeper,
		app.keyMultisigService,
	)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.setUpRouter()

	app.SetInitChainer(app.initChainer)

	app.MountStores(
		app.keyMain,
		app.keyMultisigService,
		app.keyAccount,
		app.keyFeeCollection,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}
	return app
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	multisigservice.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
