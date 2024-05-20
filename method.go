package celoutils

import (
	"encoding/hex"

	"github.com/lmittmann/w3"
)

type (
	MethodSignature struct {
		Signature string
		Hash      string
	}
)

func MethodSignatureFromString(methodSignature string) (MethodSignature, error) {
	method, err := w3.NewFunc(methodSignature, "")
	if err != nil {
		return MethodSignature{}, err
	}

	return MethodSignature{
		Signature: method.Signature,
		Hash:      hex.EncodeToString(method.Selector[:]),
	}, nil
}
