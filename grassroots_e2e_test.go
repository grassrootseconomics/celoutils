package celoutils

import (
	"context"
	"testing"
)

func TestGrassroots_RegistryMap(t *testing.T) {
	p := NewProvider("https://api.tatum.io/v3/blockchain/node/celo-mainnet", CeloMainnet)

	rMap, err := p.RegistryMap(context.Background(), SarafuNetworkRegistry)
	if err != nil {
		t.Error(err)
	}
	t.Log(rMap)
}

func TestGrassroots_GetGESmartContracts(t *testing.T) {
	p := NewProvider("https://api.tatum.io/v3/blockchain/node/celo-mainnet", CeloMainnet)

	aMap, err := p.GetGESmartContracts(context.Background(), []string{
		SarafuNetworkRegistry.Hex(),
		CustodialRegistry.Hex(),
	})
	if err != nil {
		t.Error(err)
	}
	if len(aMap) < 100 {
		t.Error("partial ge smart contracts retrieved")
	}
	t.Log(aMap)
}
