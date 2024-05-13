package main

import (
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
)

func main() {
	ethereum_client := safeethgo.NewEthereumClient("http://localhost:8545")

	tx, isPending, err := ethereum_client.GetTransaction("0xec665062d20ecfd8f4d4f2764ba967dc9220893bc64ddc52622c0c0d8d96cc1e")
	if err != nil {
		panic(err)
	}
	fmt.Println("chain id ", tx.ChainId())
	fmt.Println("cost (value + fees)", tx.Cost())
	fmt.Println("data ", tx.Data())
	fmt.Println("gas ", tx.Gas())
	fmt.Println("gas price ", tx.GasPrice())
	fmt.Println("hash ", tx.Hash())
	fmt.Println("value ", tx.Value())
	fmt.Println(isPending)

	txs, err := ethereum_client.GetTransactions(
		[]string{
			"0x03b542a2028ef4271ea6fe37ec2f86421a0df3bbddf9b4f2834c864391b9889b",
			"0xa686543e1bbe8d7cfd5a4a8f299a9a1dd371ec7df322d3c25370baf43f035cd8",
			"0x9b452aee9d436f3376beb3f8b92e467c63446770e38a574722ea1f28d0c1debb",
			"0xec665062d20ecfd8f4d4f2764ba967dc9220893bc64ddc52622c0c0d8d96cc1e",
		})
	if err != nil {
		panic(err)
	}
	for _, tx_ := range txs {
		fmt.Println("chain id ", tx_.Tx.ChainId())
		fmt.Println("cost (value + fees)", tx_.Tx.Cost())
		fmt.Println("data ", tx_.Tx.Data())
		fmt.Println("gas ", tx_.Tx.Gas())
		fmt.Println("gas price ", tx_.Tx.GasPrice())
		fmt.Println("hash ", tx_.Tx.Hash())
		fmt.Println("value ", tx_.Tx.Value())
		fmt.Println(tx_.IsPending)
	}

}
