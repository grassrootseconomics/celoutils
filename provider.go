package celoutils

import (
	"math/big"
	"net/http"
	"time"

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/rpc"
	"github.com/grassrootseconomics/w3-celo"
)

const (
	slaTimeout = 5 * time.Second

	MainnetChainId int64 = 42220
	TestnetChainId int64 = 44787
)

type ProviderOpts struct {
	ChainId          int64
	RpcEndpoint      string
	CustomHTTPClient *http.Client
}

type Provider struct {
	Client  *w3.Client
	Signer  types.Signer
	ChainId int64
}

func NewProvider(o ProviderOpts) (*Provider, error) {
	if o.CustomHTTPClient == nil {
		o.CustomHTTPClient = defaultHTTPClient()
	}

	rpcClient, err := rpc.DialHTTPWithClient(o.RpcEndpoint, o.CustomHTTPClient)
	if err != nil {
		return nil, err
	}

	return &Provider{
		ChainId: o.ChainId,
		Client:  w3.NewClient(rpcClient),
		Signer:  types.LatestSignerForChainID(big.NewInt(o.ChainId)),
	}, nil
}

func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: slaTimeout,
	}
}
