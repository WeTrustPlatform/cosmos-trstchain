package multisigservice

import (
	"crypto/sha256"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ethereum/go-ethereum/rlp"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	coinKeeper bank.Keeper

	accountKeeper auth.AccountKeeper

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

// CreateWallet creates a new multisig wallet and returns a generated AccAddress
func (k Keeper) CreateWallet(ctx sdk.Context, creator sdk.AccAddress, owners []sdk.AccAddress, requiredSignatures sdk.Int) sdk.AccAddress {
	if creator.Empty() {
		return nil
	}

	// If creator's acccount does not exist, sequence will be 0
	sequence, _ := k.accountKeeper.GetSequence(ctx, creator)

	walletAddress := DeriveAccAddress(creator, sequence)

	newAccount := k.accountKeeper.NewAccountWithAddress(ctx, walletAddress)

	if !walletAddress.Equals(newAccount.GetAddress()) {
		panic(sdk.ErrUnknownAddress(
			fmt.Sprintf(
				"auth.AccountKeeper generates wrong address. Expected %s. Got %s.",
				walletAddress,
				newAccount.GetAddress(),
			)))
	}

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

// Helpers

// DeriveAccAddress generates new AccAddress from a base address and a sequence
// deterministically
func DeriveAccAddress(creator sdk.AccAddress, sequence uint64) sdk.AccAddress {
	encoded, err := rlp.EncodeToBytes([]interface{}{creator, sequence})

	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(encoded)
	// use the last 20 bytes per Ethereum specs
	return sdk.AccAddress(hashed[12:])
}
