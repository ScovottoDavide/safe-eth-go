package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumTxSent struct {
	TxHash           common.Hash
	tx               *types.Transaction
	contractAaddress common.Address
}
