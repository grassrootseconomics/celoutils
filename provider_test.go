package celoutils

import (
	"testing"
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
					RpcEndpoint: "https://alfajores-forno.celo-testnet.org",
				},
			},
			wantErr: false,
		},
		// {
		// 	name: "Dial client with bad DSN",
		// 	args: args{
		// 		ProviderOpts{
		// 			ChainId:     TestnetChainId,
		// 			RpcEndpoint: "h:/test",
		// 		},
		// 	},
		// 	wantErr: true,
		// },
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
