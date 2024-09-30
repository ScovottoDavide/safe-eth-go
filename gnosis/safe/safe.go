package safe

import (
	"crypto/ecdsa"
	"fmt"
	"math"
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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// keccak256("fallback_manager.handler.address")
var FALLBACK_HANDLER_STORAGE_SLOT = common.HexToHash("0x6C9A6C4A39284E37ED1CF53D337577D14212A4870FB976A4366C693B939918D5")

// keccak256("guard_manager.guard.address")
var GUARD_STORAGE_SLOT = common.HexToHash("0x4A204F620C8C5CCDCA3FD54D003BADD85BA500436A431F0CBDA4F558C93C34C8")

// keccak256("SafeMessage(bytes message)");
var SAFE_MESSAGE_TYPEHASH = common.HexToHash("0x60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca")

type Safe struct {
	EthereumClient    *eth.EthereumClient
	MasterCopyAddress *common.Address
	SafeContract      *contracts.GnosisSafe
	safeAbi           *abi.ABI
	SafeAddress       *common.Address
	defualtSafeSigner *bind.SignerFn
	defualtSignerAddr *common.Address
}

func (safe *Safe) String() string {
	return "Safe=" + safe.MasterCopyAddress.Hex()
}

// Initializes a new Safe instance based on an already created Safe address
func New(SafeAddress common.Address, ethClient *eth.EthereumClient, signerSKey *ecdsa.PrivateKey) *Safe {
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
	safeContract, err = contracts.NewGnosisSafe(SafeAddress, ethClient.GetGEthClient())
	if err != nil {
		safeContract = nil
	}
	var safeAbi *abi.ABI
	safeAbi, err = contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		safeAbi = nil
	}

	signer, err := bind.NewKeyedTransactorWithChainID(signerSKey, big.NewInt(int64(chainId)))
	if err != nil {
		return nil
	}
	signerAddr := crypto.PubkeyToAddress(signerSKey.PublicKey)
	return &Safe{
		EthereumClient:    ethClient,
		MasterCopyAddress: &masterCopyAddress,
		SafeContract:      safeContract,
		safeAbi:           safeAbi,
		SafeAddress:       &SafeAddress,
		defualtSafeSigner: &signer.Signer,
		defualtSignerAddr: &signerAddr,
	}
}

// Deploys a new Safe (GnosisSafeProxy) and atomically calls the setup function of the MasterCopy
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

// Deploys a new Safe using the CREATE2 OPCODE based on the parameters of the SafeCreationTx2
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

// Deploys a new v1.3.0 Gnosis Safe Master Copy
func DeployMasterContract_v1_3_0(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	privateKey *ecdsa.PrivateKey,
) (safe_types.EthereumTxSent, error) {
	return safe_utils.DeployMasterContract(ethereumClient, sender, privateKey, nil)
}

// Deploys the Fallback Handler for maintaining compatibility with Safe older versions
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

// Estimates the gas costs needed for a new Safe
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

// Estimates the gas costs needed for a new Safe using the CREATE2 OPCODE
func EstimateSafeCreation2(
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
	safeCreationTx2, err := safecreations.NewSafeCreationTx2(
		ethereumClient,
		owners,
		threshold,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		eth.NULL_ADDRESS,
		funder,
		paymentToken,
		paymentTokenEthValue,
		fixedCreationCost,
		0,
	)
	if err != nil {
		return 0, 0, 0, err
	}
	safeCreationTx2.EstimateSafeCreation2()
	return safeCreationTx2.CreationGas, gasPrice.Int64(), safeCreationTx2.Payment, err
}

// Builds a new SafeCreationTx. It allows to Predict the Safe address that will be deployed, to estimate the creation cost, ...
func BuildSafeCreationTx(
	ethereumClient *eth.EthereumClient,
	owners []common.Address, // Owners of the Safe
	threshold int64, // Minimum number of users required to operate the Safe
	masterCopyAddress common.Address,
	fallbackHandler common.Address,
	funder common.Address, // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken common.Address, // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64, // Value of payment token per 1 Ether
	fixedCreationCost int, // Fixed creation cost of Safe (Wei)
) (*safecreations.SafeCreationTx, error) {
	return safecreations.NewSafeCreationTx(
		ethereumClient,
		owners,
		threshold,
		masterCopyAddress,
		fallbackHandler,
		funder,
		paymentToken,
		paymentTokenEthValue,
		fixedCreationCost,
	)
}

