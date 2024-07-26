package main

import (
	"fmt"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
)

func main() {
	ethereum_client := safeethgo.NewEthereumClient("http://localhost:8545")

	if txReceipt, err := ethereum_client.GetReceipt("0x5838d8d782178bbebef1b51922244674c883eabe85d3e634f446a8bc5742dd71"); err != nil {
		panic(err)
	} else {
		if txReceipt == nil {
			panic(fmt.Errorf("receipt is nil"))
		}
	}

	txs, err := ethereum_client.GetReceipts(
		[]string{
			"0x5838d8d782178bbebef1b51922244674c883eabe85d3e634f446a8bc5742dd71",
			"0xf545d5220b5e3493238b3e4811810fb733996b02f20eb0683dbcb88be01f9558",
		})
	if err != nil {
		panic(err)
	}
	for _, tx_ := range txs {
		if tx_ == nil {
			panic(fmt.Errorf("receipt is nil"))
		}
	}

}
