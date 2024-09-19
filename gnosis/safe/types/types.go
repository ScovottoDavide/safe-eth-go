package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumTxSent struct {
	TxHash          common.Hash
	Tx              *types.Transaction
	ContractAddress common.Address
}
