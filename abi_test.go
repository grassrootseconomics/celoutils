package celoutils

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseABI(t *testing.T) {
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
	eventSignatures := []EventSignature{
		{
			Signature: "received(address,uint256,bytes)",
			Hash:      "0x75fd880d39c1daf53b6547ab6cb59451fc6452d27caa90e5b6649dd8293b9eed",
		},
		{
			Signature: "receivedAddr(address)",
			Hash:      "0x46923992397eac56cf13058aced2a1871933622717e27b24eabc13bf9dd329c8",
		},
	}

	methodSignatures := []MethodSignature{
		{
			Signature: "receive(bytes)",
			Hash:      "a69b6ed0",
		},
	}

	want := &ParsedABI{
		EventSignatures:  eventSignatures,
		MethodSignatures: methodSignatures,
	}

	json := `[{"constant":false,"inputs":[{"name":"memo","type":"bytes"}],"name":"receive","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sender","type":"address"},{"indexed":false,"name":"amount","type":"uint256"},{"indexed":false,"name":"memo","type":"bytes"}],"name":"received","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"sender","type":"address"}],"name":"receivedAddr","type":"event"}]`
	got, err := ParseABI(strings.NewReader(json))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got.EventSignatures, want.EventSignatures) {
		t.Errorf("Event signatures slices should be equal, got %v and %v", got, want)
	}

	if !reflect.DeepEqual(got.MethodSignatures, want.MethodSignatures) {
		t.Errorf("Method signatures slices should be equal, got %v and %v", got, want)
	}
}
