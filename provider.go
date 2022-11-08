package celo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
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
		Signer:  types.NewEIP155Signer(big.NewInt(o.ChainId)),
	}, nil
}
