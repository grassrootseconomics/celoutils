package celo

import (
	"context"
	"testing"
)

func TestProvider_GetOptimumGasPrice(t *testing.T) {
	tests := []struct {
		name    string
		args    ProviderOpts
		want    uint64
		wantErr bool
	}{
		{
			name: "Positive mainnet min gas value",
			args: ProviderOpts{
				RpcEndpoint: MainnetRpcEndpoint,
				ChainId:     MainnetChainId,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Positive testnet min gas value",
			args: ProviderOpts{
				RpcEndpoint: TestnetRpcEndpoint,
				ChainId:     TestnetChainId,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewProvider(tt.args)
			if err != nil {
				t.Fatal("RPC endpoint parsing failed")
				return
			}

			got, err := p.GetOptimumGasPrice(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.GetOptimumGasPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Uint64() < tt.want {
				t.Errorf("Provider.GetOptimumGasPrice() = %v, want greater than %v", got, tt.want)
			}
		})
	}
}
