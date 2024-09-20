package safe

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	safecreations "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_creations"
	safe_types "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/types"
	safe_utils "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// keccak256("fallback_manager.handler.address")
const FALLBACK_HANDLER_STORAGE_SLOT = 0x6C9A6C4A39284E37ED1CF53D337577D14212A4870FB976A4366C693B939918D5

// keccak256("guard_manager.guard.address")
const GUARD_STORAGE_SLOT = 0x4A204F620C8C5CCDCA3FD54D003BADD85BA500436A431F0CBDA4F558C93C34C8

// keccak256("SafeMessage(bytes message)");
var SAFE_MESSAGE_TYPEHASH = common.Hex2Bytes("60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca")

type Safe struct {
	ethereumClient *eth.EthereumClient
	safeAddress    *common.Address
	safeContract   *contracts.GnosisSafe
	safeAbi        *abi.ABI
	proxyContract  *contracts.GnosisSafeProxy
	proxyAddress   *common.Address
}

func (safe *Safe) String() string {
	return "Safe=" + safe.safeAddress.Hex()
}

func New(safeAddress common.Address, ethClient *eth.EthereumClient) *Safe {
	/*
		Instantiates a new Safe object given an already created "Safe" (GnosisSafeProxy)
		This object will allow access to all the other required method to interact with the
		Smart Wallet (tx submit, ...)
	*/
	chainId, err := ethClient.GetChainId()
	if err != nil {
		return nil
	}
	masterCopyAddress := network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address
	if masterCopyAddress == eth.NULL_ADDRESS {
		return nil
	}
	var safeContract *contracts.GnosisSafe
	safeContract, err = contracts.NewGnosisSafe(masterCopyAddress, ethClient.GetGEthClient())
	if err != nil {
		safeContract = nil
	}
	var safeAbi *abi.ABI
	safeAbi, err = contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		safeAbi = nil
	}
	var proxyContract *contracts.GnosisSafeProxy
	proxyContract, err = contracts.NewGnosisSafeProxy(safeAddress, ethClient.GetGEthClient())
	if err != nil {
		proxyContract = nil
	}
	return &Safe{
		ethereumClient: ethClient,
		safeAddress:    &masterCopyAddress,
		safeContract:   safeContract,
		safeAbi:        safeAbi,
		proxyContract:  proxyContract,
		proxyAddress:   &safeAddress,
	}
}

func (safe *Safe) Version() (string, error) {
	GnosisSafe, err := contracts.NewGnosisSafe(*safe.safeAddress, safe.ethereumClient.GetGEthClient())
	if err != nil {
		return *new(string), err
	}
	version, err := GnosisSafe.VERSION(new(bind.CallOpts))
	return version, err
}

func (safe *Safe) DomainSeparator() (common.Hash, error) {
	GnosisSafe, err := contracts.NewGnosisSafe(*safe.safeAddress, safe.ethereumClient.GetGEthClient())
	if err != nil {
		return *new(common.Hash), err
	}
	domainSeparator, err := GnosisSafe.DomainSeparator(new(bind.CallOpts))
	return common.BytesToHash(domainSeparator[:]), err
}

func Create(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
	masterCopyAddress common.Address,
	owners []common.Address,
	threshold int64,
	fallbackHandler common.Address,
	proxyFactoryAddress common.Address,
	paymentToken common.Address,
	payment int64,
	paymentReceiver common.Address,
) (safe_types.EthereumTxSent, error) {
	/*
		Creates a new Gnosis Safe Wallet (deploys a new Gnosis Safe Proxy)
	*/
	/* owners checks */
	if len(owners) <= 0 {
		return *new(safe_types.EthereumTxSent), fmt.Errorf("at least one owner must be set")
	}
	if !(threshold >= 1 && int(threshold) <= len(owners)) {
		return *new(safe_types.EthereumTxSent), fmt.Errorf("threshold=%d must be <= %d", threshold, len(owners))
	}

	/* Get required information for tx building */
	nonce, chainId, _, gasPrice, err := safe_utils.GetDefaultTxParams(ethereumClient, sender)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	/* construct the initialization data needed for the proxy to initialize the Safe */
	safeAbi, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	method := safeAbi.Methods["setup"].Name
	initializer, err := safeAbi.Pack(
		method,
		owners,
		big.NewInt(threshold),
		eth.NULL_ADDRESS, // Contract address for optional delegate call
		make([]byte, 1),  // Data payload for optional delegate call
		fallbackHandler,
		paymentToken,
		big.NewInt(payment),
		paymentReceiver,
	)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	/* retrieve the ProxyFactory contract and deploy the new Proxy */
	proxyFactory, err := contracts.NewGnosisSafeProxyFactory(
		network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address,
		ethereumClient.GetGEthClient(),
	)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	auth, err := safe_utils.BuildTransactionWithSigner(sender, privateKey, common.Big0, int64(chainId), int64(nonce), gasPrice, nil)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	auth.GasLimit = uint64(300000) // in units

	tx, err := proxyFactory.CreateProxy(
		auth,
		masterCopyAddress,
		initializer,
	)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	ch := ethereumClient.WaitTxConfirmed(tx.Hash())
	isPending := <-ch
	if isPending {
		return *new(safe_types.EthereumTxSent), fmt.Errorf("safe creation transaction still pending. hash %s", tx.Hash().Hex())
	}
	receipt, err := ethereumClient.GetReceipt(tx.Hash().String())
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	if !eth.IsTransactionSuccessful(receipt) {
		return *new(safe_types.EthereumTxSent), fmt.Errorf("safe creation FAILED. hash %s", tx.Hash().Hex())
	}
	proxyAddress, err := safe_utils.GetProxyCreationResult(proxyFactory, receipt)
	if err != nil {
		return *new(safe_types.EthereumTxSent), fmt.Errorf("safe creation FAILED. proxyAddress not found in receipt")
	}
	return safe_types.EthereumTxSent{
		Tx:              tx,
		ContractAddress: proxyAddress,
		TxHash:          receipt.TxHash,
	}, nil
}

