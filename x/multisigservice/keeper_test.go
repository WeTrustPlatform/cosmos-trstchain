package multisigservice

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDeriveAccAddressSuccess(t *testing.T) {
	var addr, expected, actual types.AccAddress
	var err error
	addr, err = types.AccAddressFromBech32("cosmos19a99vt09df3vmtml95ylc2u6n532hv8mt9drjy")
	if err != nil {
		panic(err)
	}
	expected, err = types.AccAddressFromBech32("cosmos1kkj38xqjeezyvk3cyghjldh408aryumw2n6jey")
	if err != nil {
		panic(err)
	}

	actual = DeriveAccAddress(addr, 1)
	require.Equal(t, expected, actual)
}
