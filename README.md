# Safe ETH Go

Safe ETH Go is a Go-based rewrite of the official [Safe ETH Python (v5.8.0)](https://github.com/safe-global/safe-eth-py/tree/v5.8.0) library. This project is primarily is just an attempt to get more comfortable with `Go` and with `Account Abstraction`.

This library aims to provide a seamless transition for users of the Safe ETH Python library while helping developers build their Safe-backed application in Go.

## Disclaimer

This library is a learning project and should not be used in production environments. It may contain bugs or security vulnerabilities that have not been addressed. Use at your own risk.

## Features

- **Simplified Ethereum Client**: Provides a more direct way of interacting with the Ethereum Node.
- **Smart Contract Interaction**: Easily interact with Safe deployed smart contracts on the Ethereum network.
- **Safe Module**: Provides helper methods to easily create/manage/use Safe Accounts.
    - **Safe Transaction**: Easily create and manage Safe Transactions
    - **Sign Transaction**: Sign Safe Transactions. Allows to correctly sign the transaction by the Safe Account owners
    - **Safe Creations**: Predict the GnosisSafeProxy address. Supports both CREATE_OP and CREATE2_OP "predictions".
    - **WIP...**

## Installation

To install Safe ETH Go, use `go get`:

```bash
go get github.com/ScovottoDavide/safe-eth-go
```

## Usage

Here’s a basic example of how to use Safe ETH Go to create a wallet and send a transaction:

```go
package main

import (
	"fmt"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/safe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	uri := eth.NewURI(provider_uri)
	ethClient, err := eth.EthereumClientInit(uri)
	if err != nil {
		panic(err)
	}
	chainId, err := ethClient.GetChainId()
	if err != nil {
		panic(err)
	}
	
	var privateKey, _ = eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	creator, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		panic(err)
	}

	owner_1, _ := eth.RandomAddress()
	owner_2, _ := eth.RandomAddress()
	owner_3, _ := eth.RandomAddress()
	owners := []common.Address{*owner_1, *owner_2, *owner_3}
	threshold := int64(2)

	txSent, err := safe.Create(
		ethClient,
		*creator,
		privateKey,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address, // Safe Master Copy Address
		owners, // owners of the Safe Account
		threshold, // signatures threshold
		eth.NULL_ADDRESS,
		network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address, // Safe Proxy Factory Address
		eth.NULL_ADDRESS,
		0,
		eth.NULL_ADDRESS,
	)
	if err != nil {
		panic(err)
	}
	if txSent.ContractAddress == eth.NULL_ADDRESS {
		panic("Could not create Safe Account")
	}
	ch := ethClient.WaitTxConfirmed(txSent.TxHash)
	isPending := <-ch
	if isPending {
		fmt.Printf("unexpected pending tx %s", txSent.TxHash.Hex())
	}
	receipt, _ := ethClient.GetReceipt(txSent.TxHash.Hex())
	fmt.Println("Used gas: ", receipt.GasUsed)
	fmt.Println("Safe creation payment: ", receipt.GasUsed*receipt.EffectiveGasPrice.Uint64())
	fmt.Println(txSent.ContractAddress.Hex())
}
```

## Contact

For questions or support, please open an issue on the repository or contact the maintainer at [scovottodavide@gmail.com].

---
