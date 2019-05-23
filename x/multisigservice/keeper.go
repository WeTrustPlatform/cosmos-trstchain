package multisigservice

import (
	"crypto/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/crypto/ed25519"

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

func (k Keeper) CreateWallet(ctx sdk.Context, creator sdk.AccAddress, owners []sdk.AccAddress, requiredSignatures sdk.Int) sdk.AccAddress {
	if creator.Empty() {
		return nil
	}

	var pub ed25519.PubKeyEd25519
	// TODO Cannot randomly generate new address here!
	rand.Read(pub[:])
	walletAddress := sdk.AccAddress(pub.Address())

	multisigWallet := MultisigWallet{
		Creator:            creator,
		Owners:             owners,
		RequiredSignatures: requiredSignatures,
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(walletAddress, k.cdc.MustMarshalBinaryBare(multisigWallet))

	return walletAddress
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}
