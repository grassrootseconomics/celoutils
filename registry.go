package celoutils

import (
	"context"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/w3-celo-patch"
	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
)

const (
	TrainingVoucher = "TrainingVoucher"
	AccountIndex    = "AccountIndex"
	TokenIndex      = "TokenIndex"
	GasFaucet       = "GasFaucet"
)

// RegistryMap auto-loads all well known custodial related smart contract addresses through the CICRegistry pointer contract.
func (p *Provider) RegistryMap(ctx context.Context, registryAddress common.Address) (map[string]common.Address, error) {
	var (
		listFunc = w3.MustNewFunc("addressOf(bytes32)", "address")

		trainingVoucherAddress common.Address
		accountIndexAddress    common.Address
		tokenIndexAddress      common.Address
		gasFaucetAddress       common.Address
	)

	err := p.Client.CallCtx(
		ctx,
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(TrainingVoucher), 32))).Returns(&trainingVoucherAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(AccountIndex), 32))).Returns(&accountIndexAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(TokenIndex), 32))).Returns(&tokenIndexAddress),
		eth.CallFunc(listFunc, registryAddress, common.BytesToHash(common.RightPadBytes([]byte(GasFaucet), 32))).Returns(&gasFaucetAddress),
	)
	if err != nil {
		return nil, err
	}

	registryMap := make(map[string]common.Address)
	registryMap[TrainingVoucher] = trainingVoucherAddress
	registryMap[AccountIndex] = accountIndexAddress
	registryMap[TokenIndex] = tokenIndexAddress
	registryMap[GasFaucet] = gasFaucetAddress

	return registryMap, nil
}
