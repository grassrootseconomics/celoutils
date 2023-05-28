package celoutils

import (
	"math/big"
)

var (
	// SafeGasLimit is set to 1M gas units.
	SafeGasLimit = 1000000
	// SafeGasFeeCap is set to 4x the current MinBaseFee of 5 gwei.
	SafeGasFeeCap = big.NewInt(20000000000)
	// SafeGasTipCap is set to 10 wei.
	SafeGasTipCap = big.NewInt(10)
)
