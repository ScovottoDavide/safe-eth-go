package network

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TxSpeed struct {
	Speed int
}

type Network struct {
	ChainId int
}

type GnosisSafeV130 struct {
	Address common.Address
	Path    string
}

func (a GnosisSafeV130) GetAbi() (*abi.ABI, error) {
	contractFile, err := os.Open(a.Path)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(contractFile)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}
