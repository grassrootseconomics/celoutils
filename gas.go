package celo

import (
	"context"
	"math/big"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

const (
	CeloContractMainnet = "0x471EcE3750Da237f93B8E339c536989b8978a438"
	CeloContractTestnet = "0xF194afDf50B03e69Bd7D057c1Aa9e10c9954E4C9"

	minGasContractMainnet = "0xDfca3a8d7699D8bAfe656823AD60C17cb8270ECC"
	minGasContractTestnet = "0xd0Bf87a5936ee17014a057143a494Dc5C5d51E5e"
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
