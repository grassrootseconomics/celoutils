package celoutils

import (
	"math/big"
)

const (
	ErrInsufficientGas = "w3: call failed: insufficient funds for gas * price + value + gatewayFee"
	ErrGasPriceLow     = "w3: call failed: gasprice is less than gas price minimum floor"
	ErrNonceLow        = "w3: call failed: nonce too low"

	CeloContractMainnet   = "0x471EcE3750Da237f93B8E339c536989b8978a438"
	CeloContractTestnet   = "0xF194afDf50B03e69Bd7D057c1Aa9e10c9954E4C9"
	CUSDContractTestnet   = "0x874069Fa1Eb16D44d622F2e0Ca25eeA172369bC1"
	MinGasContractMainnet = "0xDfca3a8d7699D8bAfe656823AD60C17cb8270ECC"
	MinGasContractTestnet = "0xd0Bf87a5936ee17014a057143a494Dc5C5d51E5e"

	CUSDContractMainnet = "0x765DE816845861e75A25fCA122bb6898B8B1282a"
	CKESContractMainnet = "0x456a3D042C0DbD3db53D5489e98dFb038553B0d0"
	CEURContractmainnet = "0xD8763CBa276a3738E6DE85b4b3bF5FDed6D6cA73"
	USDCContractMainnet = "0xcebA9300f2b948710d2653dD7B07f33A8B32118C"
	USDTContractMainnet = "0x617f3112bf5397D0467D315cC709EF968D9ba546"

	CeloMainnet   int64 = 42220
	CeloAlfajores int64 = 44787
	CeloBaklava   int64 = 62320
)

var (
	// SafeGasLimit is set to 350k gas units.
	SafeGasLimit = 350000
	// SafeGasFeeCap is set to 2x the current MinBaseFee of 5 gwei.
	SafeGasFeeCap = big.NewInt(10000000000)
	// SafeGasTipCap is set to 5 wei.
	SafeGasTipCap = big.NewInt(5)
)