// Builds a new SafeCreationTx2. It allows to Predict the Safe address that will be deployed, to estimate the creation cost, ...
func BuildSafeCreationTx2(
	ethereumClient *eth.EthereumClient,
	owners []common.Address, // Owners of the Safe
	threshold int64, // Minimum number of users required to operate the Safe
	masterCopyAddress common.Address,
	fallbackHandler common.Address,
	funder common.Address, // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken common.Address, // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64, // Value of payment token per 1 Ether
	fixedCreationCost int, // Fixed creation cost of Safe (Wei)
	saltNonce int64,
) (*safecreations.SafeCreationTx2, error) {
	return safecreations.NewSafeCreationTx2(
		ethereumClient,
		owners,
		threshold,
		masterCopyAddress,
		fallbackHandler,
		funder,
		paymentToken,
		paymentTokenEthValue,
		fixedCreationCost,
		saltNonce,
	)
}

// Check Safe has enough funds to pay for tx
func (safe_ *Safe) CheckFundsForTxGas(
	safeTxGas uint64,
	baseGas uint64,
	gasPrice int64,
	gasToken common.Address,
) (bool, error) {
	if gasToken == eth.NULL_ADDRESS {
		balance, err := safe_.EthereumClient.GetBalance(safe_.SafeAddress)
		if err != nil {
			return false, err
		}
		return balance.Uint64() >= (safeTxGas+baseGas)*uint64(gasPrice), nil
	} else {
		gasTokenContract, err := contracts.NewERC20(gasToken, safe_.EthereumClient.GetGEthClient())
		if err != nil {
			return false, nil
		}
		balance, err := gasTokenContract.BalanceOf(nil, *safe_.SafeAddress)
		if err != nil {
			return false, nil
		}
		return balance.Uint64() >= (safeTxGas+baseGas)*uint64(gasPrice), nil
	}
}

// Calculate gas costs that are independent of the transaction execution(e.g. base transaction fee,
// signature check, payment of the refund...)
func (safe_ *Safe) EstimateTxBaseGas(
	to common.Address,
	value uint64,
	data []byte,
	operation int,
	gas_token common.Address,
	estimated_tx_gas uint64, // gas calculated with `estimate_tx_gas`
) (uint64, error) {
	nonce, err := safe_.RetrieveNonce()
	if err != nil {
		return 0, err
	}
	threshold, err := safe_.RetrieveThreshold()
	if err != nil {
		return 0, err
	}
	// Every byte == 0 -> 4  Gas
	// Every byte != 0 -> 16 Gas (68 before Istanbul)
	// numbers < 256 (0x00(31*2)..ff) are 192 -> 31 * 4 + 1 * GAS_CALL_DATA_BYTE
	// numbers < 65535 (0x(30*2)..ffff) are 256 -> 30 * 4 + 2 * GAS_CALL_DATA_BYTE
	// Calculate gas for signatures
	// (array count (3 -> r, s, v) + ecrecover costs) * signature count
	// ecrecover for ecdsa ~= 4K gas, we use 6K

	ecrecover_gas := 6000
	signature_gas := int(threshold.Uint64()) * (1*eth.GAS_CALL_DATA_BYTE + 2*32*eth.GAS_CALL_DATA_BYTE + ecrecover_gas)
	safe_tx_gas := estimated_tx_gas
	base_gas := 0
	gas_price := 1
	signatures := make([]byte, 0)
	refund_receiver := eth.NULL_ADDRESS

	method := safe_.safeAbi.Methods["execTransaction"].Name
	initializer, err := safe_.safeAbi.Pack(
		method,
		to,
		value,
		data,
		operation,
		safe_tx_gas,
		base_gas,
		gas_price,
		gas_token,
		refund_receiver,
		signatures,
	)
	if err != nil {
		return 0, err
	}
	// If nonce == 0, nonce storage has to be initialized
	nonce_gas := 5000
	if nonce == common.Big0 {
		nonce_gas = 20000
	}
	// Keccak costs for the hash of the safe tx
	hash_generation_gas := 1500

	base_gas = signature_gas + int(safe_.EthereumClient.EstimateDataGas(initializer)) + nonce_gas + hash_generation_gas

	// Add additional gas costs
	if base_gas > 65536 {
		base_gas += 64
	} else {
		base_gas += 128
	}

	base_gas += 32000 // Base tx costs, transfer costs...
	return uint64(base_gas), nil
}

