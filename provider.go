package celoutils

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/grassrootseconomics/w3-celo"
)

type (
	Option func(c *Provider)

	Provider struct {
		Client *w3.Client
		Signer types.Signer
	}
)

func WithSigner(signer types.Signer) Option {
	return func(p *Provider) {
		p.Signer = signer
	}
}

func WithClient(w3Client *w3.Client) Option {
	return func(p *Provider) {
		p.Client = w3Client
	}
}

func NewProvider(url string, chainID int64, opts ...Option) *Provider {
	defaultProvider := &Provider{
		Client: w3.MustDial(url),
		Signer: types.LatestSignerForChainID(big.NewInt(chainID)),
	}

	for _, opt := range opts {
		opt(defaultProvider)
	}

	return defaultProvider
}
