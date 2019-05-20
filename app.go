package app

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/gogo/protobuf/codec"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "trstchain"
)

type trstChainApp struct {
	*bam.BaseApp
}

func NewTRSTChainApp(logger log.Logger, db dbm.DB) *trstChainApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &trstChainApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app
}

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	//TODO HN implement this
	return cdc
}
