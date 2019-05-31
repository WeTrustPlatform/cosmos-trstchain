package multisigservice

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the multisigservice Querier
const (
	QueryWallet = "wallet"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryWallet:
			return queryWallet(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown multisigservice query endpoint")
		}
	}
}

func queryWallet(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	wallet := keeper.GetWalletFromBech32(ctx, path[0])
	bz, err := codec.MarshalJSONIndent(keeper.cdc, wallet)
	if err != nil {
		panic("Could not marshal result to JSON")
	}
	return bz, nil
}
