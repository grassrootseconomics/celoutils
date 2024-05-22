package celoutils

import (
	"testing"
)

func TestNewProvider_NoPanic(t *testing.T) {
	NewProvider("https://forno.celo.org", CeloMainnet)
}

func TestNewProvider_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on bad url")
		}
	}()

	NewProvider("not://a.good.url", CeloMainnet)

}
