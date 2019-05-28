package multisigservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateWallet defines a CreateWallet message
type MsgCreateWallet struct {
	Creator            sdk.AccAddress
	Owners             []sdk.AccAddress
	RequiredSignatures sdk.Int
}

// Route should return the name of the module
func (msg MsgCreateWallet) Route() string {
	return "multisigservice"
}

// Type should return the action
func (msg MsgCreateWallet) Type() string {
	return "createwallet"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateWallet) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(msg.Creator.String())
	}

	if msg.RequiredSignatures.LT(sdk.OneInt()) {
		return sdk.ErrUnknownRequest("RequiredSignatures must be equal or greater than 1.")
	}

	if msg.RequiredSignatures.GT(sdk.NewInt(int64(len(msg.Owners)))) {
		return sdk.ErrUnknownRequest("RequiredSignatures must be equal or less than Owners.length.")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateWallet) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required which is the Creator
func (msg MsgCreateWallet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// NewMsgCreateWallet creates new instances of MsgCreateWallet
func NewMsgCreateWallet(creator sdk.AccAddress, owners []sdk.AccAddress, requiredSignatures sdk.Int) MsgCreateWallet {
	return MsgCreateWallet{
		Creator:            creator,
		Owners:             owners,
		RequiredSignatures: requiredSignatures,
	}
}