// Estimate tx gas using safe `requiredTxGas` method
func (safe_ *Safe) EstimateTxGasWithSafe(
	to common.Address,
	value uint64,
	data []byte,
	operation uint8,
	gasLimit uint64,
	signer bind.SignerFn,
) (uint64, error) {
	from := *safe_.SafeAddress
	gasPrice := common.Big0
	if gasLimit > 0 {
		from = *safe_.defualtSignerAddr
		gasPrice, _ = safe_.EthereumClient.GasPrice()
	}
	_, err := safe_.SafeContract.RequiredTxGas(&bind.TransactOpts{
		From:     from,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Signer:   signer,
	}, to, big.NewInt(int64(value)), data, operation)
	if err != nil {
		jsonErr, ok := err.(safe_types.JsonError)
		if ok {
			// it is a Json Error
			errData := jsonErr.ErrorData().(string)
			// 4 bytes - error method id
			// 32 bytes - position
			// 32 bytes - length
			// Last 32 bytes - value of revert (if everything went right)
			gas_estimation_offset := 4 + 32 + 32
			revertBytes, _ := hexutil.Decode(errData)
			gas_estimation := revertBytes[gas_estimation_offset:]
			if len(gas_estimation) != 32 {
				return 0, fmt.Errorf("cannot estimate tx gas with safe")
			}
			estimatedGas := new(big.Int)
			estimatedGas = estimatedGas.SetBytes(gas_estimation)
			return estimatedGas.Uint64(), nil
		} else {
			return 0, err
		}
	}
	return 0, err
}

func (safe_ *Safe) EstimateTxGasWithWeb3(
	to common.Address,
	value uint64,
	data []byte,
) (uint64, error) {
	return safe_.EthereumClient.EstimateGas(
		*safe_.SafeAddress,
		&to,
		0,
		nil, nil, nil,
		big.NewInt(int64(value)), data,
		nil, nil, nil,
	)
}

// Try to get an estimation with Safe's `requiredTxGas`. If estimation if successful, try to set a gas limit and
// estimate again. If gas estimation is ok, same gas estimation should be returned, if it's less than required
// estimation will not be completed, so estimation was not accurate and gas limit needs to be increased.
func (safe_ *Safe) EstimateTxGasByTrying(
	to common.Address,
	value uint64,
	data []byte,
	operation uint8,
) (uint64, error) {
	gasEstimated, err := safe_.EstimateTxGasWithSafe(to, value, data, operation, 0, nil)
	if err != nil {
		return 0, err
	}
	baseGas := safe_.EthereumClient.EstimateDataGas(data)

	for i := range 30 {
		_, err = safe_.EstimateTxGasWithSafe(to, value, data, operation, gasEstimated+baseGas+32000, *safe_.defualtSafeSigner)
		if err != nil {
			blockNum, _ := safe_.EthereumClient.CurrentBlockNumber()
			_, header, err := safe_.EthereumClient.GetBlockByNumber(big.NewInt(int64(blockNum)), false)
			if err != nil {
				return 0, nil
			}
			blockGasLimit := header.GasLimit
			gasEstimated = uint64(math.Floor((1 + float64(i)*0.03) * float64(gasEstimated)))
			if gasEstimated >= blockGasLimit {
				return blockGasLimit, nil
			}
		}
	}
	return gasEstimated, nil
}

// Estimate tx gas. Use `RequiredTxGas` on the Safe contract and fallbacks to `eth_estimateGas` if that method
// fails. Note: `eth_estimateGas` cannot estimate delegate calls
func (safe_ *Safe) EstimateTxGas( // TODO: test this
	to common.Address,
	value uint64,
	data []byte,
	operation uint8,
) (uint64, error) {
	// Costs to route through the proxy and nested calls
	PROXY_GAS := 1000
	// https://github.com/ethereum/solidity/blob/dfe3193c7382c80f1814247a162663a97c3f5e67/libsolidity/codegen/ExpressionCompiler.cpp#L1764
	// This was `false` before solc 0.4.21 -> `m_context.evmVersion().canOverchargeGasForCall()`
	// So gas needed by caller will be around 35k
	OLD_CALL_GAS := 35000
	// Web3 `estimate_gas` estimates less gas
	WEB3_ESTIMATION_OFFSET := 23000
	ADDITIONAL_GAS := PROXY_GAS + OLD_CALL_GAS

	estimatedGas, err := safe_.EstimateTxGasByTrying(
		to, value, data, operation,
	)
	if err != nil {
		web3EstimateGas, err := safe_.EstimateTxGasWithWeb3(
			to, value, data,
		)
		if err != nil {
			return 0, err
		}
		return web3EstimateGas + uint64(ADDITIONAL_GAS) + uint64(WEB3_ESTIMATION_OFFSET), nil
	}
	return estimatedGas, nil
}

