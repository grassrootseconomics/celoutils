package celoutils

import (
	"strings"

	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/common"
)

func PrepareContractBytecodeData(contractBin string, contractABIJSON string, constructorArgs []any) ([]byte, error) {
	constructorBytecode, err := encodeConstructorArgs(contractABIJSON, constructorArgs)
	if err != nil {
		return nil, err
	}

	return append(common.Hex2Bytes(contractBin), constructorBytecode...), nil
}

func encodeConstructorArgs(contractABI string, constructorArgs []any) ([]byte, error) {
	abi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}

	constructorBytecode, err := abi.Pack("", constructorArgs...)
	if err != nil {
		return nil, err
	}

	return constructorBytecode, nil
}
