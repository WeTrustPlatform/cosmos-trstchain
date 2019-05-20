package multisigservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MultisigWallet struct {
	Balance            sdk.Coins        `json:"balance"`
	Creator            sdk.AccAddress   `json:"creator"`
	Owners             []sdk.AccAddress `json:"owners"`
	RequiredSignatures sdk.Int          `json:"required_signatures"`
	Address            sdk.AccAddress   `json:"address"`
}
