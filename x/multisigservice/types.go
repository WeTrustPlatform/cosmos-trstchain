package multisigservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MultisigWallet is struct that contains all the parameters of a wallet
type MultisigWallet struct {
	Creator            sdk.AccAddress   `json:"creator"`
	Owners             []sdk.AccAddress `json:"owners"`
	RequiredSignatures sdk.Int          `json:"required_signatures"`
}
