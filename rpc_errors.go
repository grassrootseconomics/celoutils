package celo

const (
	ErrInsufficientGas = "w3: call failed: insufficient funds for gas * price + value + gatewayFee"
	ErrGasPriceLow     = "w3: call failed: gasprice is less than gas price minimum floor"
	ErrNonceLow        = "w3: call failed: nonce too low"
)
