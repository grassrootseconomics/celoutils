package celoutils

import (
	"reflect"
	"testing"
)

func TestMethodSignatureFromString(t *testing.T) {
	type args struct {
		methodSignature string
	}
	tests := []struct {
		name    string
		args    args
		want    MethodSignature
		wantErr bool
	}{
		{
			name: "Transfer Method",
			args: args{
				methodSignature: "transfer(address, uint256)",
			},
			want: MethodSignature{
				Signature: "transfer(address,uint256)",
				Hash:      "a9059cbb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodSignatureFromString(tt.args.methodSignature)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodSignatureFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MethodSignatureFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
