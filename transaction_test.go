package celoutils

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/crypto"
	"github.com/grassrootseconomics/w3-celo"
)

// Throwaway test keys.
var (
	testNetRpc  = "https://alfajores-forno.celo-testnet.org"
	privateKey  = "c74ecb579f9822e196b1866fef65950f5f9b8ed128ca92260b0de3c4dca8d436"
	publicKey   = "0x3AA8028a5FD03a0D35C32347e746842689b30987"
	deadAddress = w3.A("0x000000000000000000000000000000000000dEaD")
)

func TestProvider_SignContractExecutionTx(t *testing.T) {
	p, err := NewProvider(ProviderOpts{
		ChainId:     TestnetChainId,
		RpcEndpoint: testNetRpc,
	})
	if err != nil {
		t.Fatal("RPC endpoint parsing failed")
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key %v", err)
	}

	sampleFunc := w3.MustNewFunc("transfer(address to, uint256 amount)", "bool")
	input, err := sampleFunc.EncodeArgs(deadAddress, w3.I("1"))
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		privateKey *ecdsa.PrivateKey
		txData     ContractExecutionTxOpts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sign ERC20 transfer",
			args: args{
				privateKey: privateKey,
				txData: ContractExecutionTxOpts{
					// cUSD
					ContractAddress: w3.A("0x874069Fa1Eb16D44d622F2e0Ca25eeA172369bC1"),
					InputData:       input,
					GasLimit:        250000,
					GasFeeCap:       SafeGasFeeCap,
					GasTipCap:       SafeGasTipCap,
					Nonce:           0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.SignContractExecutionTx(tt.args.privateKey, tt.args.txData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SignContractExecutionTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProvider_SignGasTransferTx(t *testing.T) {
	p, err := NewProvider(ProviderOpts{
		ChainId:     TestnetChainId,
		RpcEndpoint: testNetRpc,
	})
	if err != nil {
		t.Fatal("RPC endpoint parsing failed")
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key %v", err)
	}

	type args struct {
		privateKey *ecdsa.PrivateKey
		txData     GasTransferTxOpts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sign gas transfer",
			args: args{
				privateKey: privateKey,
				txData: GasTransferTxOpts{
					To:        deadAddress,
					Value:     big.NewInt(1),
					GasFeeCap: SafeGasFeeCap,
					GasTipCap: SafeGasTipCap,
					Nonce:     0,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.SignGasTransferTx(tt.args.privateKey, tt.args.txData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SignGasTransferTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProvider_SignGasTransferTxPayWithCUSD(t *testing.T) {
	cUSD := w3.A(CUSDContractTestnet)

	p, err := NewProvider(ProviderOpts{
		ChainId:     TestnetChainId,
		RpcEndpoint: testNetRpc,
	})
	if err != nil {
		t.Fatal("RPC endpoint parsing failed")
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key %v", err)
	}

	type args struct {
		privateKey *ecdsa.PrivateKey
		txData     types.TxData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sign gas transfer, pay with cUSD",
			args: args{
				privateKey: privateKey,
				txData: &types.CeloDynamicFeeTx{
					To:          &deadAddress,
					Gas:         21000 + 50000,
					FeeCurrency: &cUSD,
					GasFeeCap:   SafeGasFeeCap,
					GasTipCap:   SafeGasTipCap,
					Nonce:       0,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		_, err := types.SignNewTx(tt.args.privateKey, p.Signer, tt.args.txData)
		if (err != nil) != tt.wantErr {
			t.Errorf("types.SignNewTx error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
