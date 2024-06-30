package main

import (
	"context"
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const S_KEY_0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

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

	fmt.Println(ethereum_client.IsEip1559Supported())
	fmt.Println(ethereum_client.IsEip1559Supported())
	address, err := eth.AddressFromPrivKey(hexutil.MustDecode(S_KEY_0))
	if err != nil {
		panic(err)
	}
	fmt.Println(address)

	nonce, err := ethereum_client.GetNonceForAccount(*address, "latest")
	if err != nil {
		panic(err)
	}
	fmt.Println("Nonce: ", nonce)

	to_address, err := eth.StringToAddress("0x63f57569bC23c52165A98144D896a2AFC6fF8DFA")
	if err != nil {
		panic(err)
	}
	gas_price, err := ethereum_client.GasPrice()
	if err != nil {
		panic(err)
	}

	amount_to_transfer := eth.ToWei(0.5, 18)
	estimated_gas, err := ethereum_client.EstimateGas(*address, to_address, 0, gas_price, nil, nil, amount_to_transfer, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Estimated gas for transferring ", amount_to_transfer, "(", eth.ToDecimal(amount_to_transfer, 18), ")", ": ", estimated_gas)

	txHash, err := ethereum_client.SendEthTo(
		S_KEY_0,
		to_address,
		gas_price,
		amount_to_transfer,
		estimated_gas,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(txHash)
}
