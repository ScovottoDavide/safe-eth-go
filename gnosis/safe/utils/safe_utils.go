package utils

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	safe_types "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/**
	PRIVATE METHODS
**/

func GetProxyCreationResult(proxyFactory *contracts.GnosisSafeProxyFactory, receipt *types.Receipt) (common.Address, error) {
	/*
		Get the address of the newly deployed GnosisSafeProxy from the receipt (the address is returned by an event)
		The ParseProxyCreation catches the following event:
			`event ProxyCreation(GnosisSafeProxy proxy, address singleton);`
	*/
	var result *contracts.GnosisSafeProxyFactoryProxyCreation
	for _, log := range receipt.Logs {
		result, _ = proxyFactory.ParseProxyCreation(*log)
	}
	if result == nil {
		return common.Address{}, nil
	}
	return result.Proxy, nil
}

func DeployMasterContract(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
	constructorData []byte, // for Safe version < 1.1.1
) (safe_types.EthereumTxSent, error) {
	/*
		Private method for deploying a new Gnosis Safe Master Copy. For only version 1.3.0 is tested.
	*/
	var _ = constructorData
	nonce, chainId, estimatedEIP1559Gas, gasPrice, err := GetDefaultTxParams(ethereumClient, sender)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	auth, err := BuildTransactionWithSigner(sender, privateKey, int64(chainId), int64(nonce), gasPrice, estimatedEIP1559Gas)
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}

	contractAddress, tx, gnosisSafe, err := contracts.DeployGnosisSafe(auth, ethereumClient.GetGEthClient())
	if err != nil {
		return *new(safe_types.EthereumTxSent), err
	}
	if gnosisSafe == nil {
		return *new(safe_types.EthereumTxSent), err
	}
	return safe_types.EthereumTxSent{
		Tx:              tx,
		ContractAddress: contractAddress,
		TxHash:          tx.Hash(),
	}, nil
}

func GetDefaultTxParams(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
) (uint64, int, *eth.EIP1559EstimatedGas, *big.Int, error) {
	nonce, err := ethereumClient.GetNonceForAccount(sender, "pending")
	if err != nil {
		return 0, 0, nil, nil, err
	}
	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return 0, 0, nil, nil, err
	}

	estimatedEIP1559Gas := &eth.EIP1559EstimatedGas{GasTipCap: nil, GasFeeCap: nil}
	var gasPrice *big.Int
	if ethereumClient.IsEip1559Supported() {
		estimatedEIP1559Gas, err = ethereumClient.EstimateFeeEip1559(network.Normal)
		if err != nil {
			return 0, 0, nil, nil, err
		}
	} else {
		gasPrice, err = ethereumClient.GasPrice()
		if err != nil {
			return 0, 0, nil, nil, err
		}
	}
	return nonce, chainId, estimatedEIP1559Gas, gasPrice, nil
}

func BuildTransactionWithSigner(
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
	chainId int64,
	nonce int64,
	gasPrice *big.Int,
	estimatedEIP1559Gas *eth.EIP1559EstimatedGas,
) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey, big.NewInt(int64(chainId)),
	)
	if err != nil {
		return nil, err
	}
	auth.From = sender
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = common.Big0
	if estimatedEIP1559Gas != nil {
		auth.GasFeeCap = estimatedEIP1559Gas.GasFeeCap
		auth.GasTipCap = estimatedEIP1559Gas.GasTipCap
	} else {
		auth.GasPrice = gasPrice
	}
	return auth, nil
}
