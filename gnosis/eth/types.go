package eth

import (
	"fmt"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type URI struct {
	address string
}

func NewURI(uri string) *URI {
	new_uri := &URI{address: uri}
	if new_uri.isValid() {
		return new_uri
	}
	return nil
}

func (uri *URI) isValid() bool {
	_, err := url.ParseRequestURI(uri.address)
	if err != nil {
		fmt.Printf("Provided Invalid URI for the EthereumClient: %s", err)
		return false
	}
	return true
}

func (uri *URI) GetAddress() string {
	return uri.address
}

type Network struct {
	chainId int
}

type TxSpeed struct {
	speed int
}

type EIP1559EstimatedGas struct {
	Reward  *big.Int
	BaseFee *big.Int
}

// from ethclient
type rpcTransaction struct {
	Tx *types.Transaction
	txExtraInfo
}
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

type batcthedTransactionResult struct {
	Tx        *types.Transaction
	IsPending bool
}

type multipleTxData struct {
	constructorData []byte
	initializerData []byte
}

func (m multipleTxData) toArray() [][]byte {
	var arr [][]byte
	if m.initializerData == nil {
		arr = make([][]byte, 1)
		arr[0] = m.constructorData
	} else {
		arr = make([][]byte, 2)
		arr[0] = m.constructorData
		arr[1] = m.initializerData
	}
	return arr
}
