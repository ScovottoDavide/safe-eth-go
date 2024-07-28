package eth

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const HARDHAT_S_KEY1 = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

var ethereum_client, _ = EthereumClientInit(NewURI("http://localhost:8545"))

func TestEthClientInit(t *testing.T) {
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
	address, err := AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
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
	address, err = AddressFromPrivKey(randKey)
	if err == nil && address != nil {
		t.Fatalf("address should be nil and an error should occur instead")
	}
}

func TestSendEthTo(t *testing.T) {
	address, err := AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}

	to_address, _ := StringToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	gas_price, _ := ethereum_client.GasPrice()
	amount_to_transfer := ToWei(0.5, 18)
	estimated_gas, err := ethereum_client.EstimateGas(
		*address, to_address, 0, gas_price, nil, nil, amount_to_transfer, nil, nil, nil, nil,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println("Estimated gas for transferring ", amount_to_transfer, "(", ToDecimal(amount_to_transfer, 18), ")", ": ", estimated_gas)

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

	receipt, err := ethereum_client.GetReceipt(txHash.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}
	if receipt == nil {
		t.Fatalf("receipt not found")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Fatalf("Transaction failedddd")
	}
	fmt.Println("Transaction successful")
}

func TestDeployContract(t *testing.T) {
	hexBytecode, err := os.ReadFile("../contracts/MyToken.bin")
	if err != nil {
		t.Fatal(err.Error())
	}
	// append the constructor params at the end of the bytecode
	ownerAddress, _ := AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	// pad the 20byte addr into a 32byte word
	binOwnerAddress := append(make([]byte, 12), ownerAddress.Bytes()...)
	data := multipleTxData{
		constructorData: append(common.Hex2Bytes(string(hexBytecode)), binOwnerAddress...),
	}
	_, _, contractAddress, err := ethereum_client.DeployAndInitializeContract(
		HARDHAT_S_KEY0,
		data,
		true,
		Normal,
		*big.NewInt(0),
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	isContract, err := ethereum_client.IsContract(contractAddress)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !isContract {
		t.Fatalf("contract address is not an actual contract")
	}
}
