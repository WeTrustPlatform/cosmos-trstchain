package multisigservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

func (k Keeper) AddOwner(ctx sdk.Context, walletAddress sdk.AccAddress, owner sdk.AccAddress) {
	//TODO Implement this
}

func (k Keeper) RemoveOwner(ctx sdk.Context, walletAddress sdk.AccAddress, owner sdk.AccAddress) {
	//TODO Implement this
}

func (k Keeper) Send(ctx sdk.Context, walletAddress sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) {
	//TODO Implement this
}

func (k Keeper) SetRequiredSignatures(ctx sdk.Context, walletAddress sdk.AccAddress, num sdk.Int) {
	//TODO Implement this
}

func (k Keeper) CreateWallet(ctx sdk.Context, creator sdk.AccAddress, owners []sdk.AccAddress, requiredSignatures sdk.Int) {
	//TODO Implement this
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}