// Return hash of a message that can be signed by owners.
func (safe_ *Safe) GetMessageHash(message interface{}) (*common.Hash, error) {
	messageHash := *new(common.Hash)
	switch m := message.(type) {
	case string:
		messageHash = crypto.Keccak256Hash([]byte(m))
	case []byte:
		messageHash = crypto.Keccak256Hash(m)
	}
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
	}
	messageHashPack, err := arguments.Pack(SAFE_MESSAGE_TYPEHASH, messageHash)
	if err != nil {
		return nil, err
	}
	safeMessageHash := crypto.Keccak256Hash(messageHashPack)
	domainSeparator, err := safe_.DomainSeparator()
	if err != nil {
		return nil, err
	}
	res := crypto.Keccak256Hash(
		[]byte{0x19},
		[]byte{0x01},
		domainSeparator.Bytes(),
		safeMessageHash.Bytes(),
	)
	return &res, nil
}

func (safe_ *Safe) RetrieveAllInfo() {
	// TODO:
}

// Retrieve Safe Version
func (safe *Safe) Version() (string, error) {
	return safe.SafeContract.VERSION(new(bind.CallOpts))
}

// Retrieve Safe Domain Separator
func (safe *Safe) DomainSeparator() (common.Hash, error) {
	GnosisSafe, err := contracts.NewGnosisSafe(*safe.SafeAddress, safe.EthereumClient.GetGEthClient())
	if err != nil {
		return *new(common.Hash), err
	}
	domainSeparator, err := GnosisSafe.DomainSeparator(new(bind.CallOpts))
	return common.BytesToHash(domainSeparator[:]), err
}

// Retrieve the threshold from the Safe
func (safe_ *Safe) RetrieveThreshold() (*big.Int, error) {
	threshold, err := safe_.SafeContract.GetThreshold(nil)
	if err != nil {
		return nil, err
	}
	return threshold, err
}

// Retrieve the Safe Nonce
func (safe_ *Safe) RetrieveNonce() (*big.Int, error) {
	nonce, err := safe_.SafeContract.Nonce(nil)
	if err != nil {
		return nil, err
	}
	return nonce, err
}

func (safe_ *Safe) GetCode() ([]byte, error) {
	return safe_.EthereumClient.GetCode(safe_.SafeAddress)
}

func (safe_ *Safe) RetrieveFallbackHandler() (common.Address, error) {
	storageSlot, err := safe_.EthereumClient.GetStorageAt(safe_.SafeAddress, FALLBACK_HANDLER_STORAGE_SLOT)
	if err != nil {
		return eth.NULL_ADDRESS, err
	}
	return common.BytesToAddress(storageSlot[len(storageSlot)-20:]), nil
}

func (safe_ *Safe) RetrieveGuard() (common.Address, error) {
	storageSlot, err := safe_.EthereumClient.GetStorageAt(safe_.SafeAddress, GUARD_STORAGE_SLOT)
	if err != nil {
		return eth.NULL_ADDRESS, err
	}
	return common.BytesToAddress(storageSlot[len(storageSlot)-20:]), nil
}

func (safe_ *Safe) RetrieveMastercopyAddress() (common.Address, error) {
	storageSlot, err := safe_.EthereumClient.GetStorageAt(safe_.SafeAddress, common.HexToHash("0x00"))
	if err != nil {
		return eth.NULL_ADDRESS, err
	}
	return common.BytesToAddress(storageSlot[len(storageSlot)-20:]), nil
}

func (safe_ *Safe) RetreiveModules() {
	// TODO:
}

func (safe_ *Safe) RetrieveIsHashApproved(owner common.Address, safeHash []byte) (bool, error) {
	res, err := safe_.SafeContract.ApprovedHashes(nil, owner, [32]byte(safeHash))
	if res.Uint64() == 0 {
		return false, nil
	} else if res.Uint64() == 1 {
		return true, nil
	}
	return false, err
}

func (safe_ *Safe) RetrieveIsMessageSigned(safeHash []byte) (bool, error) {
	res, err := safe_.SafeContract.SignedMessages(nil, [32]byte(safeHash))
	if res.Uint64() == 0 {
		return false, nil
	} else if res.Uint64() == 1 {
		return true, nil
	}
	return false, err
}

func (safe_ *Safe) RetrieveIsOwner(owner common.Address) (bool, error) {
	return safe_.SafeContract.IsOwner(nil, owner)
}

func (safe_ *Safe) RetrieveOwners() ([]common.Address, error) {
	return safe_.SafeContract.GetOwners(nil)
}
