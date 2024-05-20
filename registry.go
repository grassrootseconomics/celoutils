package celoutils

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

const (
	AccountIndex    = "AccountIndex"
	CustodialProxy  = "CustodialRegistrationProxy"
	GasFaucet       = "GasFaucet"
	TrainingVoucher = "TrainingVoucher"
	TokenIndex      = "TokenIndex"
	PoolIndex       = "PoolIndex"
)

// RegistryMap auto-loads all well known custodial related smart contract addresses through the CICRegistry pointer contract.
func (p *Provider) RegistryMap(ctx context.Context, registryAddress common.Address) (map[string]common.Address, error) {
	var (
		listFunc = w3.MustNewFunc("addressOf(bytes32)", "address")

		accountIndexAddress    common.Address
		custodialProxy         common.Address
		gasFaucetAddress       common.Address
		tokenIndexAddress      common.Address
		trainingVoucherAddress common.Address
		poolIndex              common.Address
	)

	err := p.Client.CallCtx(
		ctx,
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(AccountIndex), 32))).Returns(&accountIndexAddress),
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(CustodialProxy), 32))).Returns(&custodialProxy),
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(GasFaucet), 32))).Returns(&gasFaucetAddress),
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(TokenIndex), 32))).Returns(&tokenIndexAddress),
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(TrainingVoucher), 32))).Returns(&trainingVoucherAddress),
		eth.CallFunc(registryAddress, listFunc, common.BytesToHash(common.RightPadBytes([]byte(PoolIndex), 32))).Returns(&poolIndex),
	)
	if err != nil {
		return nil, err
	}

	registryMap := make(map[string]common.Address)
	registryMap[CustodialProxy] = custodialProxy
	registryMap[AccountIndex] = accountIndexAddress
	registryMap[GasFaucet] = gasFaucetAddress
	registryMap[TokenIndex] = tokenIndexAddress
	registryMap[TrainingVoucher] = trainingVoucherAddress
	registryMap[PoolIndex] = poolIndex

	return registryMap, nil
}
