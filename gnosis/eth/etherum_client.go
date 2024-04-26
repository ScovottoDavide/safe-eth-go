package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	ethereum_node_url URI
	ethereum_client   *ethclient.Client
}

func ethereum_client_init(uri URI) *EthereumClient {
	if uri.isValid() {
		client, err := ethclient.Dial(uri.address)
		if err != nil {
			panic(err)
		}
		fmt.Println("we have a connection")

		ethereum_client := new(EthereumClient)
		ethereum_client.ethereum_node_url = uri
		ethereum_client.ethereum_client = client
	}
	return nil
}
