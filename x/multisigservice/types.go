package multisigservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MultisigWallet is struct that contains all the parameters of a wallet
type MultisigWallet struct {
	Creator            sdk.AccAddress   `json:"creator"`
	Owners             []sdk.AccAddress `json:"owners"`
	RequiredSignatures sdk.Int          `json:"required_signatures"`
}

// Implement fmt.Stringer to use with cliCtx.PrintOutput
func (w MultisigWallet) String() string {
	j, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	return string(j)
}
