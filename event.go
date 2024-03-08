package celoutils

import (
	"io"
	"strings"

	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/grassrootseconomics/w3-celo"
)

type (
	EventSignature struct {
		Signature string
		Hash      string
	}
)

func EventSignatureHash(eventSignature string) (EventSignature, error) {
	event, err := w3.NewEvent(strings.TrimSuffix(eventSignature, ";"))
	if err != nil {
		return EventSignature{}, err
	}

	return EventSignature{
		Signature: event.Signature,
		Hash:      event.Topic0.Hex(),
	}, nil
}

func EventSignatureHashesFromABI(abiReader io.Reader) ([]EventSignature, error) {
	abi, err := abi.JSON(abiReader)
	if err != nil {
		return nil, err
	}

	var eventSignatures []EventSignature

	for _, v := range abi.Events {
		eventSignatureHash, err := EventSignatureHash(v.Sig)
		if err != nil {
			return nil, err
		}

		eventSignatures = append(eventSignatures, eventSignatureHash)
	}

	return eventSignatures, nil
}
