package eth

import (
	"fmt"
	"math/big"
	"net/url"
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
	Reward  big.Int
	BaseFee big.Int
}
