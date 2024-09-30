package safe

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
var testSafe = eth.NULL_ADDRESS
var privateKey, _ = eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
var signer, _ = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))

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

func createTestSafe() error {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		return err
	}

	newSafe := deploySafe(sender, ethClient, privateKey)
	if newSafe == eth.NULL_ADDRESS {
		return err
	}
	testSafe = newSafe
	return nil
}

func TestCreateNewSafe(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}

	newSafe := deploySafe(sender, ethClient, privateKey)
	if newSafe == eth.NULL_ADDRESS {
		t.Fatalf("new safe not deployed")
	}
}

func TestSafeVersion(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe := New(testSafe, ethClient, privateKey)
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
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
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
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe := New(testSafe, ethClient, privateKey)

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

func TestEstimateTxWithSafe(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)

	// we want to estimate the Safe sending 1 eth to this randAddr
	// so we need to fund the safe
	randAddr, _ := eth.RandomAddress()
	balanceBefore, _ := ethClient.GetBalance(safe_)

	gasPrice, _ := ethClient.GasPrice()
	txHash, err := ethClient.SendEthTo(hexutil.EncodeBig(privateKey.D), safe_.SafeAddress, gasPrice, big.NewInt(1e18), uint64(3000000))
	if err != nil {
		t.Fatalf(err.Error())
	}
	ch := ethClient.WaitTxConfirmed(txHash)
	isPending := <-ch
	if isPending {
		t.Fatalf("unexpected pending tx %s", txHash.Hex())
	}

	balanceAfter, _ := ethClient.GetBalance(safe_.SafeAddress)

	if balanceAfter.Uint64() <= balanceBefore.Uint64() {
		t.Fatalf("did not send 1 eth to randAddr")
	}

	estimatedGas, err := safe_.EstimateTxGasWithSafe(
		*randAddr,
		1e17,
		nil,
		0, 0, signer.Signer,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if estimatedGas <= 0 {
		t.Fatalf("could not estimate gas with safe")
	}
	t.Log(estimatedGas)
}

func TestEstimateTxGasWithWeb3(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)

	randAddr, _ := eth.RandomAddress()

	balanceBefore, _ := ethClient.GetBalance(safe_)
	if balanceBefore.Uint64() < 1e17 {
		gasPrice, _ := ethClient.GasPrice()
		txHash, err := ethClient.SendEthTo(hexutil.EncodeBig(privateKey.D), safe_.SafeAddress, gasPrice, big.NewInt(1e18), uint64(3000000))
		if err != nil {
			t.Fatalf(err.Error())
		}
		ch := ethClient.WaitTxConfirmed(txHash)
		isPending := <-ch
		if isPending {
			t.Fatalf("unexpected pending tx %s", txHash.Hex())
		}

		balanceAfter, _ := ethClient.GetBalance(safe_.SafeAddress)

		if balanceAfter.Uint64() <= balanceBefore.Uint64() {
			t.Fatalf("did not send 1 eth to randAddr")
		}
	}
	web3EstimatedGas, err := safe_.EstimateTxGasWithWeb3(
		*randAddr,
		1e15,
		nil,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if web3EstimatedGas == 0 {
		t.Fatalf("could not estimate gas with eth node estimateGas method")
	}

	t.Log(web3EstimatedGas)
}

func TestEstimateTxGasByTrying(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
	randAddr, _ := eth.RandomAddress()

	balanceBefore, _ := ethClient.GetBalance(safe_)
	if balanceBefore.Uint64() < 1e17 {
		gasPrice, _ := ethClient.GasPrice()
		txHash, err := ethClient.SendEthTo(hexutil.EncodeBig(privateKey.D), safe_.SafeAddress, gasPrice, big.NewInt(1e18), uint64(3000000))
		if err != nil {
			t.Fatalf(err.Error())
		}
		ch := ethClient.WaitTxConfirmed(txHash)
		isPending := <-ch
		if isPending {
			t.Fatalf("unexpected pending tx %s", txHash.Hex())
		}

		balanceAfter, _ := ethClient.GetBalance(safe_.SafeAddress)

		if balanceAfter.Uint64() <= balanceBefore.Uint64() {
			t.Fatalf("did not send 1 eth to randAddr")
		}
	}
	EstimatedGas, err := safe_.EstimateTxGasByTrying(
		*randAddr,
		1e15,
		nil,
		0,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if EstimatedGas == 0 {
		t.Fatalf("could not estimate gas with eth node estimateGas method")
	}

	t.Log(EstimatedGas)
}

func TestGetMessageHash(t *testing.T) { // tested with official safe lib using the same Safe. OK!
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
	_, err := safe_.GetMessageHash("abc")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRetreiveFallbackHandler(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
	handler, err := safe_.RetrieveFallbackHandler()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if handler != eth.NULL_ADDRESS {
		t.Fatalf("default safe should not have any fallback handler set")
	}
}

func TestRetreiveMastercopyAddress(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
	mastercopyAddress, err := safe_.RetrieveMastercopyAddress()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if mastercopyAddress != network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address {
		t.Fatalf("mastercopy Address is wrong")
	}
}

func TestRetreiveIsHashApproved(t *testing.T) {
	if testSafe == eth.NULL_ADDRESS {
		err := createTestSafe()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	safe_ := New(testSafe, ethClient, privateKey)
	isApproved, err := safe_.RetrieveIsHashApproved(eth.NULL_ADDRESS, common.BigToHash(common.Big0).Bytes())
	if err != nil {
		t.Fatalf(err.Error())
	}
	if isApproved == true {
		t.Fatalf("hash should not be approved")
	}
}
