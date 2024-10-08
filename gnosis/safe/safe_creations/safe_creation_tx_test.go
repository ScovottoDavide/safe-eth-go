package safecreations_test

import (
	"math/big"
	"math/rand/v2"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/safe"
	safecreations "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_creations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const owner2 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const owner3 = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

var uri = eth.NewURI("http://127.0.0.1:8546")
var ethClient, _ = eth.EthereumClientInit(uri)
var chainId, _ = ethClient.GetChainId()

func TestEstimateSafeCreation(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	safeCreationTx, err := safecreations.NewSafeCreationTx(
		ethClient,
		owners,
		2,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS, 1.0, 0,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = safeCreationTx.EstimateSafeCreation()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(safeCreationTx.CreationGas)
	t.Log(safeCreationTx.Payment)
	/* SafeCreationEstimate from safe-eth-py (gas=394016, gas_price=1875000000, payment=781905000000000,
	payment_token='0x0000000000000000000000000000000000000000')*/
}

func TestPredictSafeAddressWithCREATE(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))

	safeCreationTx, err := safecreations.NewSafeCreationTx(
		ethClient,
		owners,
		2,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS, 1.0, 0,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	safeCreationTx.PredictSafeAddress_CREATE()

	txSent, err := safe.Create(
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
		t.Fatalf(err.Error())
	}

	if txSent.ContractAddress.Hex() != safeCreationTx.ExpectedSafeAddress.Hex() {
		t.Fatalf("Deployed Safe address differs from Predicted Safe address")
	}
}

func TestEstimateSafeCreation2(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	safeCreationTx2, err := safecreations.NewSafeCreationTx2(
		ethClient,
		owners,
		2,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS, 1.0, 0, rand.Int64(),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = safeCreationTx2.EstimateSafeCreation2()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// 307874 -> got from running the official safe-eth-py with the same setup parameters
	if safeCreationTx2.CreationGas != 307874 {
		t.Fatalf("CreationGas differs from what expected for this initializer")
	}
}

func TestPredictSafeAddressWithCREATE2(t *testing.T) {
	sender, err := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	if err != nil {
		t.Fatalf(err.Error())
	}
	privateKey, err := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))

	/* Create a new SafeTx2 ->
	it will estimate creation gas and the eventual refund payment and prepare the initializer
	*/
	safeCreationTx2, err := safecreations.NewSafeCreationTx2(
		ethClient,
		owners,
		2,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS, 1.0, 0, rand.Int64(),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	/*it will calculate the Safe address that will be deployed given the provided params and SaltNonce*/
	safeCreationTx2.PredictSafeAddress_CREATE2()
	if safeCreationTx2.ExpectedSafeAddress2 == eth.NULL_ADDRESS {
		t.Fatalf("Could not predict Safe address")
	}

	/* If funder is set fund the future Safe address so that it can pay back the funder, otherwise deploy will fail */
	if safeCreationTx2.Funder != eth.NULL_ADDRESS {
		gasPrice, _ := ethClient.GasPrice()
		txHash, err := ethClient.SendEthTo(
			hexutil.EncodeBig(privateKey.D), &safeCreationTx2.ExpectedSafeAddress2, gasPrice, big.NewInt(1e18), 21000)
		if err != nil {
			t.Fatalf(err.Error())
		}
		ch := ethClient.WaitTxConfirmed(txHash)
		isPending := <-ch
		if isPending {
			t.Fatalf("safe creation transaction still pending. hash %s", txHash)
		}
	}

	/* Now deploy a new Safe with the init params and Nonce provided in the safeCreationTx2 obj */
	// contract address check is made inside the func so we check only if there error
	ethTxSent, err := safe.CreateWithNonce(*sender, privateKey, safeCreationTx2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	/* Check that the actual deployed Safe coincides with the predicted one */
	if ethTxSent.ContractAddress != safeCreationTx2.ExpectedSafeAddress2 {
		t.Fatalf("Deployed Safe address differs from predicted Safe address")
	}
}
