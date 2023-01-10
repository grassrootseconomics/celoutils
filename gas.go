package celo

import (
	"context"
	"math/big"

	"github.com/grassrootseconomics/w3-celo-patch"
	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
)

var (
	FixedMinGas = big.NewInt(500000001)
)

func (p *Provider) GetOptimumGasPrice(ctx context.Context) (*big.Int, error) {
	var (
		callFunc        *eth.CallFuncFactory
		optimumGasPrice big.Int
	)

	abi, err := w3.NewFunc("getGasPriceMinimum(address tokenAddress)", "uint256")
	if err != nil {
		return big.NewInt(0), nil
	}

	if p.ChainId == MainnetChainId {
		callFunc = eth.CallFunc(abi, w3.A(minGasContractMainnet), w3.A(CeloContractMainnet))
	} else {
		callFunc = eth.CallFunc(abi, w3.A(minGasContractTestnet), w3.A(CeloContractTestnet))
	}

	err = p.Client.CallCtx(
		ctx,
		callFunc.Returns(&optimumGasPrice),
	)
	if err != nil {
		return big.NewInt(0), nil
	}

	return &optimumGasPrice, nil
}
