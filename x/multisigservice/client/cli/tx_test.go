package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAddresses(t *testing.T) {
	addrs, _ := ParseAddresses("cosmos19a99vt09df3vmtml95ylc2u6n532hv8mt9drjy,cosmos19svpfaujcn5vqvgqcylwvakq9xeh7vnx8ufr00")
	require.Equal(t, "cosmos19a99vt09df3vmtml95ylc2u6n532hv8mt9drjy", addrs[0].String())
	require.Equal(t, "cosmos19svpfaujcn5vqvgqcylwvakq9xeh7vnx8ufr00", addrs[1].String())
}
