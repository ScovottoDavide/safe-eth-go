package main

import (
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
)

func main() {
	ethereum_client := safeethgo.NewEthereumClient("http://localhost:8545")

	tx, isPending, err := ethereum_client.GetTransaction("0x5838d8d782178bbebef1b51922244674c883eabe85d3e634f446a8bc5742dd71")
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
			"0x5838d8d782178bbebef1b51922244674c883eabe85d3e634f446a8bc5742dd71",
			"0xf545d5220b5e3493238b3e4811810fb733996b02f20eb0683dbcb88be01f9558",
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
