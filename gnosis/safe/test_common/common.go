package testcommon

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/safe"
	safecreations "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_creations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/joho/godotenv"
)

func envPath() string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../../")
	return filepath.Join(root, ".env")
}

func loadEnv() {
	env := envPath()
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func CreateTestSafe() (*common.Address, *eth.EthereumClient, *ecdsa.PrivateKey) {
	loadEnv()

	ethClient := EthClient()
	sender, sKey := SenderAccount()

	defaultSafeAddress := os.Getenv("DEFAULT_SAFE_ADDRESS")
	isContract, err := ethClient.IsContract(defaultSafeAddress)
	if err != nil {
		panic(err)
	}
	if !isContract {
		fmt.Println("Creating new Safe")
		safeAddr := deployPredicatbleSafe(ethClient, sender, sKey)
		return safeAddr, ethClient, sKey
	}
	fmt.Println("Default Safe already deployed...")
	safeAddr := common.HexToAddress(defaultSafeAddress)
	return &safeAddr, ethClient, sKey
}

func deployPredicatbleSafe(ethClient *eth.EthereumClient, sender *common.Address, sKey *ecdsa.PrivateKey) *common.Address {
	owner2 := Owner2()
	owner3 := Owner3()

	chainId, _ := ethClient.GetChainId()

	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, *owner2)
	owners = append(owners, *owner3)
	safeCreationTx2, err := safecreations.NewSafeCreationTx2(
		ethClient,
		owners,
		2,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS, 1.0, 0, 1,
	)
	if err != nil {
		panic(err)
	}
	safeCreationTx2.PredictSafeAddress_CREATE2()
	if safeCreationTx2.ExpectedSafeAddress2 == eth.NULL_ADDRESS {
		panic("Could not predict Safe address")
	}

	/* If funder is set fund the future Safe address so that it can pay back the funder, otherwise deploy will fail */
	if safeCreationTx2.Funder != eth.NULL_ADDRESS {
		gasPrice, _ := ethClient.GasPrice()
		txHash, err := ethClient.SendEthTo(
			hexutil.EncodeBig(sKey.D), &safeCreationTx2.ExpectedSafeAddress2, gasPrice, big.NewInt(1e18), 21000)
		if err != nil {
			panic(err.Error())
		}
		ch := ethClient.WaitTxConfirmed(txHash)
		isPending := <-ch
		if isPending {
			panic("safe creation transaction still pending")
		}
	}

	/* Now deploy a new Safe with the init params and Nonce provided in the safeCreationTx2 obj */
	// contract address check is made inside the func so we check only if there error
	ethTxSent, err := safe.CreateWithNonce(*sender, sKey, safeCreationTx2)
	if err != nil {
		panic(err.Error())
	}

	/* Check that the actual deployed Safe coincides with the predicted one */
	if ethTxSent.ContractAddress != safeCreationTx2.ExpectedSafeAddress2 {
		panic("Deployed Safe address differs from predicted Safe address")
	}
	fmt.Println(ethTxSent.ContractAddress)
	return &ethTxSent.ContractAddress
}

func SenderAccount() (*common.Address, *ecdsa.PrivateKey) {
	loadEnv()
	HARDHAT_S_KEY0 := os.Getenv("HARDHAT_S_KEY0")
	sender, _ := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	sKey, _ := eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
	return sender, sKey
}

func Sender() *common.Address {
	loadEnv()
	HARDHAT_S_KEY0 := os.Getenv("HARDHAT_S_KEY0")
	sender, _ := eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))
	return sender
}

func EthClient() *eth.EthereumClient {
	loadEnv()
	ethereumNodeUrl := os.Getenv("ETHEREUM_NODE_URL")
	if ethereumNodeUrl == "" {
		panic("empty ethereum node url")
	}
	ethClient, _ := eth.EthereumClientInit(eth.NewURI(ethereumNodeUrl))
	return ethClient
}

func Owner2() *common.Address {
	loadEnv()
	owner2 := common.HexToAddress(os.Getenv("OWNER_2"))
	return &owner2
}

func Owner3() *common.Address {
	loadEnv()
	owner3 := common.HexToAddress(os.Getenv("OWNER_3"))
	return &owner3
}
