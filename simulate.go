package celoutils

import (
	"context"
	"math/big"
	"strings"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/grassrootseconomics/w3-celo/module/eth"
	"github.com/grassrootseconomics/w3-celo/w3types"
)

// SimulateRevertedTx attempts to simulate a reverted tx and dump its revert reason
func (p *Provider) SimulateRevertedTx(ctx context.Context, txHash common.Hash, blockNumber *big.Int) (string, error) {
	var (
		tx     types.Transaction
		output []byte
	)

	if err := p.Client.CallCtx(
		ctx,
		eth.Tx(txHash).Returns(&tx),
	); err != nil {
		return "", err
	}

	txMsg, err := tx.AsMessage(p.Signer, nil)
	if err != nil {
		return "", err
	}

	if err := p.Client.CallCtx(
		ctx,
		eth.Call(&w3types.Message{
			From:     txMsg.From(),
			To:       tx.To(),
			Input:    tx.Data(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice(),
			Value:    tx.Value(),
		}, blockNumber, nil).Returns(&output),
	); err != nil {
		revert, reason := parseError(err)
		if revert {
			return reason, nil
		}
		return "", err
	}

	return "", nil
}

func parseError(err error) (bool, string) {
	executionRevertMsg := "w3: call failed: execution reverted: "

	if strings.Contains(err.Error(), executionRevertMsg) {
		return true, strings.TrimLeft(err.Error(), executionRevertMsg)
	}

	return false, ""
}
