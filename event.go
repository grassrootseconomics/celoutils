package celoutils

import (
	"io"
	"strings"

	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/grassrootseconomics/w3-celo"
)

func EventSignatureHash(eventSignature string) (string, error) {
	event, err := w3.NewEvent(strings.TrimSuffix(eventSignature, ";"))
	if err != nil {
		return "", err
	}

	return event.Topic0.Hex(), nil
}

func EventSignatureHashesFromABI(abiReader io.Reader) ([]string, error) {
	abi, err := abi.JSON(abiReader)
	if err != nil {
		return nil, err
	}

	var eventSignatureHashes []string

	for _, v := range abi.Events {
		eventSignatureHash, err := EventSignatureHash(v.Sig)
		if err != nil {
			return nil, err
		}

		eventSignatureHashes = append(eventSignatureHashes, eventSignatureHash)
	}

	return eventSignatureHashes, nil
}
