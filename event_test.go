package celoutils

import (
	"reflect"
	"strings"
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
				event: "Transfer(address indexed _from, address indexed _to, uint256 _value)",
			},
			want: EventSignature{
				Signature: "Transfer(address,address,uint256)",
				Hash:      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EventSignatureHash(tt.args.event)
			if err != nil {
				t.Errorf("EventSignatureHash() got error: %v", err)
			}
			if got != tt.want {
				t.Errorf("EventSignatureHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventSignatureHashesFromABI(t *testing.T) {
	//    contract T {
	//      event received(address sender, uint amount, bytes memo);
	//      event receivedAddr(address sender);
	//      function receive(bytes memo) external payable {
	//        received(msg.sender, msg.value, memo);
	//        receivedAddr(msg.sender);
	//      }
	//    }
	// https://www.4byte.directory/event-signatures
	// 253837 	received(address,uint256,bytes) 	0x75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed
	// 253838 	receivedAddr(address) 	0x46923992397eac56cf13058aced2a1871933622717e27b24eabc13bf9dd329c8
	want := []EventSignature{
		{
			Signature: "received(address,uint256,bytes)",
			Hash:      "0x75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed",
		},
		{
			Signature: "receivedAddr(address)",
			Hash:      "0x46923992397eac56cf13058aced2a1871933622717e27b24eabc13bf9dd329c8",
		},
	}

	json := `[{"constant":false,"inputs":[{"name":"memo","type":"bytes"}],"name":"receive","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sender","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"memo","type":"bytes"}],"name":"received","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sender","type":"address"}],"name":"receivedAddr","type":"event"}]`
	got, err := EventSignatureHashesFromABI(strings.NewReader(json))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Slices should be equal, got %v and %v", got, want)
	}
}
