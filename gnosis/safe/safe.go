package safe

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	masterCopyAddress := network.NetworkToSafeAddress[network.Gochain_Testnet].Address
	var safeContract *contracts.GnosisSafe
	safeContract, err := contracts.NewGnosisSafe(masterCopyAddress, ethClient.GetGEthClient())
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
) (EthereumTxSent, error) {
	/* owners checks */
	if len(owners) <= 0 {
		return *new(EthereumTxSent), fmt.Errorf("at least one owner must be set")
	}
	if !(threshold >= 1 && int(threshold) <= len(owners)) {
		return *new(EthereumTxSent), fmt.Errorf("threshold=%d must be <= %d", threshold, len(owners))
	}

	/* Get required information for tx building */
	nonce, err := ethereumClient.GetNonceForAccount(sender, "pending")
	if err != nil {
		return *new(EthereumTxSent), err
	}
	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return *new(EthereumTxSent), err
	}
	gasPrice, err := ethereumClient.GasPrice()
	if err != nil {
		return *new(EthereumTxSent), err
	}

	/* construct the initialization data needed for the proxy to initialize the Safe */
	safeAbi, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		return *new(EthereumTxSent), err
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
		return *new(EthereumTxSent), err
	}

	/* retrieve the ProxyFactory contract and deploy the new Proxy */
	proxyFactory, err := contracts.NewGnosisSafeProxyFactory(
		network.NetworkToSafeProxyFactoryAddress[network.Gochain_Testnet].Address,
		ethereumClient.GetGEthClient(),
	)
	if err != nil {
		return *new(EthereumTxSent), err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey, big.NewInt(int64(chainId)),
	)
	if err != nil {
		return *new(EthereumTxSent), err
	}
	auth.From = sender
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := proxyFactory.CreateProxy(
		auth,
		masterCopyAddress,
		initializer,
	)
	if err != nil {
		return *new(EthereumTxSent), err
	}
	ch := ethereumClient.WaitTxConfirmed(tx.Hash())
	isPending := <-ch
	if isPending {
		return *new(EthereumTxSent), fmt.Errorf("safe creation transaction still pending. hash %s", tx.Hash().Hex())
	}
	receipt, err := ethereumClient.GetReceipt(tx.Hash().String())
	if err != nil {
		return *new(EthereumTxSent), err
	}
	if !eth.IsTransactionSuccessful(receipt) {
		return *new(EthereumTxSent), fmt.Errorf("safe creation FAILED. hash %s", tx.Hash().Hex())
	}
	proxyAddress, err := getProxyCreationResult(proxyFactory, receipt)
	if err != nil {
		return *new(EthereumTxSent), fmt.Errorf("safe creation FAILED. proxyAddress not found in receipt")
	}
	return EthereumTxSent{
		tx:               tx,
		contractAaddress: proxyAddress,
		TxHash:           receipt.TxHash,
	}, nil
}

func getProxyCreationResult(proxyFactory *contracts.GnosisSafeProxyFactory, receipt *types.Receipt) (common.Address, error) {
	var result *contracts.GnosisSafeProxyFactoryProxyCreation
	for _, log := range receipt.Logs {
		result, _ = proxyFactory.ParseProxyCreation(*log)
	}
	if result == nil {
		return common.Address{}, nil
	}
	return result.Proxy, nil
}
