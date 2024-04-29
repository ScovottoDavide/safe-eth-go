package main

import (
	"context"
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

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

	clientVersion, err := ethereum_client.GetClientVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(clientVersion)

	// fmt.Println(ethereum_client.IsEip1559Supported())
	address, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
