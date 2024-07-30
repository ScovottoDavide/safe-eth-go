package safe

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const owner2 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const owner3 = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

func deploySafe(sender *common.Address, ethClient *eth.EthereumClient, privateKey *ecdsa.PrivateKey) common.Address {
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	txSent, err := Create(
		ethClient,
		*sender,
		privateKey,
		common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"),
		owners,
		2,
		eth.NULL_ADDRESS,
		common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
		eth.NULL_ADDRESS,
		0,
		eth.NULL_ADDRESS,
	)
	if err != nil {
		return eth.NULL_ADDRESS
	}
	if txSent.contractAaddress == eth.NULL_ADDRESS {
		return eth.NULL_ADDRESS
	}
	fmt.Println(txSent.contractAaddress.Hex())
	return txSent.contractAaddress
}

func TestCreateNewSafe(t *testing.T) {
	uri := eth.NewURI("http://127.0.0.1:8545")
	ethClient, err := eth.EthereumClientInit(uri)
	if err != nil {
		t.Fatalf(err.Error())
	}

	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	newSafe := deploySafe(sender, ethClient, privateKey)
	if newSafe == eth.NULL_ADDRESS {
		t.Fatalf("new safe not deployed")
	}
}

func TestSafeCall(t *testing.T) {
	uri := eth.NewURI("http://127.0.0.1:8545")
	ethClient, err := eth.EthereumClientInit(uri)
	if err != nil {
		panic(err)
	}
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	deployedNewSafe := deploySafe(sender, ethClient, privateKey)
	if deployedNewSafe == eth.NULL_ADDRESS {
		t.Fatalf("new safe could not be deployed")
	}
	safe := New(deployedNewSafe, ethClient)
	if safe == nil {
		t.Fatalf("Got nil Safe object")
	}

	version, err := safe.Version()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(version)
}
