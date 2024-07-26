package eth_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"

	safeethgo "github.com/ScovottoDavide/safe-eth-go"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

var ethereum_client = safeethgo.NewEthereumClient("http://localhost:8545")

func TestEthClientInit(t *testing.T) {
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
}

func TestAddressFromPrivKey(t *testing.T) {
	address, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	if address.Hex() != "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" {
		t.Fatalf("address mismatch")
	}

	randKey := make([]byte, 30)
	n, _ := rand.Read(randKey)
	if n != 30 {
		t.Fatalf("cannot generate random key")
	}
	address, err = eth.AddressFromPrivKey(randKey)
	if err == nil && address != nil {
		t.Fatalf("address should be nil and an error should occur instead")
	}
}

func TestSendEthTo(t *testing.T) {
	address, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}

	to_address, _ := eth.StringToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	gas_price, _ := ethereum_client.GasPrice()
	amount_to_transfer := eth.ToWei(0.5, 18)
	estimated_gas, err := ethereum_client.EstimateGas(
		*address, to_address, 0, gas_price, nil, nil, amount_to_transfer, nil, nil, nil, nil,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println("Estimated gas for transferring ", amount_to_transfer, "(", eth.ToDecimal(amount_to_transfer, 18), ")", ": ", estimated_gas)

	txHash, err := ethereum_client.SendEthTo(
		HARDHAT_S_KEY0,
		to_address,
		gas_price,
		amount_to_transfer,
		estimated_gas,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(txHash)
	fmt.Println("Waiting for tx confirmation")
	ch := ethereum_client.WaitTxConfirmed(txHash)
	isPending := <-ch
	if isPending {
		t.Fatalf("transaction should be confirmed")
	}
	fmt.Println("Transaction confirmed")

	receipt := ethereum_client.GetTransactionReceipt(txHash)
	if receipt == nil {
		t.Fatalf("receipt not found")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Fatalf("Transaction failedddd")
	}
	fmt.Println("Transaction successful")
}
