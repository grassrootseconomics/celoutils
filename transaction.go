package celoutils

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
)

type ContractExecutionTxOpts struct {
	ContractAddress common.Address
	InputData       []byte
	GasFeeCap       *big.Int
	GasTipCap       *big.Int
	GasLimit        uint64
	Nonce           uint64
}

type GasTransferTxOpts struct {
	To        common.Address
	Value     *big.Int
	GasFeeCap *big.Int
	GasTipCap *big.Int
	Nonce     uint64
}

func (p *Provider) SignContractExecutionTx(privateKey *ecdsa.PrivateKey, txData ContractExecutionTxOpts) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.CeloDynamicFeeTx{
		To:        &txData.ContractAddress,
		Nonce:     txData.Nonce,
		Data:      txData.InputData,
		Gas:       txData.GasLimit,
		GasFeeCap: txData.GasFeeCap,
		GasTipCap: txData.GasTipCap,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *Provider) SignGasTransferTx(privateKey *ecdsa.PrivateKey, txData GasTransferTxOpts) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.CeloDynamicFeeTx{
		Value:     txData.Value,
		To:        &txData.To,
		Nonce:     txData.Nonce,
		Gas:       21000,
		GasFeeCap: txData.GasFeeCap,
		GasTipCap: txData.GasTipCap,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
