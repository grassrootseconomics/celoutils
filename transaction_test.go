package celo

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	PrivateKey = "c74ecb579f9822e196b1866fef65950f5f9b8ed128ca92260b0de3c4dca8d436"
	PublicKey  = "0x3AA8028a5FD03a0D35C32347e746842689b30987"
)

func TestProvider_SignContractExecutionTx(t *testing.T) {
	var nonce uint64

	p, err := NewProvider(ProviderOpts{
		ChainId:     TestnetChainId,
		RpcEndpoint: TestnetRpcEndpoint,
	})
	if err != nil {
		t.Fatal("RPC endpoint parsing failed")
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key %v", err)
	}

	err = p.Client.CallCtx(
		context.Background(),
		eth.Nonce(w3.A(PublicKey), nil).Returns(&nonce),
	)
	if err != nil {
		t.Fatal("Failed to fetch test account nonce")
	}

	gasPrice, err := p.GetOptimumGasPrice(context.Background())
	if err != nil {
		t.Fatal("Failed to fetch gas price")
	}

	sampleFunc := w3.MustNewFunc("transfer(address to, uint256 amount)", "bool")
	input, err := sampleFunc.EncodeArgs(w3.A("0x000000000000000000000000000000000000dEaD"), w3.I("1"))
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
					GasPrice:        gasPrice,
					Nonce:           nonce,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawSignedTx, err := p.SignContractExecutionTx(tt.args.privateKey, tt.args.txData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SignContractExecutionTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var txH common.Hash
			err = p.Client.CallCtx(
				context.Background(),
				eth.SendTx(rawSignedTx).Returns(&txH),
			)
			if err != nil {
				t.Fatal("Failed to submit signed tx")
			}

			t.Logf("Testnet transfer hash: %s", txH.Hex())
			t.Log("Waiting 7 seconds for the next block to be mined...")
			// Wait for the next block to be mined so as to safely use nonce+1 in the next test
			time.Sleep(time.Second * 7)
		})
	}
}

func TestProvider_SignGasTransferTx(t *testing.T) {
	var nonce uint64

	p, err := NewProvider(ProviderOpts{
		ChainId:     TestnetChainId,
		RpcEndpoint: TestnetRpcEndpoint,
	})
	if err != nil {
		t.Fatal("RPC endpoint parsing failed")
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key %v", err)
	}

	err = p.Client.CallCtx(
		context.Background(),
		eth.Nonce(w3.A(PublicKey), nil).Returns(&nonce),
	)
	if err != nil {
		t.Fatal("Failed to fetch test account nonce")
	}

	gasPrice, err := p.GetOptimumGasPrice(context.Background())
	if err != nil {
		t.Fatal("Failed to fetch gas price")
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
					// CELO
					To:       w3.A("0x000000000000000000000000000000000000dEaD"),
					Value:    big.NewInt(1),
					GasPrice: gasPrice,
					Nonce:    nonce,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rawSignedTx, err := p.SignGasTransferTx(tt.args.privateKey, tt.args.txData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.SignGasTransferTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var txH common.Hash
			err = p.Client.CallCtx(
				context.Background(),
				eth.SendTx(rawSignedTx).Returns(&txH),
			)
			if err != nil {
				t.Fatal("Failed to submit signed tx")
			}

			t.Logf("Testnet gas transfer hash: %s", txH.Hex())
		})
	}
}
