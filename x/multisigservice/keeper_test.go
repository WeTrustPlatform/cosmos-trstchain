package multisigservice

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func getAccAddressFromBech32(addr string) (accAddr types.AccAddress) {
	var err error
	accAddr, err = types.AccAddressFromBech32(addr)
	if err != nil {
		panic(err)
	}
	return accAddr
}

func TestDeriveAccAddressSuccess(t *testing.T) {
	addr := getAccAddressFromBech32("cosmos19a99vt09df3vmtml95ylc2u6n532hv8mt9drjy")

	actual := DeriveAccAddress(addr, 0)
	require.Equal(t, getAccAddressFromBech32("cosmos19svpfaujcn5vqvgqcylwvakq9xeh7vnx8ufr00"), actual)

	actual = DeriveAccAddress(addr, 1)
	require.Equal(t, getAccAddressFromBech32("cosmos1kkj38xqjeezyvk3cyghjldh408aryumw2n6jey"), actual)
}
