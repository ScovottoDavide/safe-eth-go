package main

import (
	"context"
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
)

func main() {
	ethereum_client := safeethgo.NewEthereumClient("http://localhost:8545")

	uri := ethereum_client.GetUri()
	fmt.Println(uri.GetAddress())

	geth_client := ethereum_client.GetGEthClient()
	chain_id, err := geth_client.ChainID(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Println(chain_id)
}
