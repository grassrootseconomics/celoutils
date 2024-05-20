package celoutils

import (
	"io"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type (
	ParsedABI struct {
		EventSignatures  []EventSignature
		MethodSignatures []MethodSignature
	}
)

func ParseABI(abiReader io.Reader) (*ParsedABI, error) {
	abi, err := abi.JSON(abiReader)
	if err != nil {
		return nil, err
	}

	var (
		eventSignatures   []EventSignature
		methodsSignatures []MethodSignature
	)

	for _, v := range abi.Events {
		eventSignature, err := EventSignatureFromString(v.Sig)
		if err != nil {
			return nil, err
		}

		eventSignatures = append(eventSignatures, eventSignature)
	}

	for _, v := range abi.Methods {
		methodSignature, err := MethodSignatureFromString(v.Sig)
		if err != nil {
			return nil, err
		}

		methodsSignatures = append(methodsSignatures, methodSignature)
	}

	return &ParsedABI{
		EventSignatures:  eventSignatures,
		MethodSignatures: methodsSignatures,
	}, nil
}
