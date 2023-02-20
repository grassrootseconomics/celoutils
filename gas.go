package celoutils

import (
	"math/big"
)

var (
	// SafeGasFeeCap is set to 4x the current MinBaseFee of 5 gwei.
	SafeGasFeeCap = big.NewInt(20000000000)
	// SafeGasTipCap is set to 10 wei.
	SafeGasTipCap = big.NewInt(10)
)
