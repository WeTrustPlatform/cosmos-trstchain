package multisigservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for multisigservice type messages
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateWallet:
			return handleMsgCreateWallet(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized multisigservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreateWallet(ctx sdk.Context, keeper Keeper, msg MsgCreateWallet) sdk.Result {
	//TODO Implement this
	return sdk.Result{}
}
