package celoutils

import (
	"math/big"
)

var (
	// SafeGasLimit is set to 350k gas units.
	SafeGasLimit = 550000
	// SafeGasFeeCap is set to 2x the current MinBaseFee of 5 gwei.
	SafeGasFeeCap = big.NewInt(10000000000)
	// SafeGasTipCap is set to 5 wei.
	SafeGasTipCap = big.NewInt(5)
)