func CreateWithNonce(
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
	safeCreateTx2 *safecreations.SafeCreationTx2,
) (*safe_types.EthereumTxSent, error) {
	/* Get required information for tx building */
	nonce, chainId, _, gasPrice, err := safe_utils.GetDefaultTxParams(
		safeCreateTx2.EthereumClient,
		sender,
	)
	if err != nil {
		return nil, err
	}

	/* Calculate the address that will be deployed to see if it matches with the real one */
	if safeCreateTx2.ExpectedSafeAddress2 == eth.NULL_ADDRESS {
		safeCreateTx2.PredictSafeAddress_CREATE2()
	}

	/* retrieve the ProxyFactory contract and deploy the new Proxy */
	proxyFactory, err := contracts.NewGnosisSafeProxyFactory(
		network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address,
		safeCreateTx2.EthereumClient.GetGEthClient(),
	)
	if err != nil {
		return nil, err
	}

	auth, err := safe_utils.BuildTransactionWithSigner(
		sender, privateKey, nil, int64(chainId), int64(nonce), gasPrice, nil)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = safeCreateTx2.CreationGas + 50000

	tx, err := proxyFactory.CreateProxyWithNonce(
		auth,
		safeCreateTx2.MasterCopy,
		safeCreateTx2.GetInitializer2(),
		big.NewInt(safeCreateTx2.SaltNonce),
	)
	if err != nil {
		return nil, err
	}
	ch := safeCreateTx2.EthereumClient.WaitTxConfirmed(tx.Hash())
	isPending := <-ch
	if isPending {
		return nil, fmt.Errorf("safe creation transaction still pending. hash %s", tx.Hash().Hex())
	}
	receipt, err := safeCreateTx2.EthereumClient.GetReceipt(tx.Hash().String())
	if err != nil {
		return nil, err
	}
	if !eth.IsTransactionSuccessful(receipt) {
		return nil, fmt.Errorf("safe creation FAILED. hash %s", tx.Hash().Hex())
	}
	proxyAddress, err := safe_utils.GetProxyCreationResult(proxyFactory, receipt)
	if err != nil {
		return nil, fmt.Errorf("safe creation FAILED. proxyAddress not found in receipt")
	}

	ethTxSent := safe_types.EthereumTxSent{
		Tx:              tx,
		ContractAddress: proxyAddress,
		TxHash:          receipt.TxHash,
	}
	if safeCreateTx2.ExpectedSafeAddress2 != proxyAddress {
		return &ethTxSent, fmt.Errorf("predicted Safe address differs from deployed one")
	}
	return &ethTxSent, nil
}

func DeployMasterContract_v1_3_0(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
) (safe_types.EthereumTxSent, error) {
	/*
		Deploys a new v1.3.0 Gnosis Safe Master Copy
	*/
	return safe_utils.DeployMasterContract(ethereumClient, sender, privateKey, nil)
}

func DeployCompatibilityFallbackHandler(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
) (safe_types.EthereumTxSent, error) {
	/*
		Deploy Compatibility Fallback handler v1.3.0
	*/
	nonce, chainId, estimatedEIP1559Gas, gasPrice, err := safe_utils.GetDefaultTxParams(ethereumClient, sender)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	auth, err := safe_utils.BuildTransactionWithSigner(sender, privateKey, common.Big0, int64(chainId), int64(nonce), gasPrice, estimatedEIP1559Gas)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	contractAddress, tx, fallbackHandler, err := contracts.DeployCompatibiliyFallbackHandler(auth, ethereumClient.GetGEthClient())
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	if fallbackHandler == nil {
		return *new(safe_types.EthereumTxSent), err
	}
	return safe_types.EthereumTxSent{
		Tx:              tx,
		ContractAddress: contractAddress,
		TxHash:          tx.Hash(),
	}, nil
}

func EstimateSafeCreation(
	ethereumClient *eth.EthereumClient,
	owners []common.Address, // Owners of the Safe
	threshold int64, // Minimum number of users required to operate the Safe
	funder common.Address, // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken common.Address, // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64, // Value of payment token per 1 Ether
	fixedCreationCost int, // Fixed creation cost of Safe (Wei)
) (uint64, int64, uint64, error) {
	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return 0, 0, 0, err
	}
	gasPrice, err := ethereumClient.GasPrice()
	if err != nil {
		return 0, 0, 0, err
	}
	safeCreationTx, err := safecreations.NewSafeCreationTx(
		ethereumClient,
		owners,
		threshold,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		funder,
		paymentToken,
		paymentTokenEthValue,
		fixedCreationCost,
	)
	if err != nil {
		return 0, 0, 0, err
	}
	safeCreationTx.EstimateSafeCreation()
	return safeCreationTx.CreationGas, gasPrice.Int64(), safeCreationTx.Payment, err
}
