package safe

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const owner2 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const owner3 = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

var uri = eth.NewURI("http://127.0.0.1:8545")
var ethClient, _ = eth.EthereumClientInit(uri)

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

func TestSafeVersion(t *testing.T) {
	safe := New(eth.NULL_ADDRESS, ethClient)
	if safe == nil {
		t.Fatalf("Got nil Safe object")
	}

	version, err := safe.Version()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if version == "" {
		t.Fatalf("Empty version string: %s", version)
	}
}

func TestSafeDomainSeparator(t *testing.T) {
	safe := New(eth.NULL_ADDRESS, ethClient)
	if safe == nil {
		t.Fatalf("Got nil Safe object")
	}

	domainSeparator, err := safe.DomainSeparator()
	if err != nil {
		t.Fatalf(err.Error())
	}
	chainId, _ := ethClient.GetChainId()
	verifyingContract := network.NetworkToSafeAddress[network.GetNetwork(chainId)]
	dataDomain := apitypes.TypedDataDomain{
		ChainId:           math.NewHexOrDecimal256(int64(chainId)),
		VerifyingContract: verifyingContract.Address.Hex(),
	}
	types := apitypes.Types{
		"EIP712Domain": {
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}
	typedData := apitypes.TypedData{
		Types:  types,
		Domain: dataDomain,
	}
	expectedDomainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Fatalf(err.Error())
	}
	if domainSeparator != common.Hash(expectedDomainSeparator) {
		t.Fatalf("domain separator mismatch")
	}
}

func TestDeployMasterContract_v1_3_0(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ethTxSent, err := DeployMasterContract_v1_3_0(ethClient, *sender, privateKey)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ch := ethClient.WaitTxConfirmed(ethTxSent.TxHash)
	isPending := <-ch
	if isPending {
		t.Fatalf("unexpected pending tx %s", ethTxSent.TxHash.Hex())
	}
	isContract, err := ethClient.IsContract(ethTxSent.contractAaddress.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !isContract {
		t.Fatalf("could not deploy master copy contract v1.3.0")
	}
}
