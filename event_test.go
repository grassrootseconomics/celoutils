package celoutils

import (
	"testing"
)

func TestEventSignatureHash(t *testing.T) {
	type args struct {
		event string
	}
	tests := []struct {
		name string
		args args
		want EventSignature
	}{
		{
			name: "Transfer Event variant 1",
			args: args{
				event: "Transfer(address,address,uint256)",
			},
			want: EventSignature{
				Signature: "Transfer(address,address,uint256)",
				Hash:      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
		},
		{
			name: "Transfer Event variant 2",
			args: args{
				event: "Transfer(address indexed _from, address indexed _to, uint256 _value)",
			},
			want: EventSignature{
				Signature: "Transfer(address,address,uint256)",
				Hash:      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
		},
		{
			name: "Transfer Event variant 3",
			args: args{
				event: "Transfer(address indexed _from, address indexed _to, uint256 _value);",
			},
			want: EventSignature{
				Signature: "Transfer(address,address,uint256)",
				Hash:      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EventSignatureFromString(tt.args.event)
			if err != nil {
				t.Errorf("EventSignatureHash() got error: %v", err)
			}
			if got != tt.want {
				t.Errorf("EventSignatureHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
