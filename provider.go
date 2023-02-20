package celo

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/grassrootseconomics/w3-celo-patch"
)

const (
	MainnetChainId int64 = 42220
	TestnetChainId int64 = 44787
)

type ProviderOpts struct {
	ChainId     int64
	RpcEndpoint string
}

type Provider struct {
	Client  *w3.Client
	Signer  types.Signer
	ChainId int64
}

func NewProvider(o ProviderOpts) (*Provider, error) {
	client, err := w3.Dial(o.RpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &Provider{
		ChainId: o.ChainId,
		Client:  client,
		Signer:  types.NewLondonSigner(big.NewInt(o.ChainId)),
	}, nil
}
