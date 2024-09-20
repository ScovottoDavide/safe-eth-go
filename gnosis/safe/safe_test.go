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

var uri = eth.NewURI("http://127.0.0.1:8546")
var ethClient, _ = eth.EthereumClientInit(uri)
var chainId, _ = ethClient.GetChainId()

func deploySafe(sender *common.Address, ethClient *eth.EthereumClient, privateKey *ecdsa.PrivateKey) common.Address {
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	txSent, err := Create(
		ethClient,
		*sender,
		privateKey,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
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
	if txSent.ContractAddress == eth.NULL_ADDRESS {
		return eth.NULL_ADDRESS
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
	return txSent.ContractAddress
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
	safe := New(common.HexToAddress("0x43E77Ba9F5E59CefB97D55CF58641EBb7bEB22c4"), ethClient)
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

func TestGetThreshold(t *testing.T) {
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

	safe_ := New(newSafe, ethClient)
	threshold, err := safe_.RetrieveThreshold()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if threshold == nil {
		t.Fatalf("Could not retrieve threshold")
	}
	t.Log(threshold.Uint64())
	if threshold.Uint64() != 2 {
		t.Fatalf("Given threshold differs from what expected")
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
	verifyingContract := network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)]
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
	_, err = ethClient.GetReceipt(ethTxSent.TxHash.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}
	isContract, err := ethClient.IsContract(ethTxSent.ContractAddress.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !isContract {
		t.Fatalf("could not deploy master copy contract v1.3.0")
	}
}

func TestDeployCompatibilityFallbackHandler(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ethTxSent, err := DeployCompatibilityFallbackHandler(ethClient, *sender, privateKey)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ch := ethClient.WaitTxConfirmed(ethTxSent.TxHash)
	isPending := <-ch
	if isPending {
		t.Fatalf("unexpected pending tx %s", ethTxSent.TxHash.Hex())
	}
	_, err = ethClient.GetReceipt(ethTxSent.TxHash.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}
	isContract, err := ethClient.IsContract(ethTxSent.ContractAddress.Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !isContract {
		t.Fatalf("could not deploy master copy contract v1.3.0")
	}
}

func TestEstimateSafeCreation(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	gas, _gasPrice, payment, err := EstimateSafeCreation(
		ethClient, owners, 2, eth.NULL_ADDRESS, eth.NULL_ADDRESS, 1.0, 0,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(gas)
	t.Log(_gasPrice)
	t.Log(payment)
	/* SafeCreationEstimate from safe-eth-py (gas=394016, gas_price=1875000000, payment=781905000000000,
	payment_token='0x0000000000000000000000000000000000000000')*/
}
