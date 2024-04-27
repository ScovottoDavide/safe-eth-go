package safeethgo

import (
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
)

func NewEthereumClient(provider_uri string) *eth.EthereumClient {
	uri := eth.NewURI(provider_uri)
	client, err := eth.EthereumClientInit(uri)

	if err != nil {
		panic(err)
	}
	return client
}
