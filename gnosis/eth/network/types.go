package network

import (
	"github.com/ethereum/go-ethereum/common"
)

type TxSpeed struct {
	Speed int
}

type Network struct {
	ChainId int
}

type GnosisSafeContract struct {
	Address common.Address
}
