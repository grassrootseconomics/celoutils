package celoutils

import (
	"strings"

	"github.com/grassrootseconomics/w3-celo"
)

type (
	EventSignature struct {
		Signature string
		Hash      string
	}
)

func EventSignatureFromString(eventSignature string) (EventSignature, error) {
	event, err := w3.NewEvent(strings.TrimSuffix(eventSignature, ";"))
	if err != nil {
		return EventSignature{}, err
	}

	return EventSignature{
		Signature: event.Signature,
		Hash:      event.Topic0.Hex(),
	}, nil
}
