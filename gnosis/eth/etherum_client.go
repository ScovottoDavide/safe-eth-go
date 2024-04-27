package eth

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	ethereumNodeUrl URI
	ethereumClient  *ethclient.Client
}

func EthereumClientInit(uri *URI) (*EthereumClient, error) {
	client, err := ethclient.Dial(uri.address)
	if err != nil {
		return nil, err
	}

	chain_id, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	if !isSupported(int(chain_id.Int64())) {
		return nil, errors.New("chain not supported")
	}
	fmt.Println("Successfully connected to network with {address: ", uri.address, ", chainId: ", chain_id, "}")

	ethereum_client := &EthereumClient{
		ethereumNodeUrl: *uri,
		ethereumClient:  client,
	}
	return ethereum_client, nil
}

func (ethereumClient *EthereumClient) GetGEthClient() *ethclient.Client {
	return ethereumClient.ethereumClient
}

func (ethereumClient *EthereumClient) GetUri() URI {
	return ethereumClient.ethereumNodeUrl
}
