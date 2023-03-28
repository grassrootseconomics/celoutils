package celoutils

import (
	"context"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/w3-celo-patch"
	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
)

const (
	AccountIndex    = "AccountIndex"
	CustodialProxy  = "CustodialRegistrationProxy"
	GasFaucet       = "GasFaucet"
	TrainingVoucher = "TrainingVoucher"
	TokenIndex      = "TokenIndex"
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
	)

	err := p.Client.CallCtx(
		ctx,
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(AccountIndex), 32))).Returns(&accountIndexAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(CustodialProxy), 32))).Returns(&custodialProxy),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(GasFaucet), 32))).Returns(&gasFaucetAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(TokenIndex), 32))).Returns(&tokenIndexAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(TrainingVoucher), 32))).Returns(&trainingVoucherAddress),
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

	return registryMap, nil
}
