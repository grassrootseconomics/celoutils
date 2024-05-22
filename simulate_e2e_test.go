package celoutils

import (
	"context"
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/w3-celo"
)

func TestProvider_SimulateRevertedTx(t *testing.T) {
	p := NewProvider("https://forno.celo.org", CeloMainnet)

	type args struct {
		txHash      common.Hash
		blockNumber *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Reverted reason",
			args: args{
				txHash:      w3.H("0x0fd502e9259424ebea2cd0c170adcf478bf32c2a575395649e43954d3b70d071"),
				blockNumber: big.NewInt(25736867),
			},
			want:    "ERC20: transfer amount exceeds balance",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.SimulateRevertedTx(context.Background(), tt.args.txHash, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SimulateRevertedTx() error = %v, wantErr %v, got = %s", err, tt.wantErr, got)
				return
			}
			if got != tt.want {
				t.Errorf("Provider.SimulateRevertedTx() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
