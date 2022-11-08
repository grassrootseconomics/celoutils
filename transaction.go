package celo

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractExecutionTxOpts struct {
	ContractAddress common.Address
	InputData       []byte
	GasPrice        big.Int
	GasLimit        uint64
	Nonce           uint64
}

type GasTransferTxOpts struct {
	To       common.Address
	Value    big.Int
	GasPrice big.Int
	Nonce    uint64
}

func (p *Provider) SignContractExecutionTx(privateKey *ecdsa.PrivateKey, txData ContractExecutionTxOpts) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.LegacyTx{
		To:       &txData.ContractAddress,
		Nonce:    txData.Nonce,
		Data:     txData.InputData,
		Gas:      txData.GasLimit,
		GasPrice: &txData.GasPrice,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *Provider) SignGasTransferTx(privateKey *ecdsa.PrivateKey, txData GasTransferTxOpts) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.LegacyTx{
		Value:    &txData.Value,
		To:       &txData.To,
		Nonce:    txData.Nonce,
		Gas:      21000,
		GasPrice: &txData.GasPrice,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
