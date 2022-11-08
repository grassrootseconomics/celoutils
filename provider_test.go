package celo

import (
	"testing"
)

const (
	TestnetRpcEndpoint = "https://alfajores-forno.celo-testnet.org"
	MainnetRpcEndpoint = "https://forno.celo.org"
)

func TestNewProvider(t *testing.T) {
	type args struct {
		o ProviderOpts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Dial client",
			args: args{
				ProviderOpts{
					ChainId:     TestnetChainId,
					RpcEndpoint: TestnetRpcEndpoint,
				},
			},
			wantErr: false,
		},
		{
			name: "Dial client with bad DSN",
			args: args{
				ProviderOpts{
					ChainId:     TestnetChainId,
					RpcEndpoint: "h:/test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProvider(tt.args.o)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
