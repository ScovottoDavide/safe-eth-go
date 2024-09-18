package safecreations

import (
	"errors"
	"math"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	safe_utils "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/utils"
	"github.com/ethereum/go-ethereum/common"
)

type SafeCreationTx struct {
	ethereumClient       *eth.EthereumClient // Web3 instance
	owners               []common.Address    // Owners of the Safe
	threshold            int64               // Minimum number of users required to operate the Safe
	signature_s          uint64              // Random s value for ecdsa signature
	masterCopy           common.Address      // Safe master copy address
	funder               common.Address      // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken         common.Address      // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64             // Value of payment token per 1 Ether
	fixedCreationCost    int                 // Fixed creation cost of Safe (Wei)
}

func NewSafeCreationTx(
	ethereumClient *eth.EthereumClient,
	owners []common.Address,
	threshold int64,
	signature_s uint64,
	masterCopy common.Address,
	funder common.Address,
	paymentToken common.Address,
	paymentTokenEthValue float64,
	fixedCreationCost int,
) (*SafeCreationTx, error) {
	if threshold <= 0 || int(threshold) > len(owners) {
		return nil, errors.New("threshold cannot be negative or greter the number of Safe owners")
	}

	return &SafeCreationTx{
		ethereumClient:       ethereumClient,
		owners:               owners,
		threshold:            threshold,
		signature_s:          signature_s,
		masterCopy:           masterCopy,
		funder:               funder,
		paymentToken:         paymentToken,
		paymentTokenEthValue: paymentTokenEthValue,
		fixedCreationCost:    fixedCreationCost,
	}, nil
}

func (safeCreationTx *SafeCreationTx) EstimateSafeCreation(sender common.Address, gasPrice int64) (int64, int64, uint64, error) {
	// Prepare Safe creation

	// This initializer will be passed to the proxy and will be called right after proxy is deployed
	safeSetupData, err := getInitialSetupSafeData(
		safeCreationTx.owners, safeCreationTx.threshold, safeCreationTx.paymentToken, safeCreationTx.funder,
	)
	if err != nil {
		return 0, 0, 0, err
	}

	// Calculate gas based on experience of previous deployments of the safe
	calculated_gas := calculateCreationGas(safeCreationTx.owners, safeSetupData, safeCreationTx.paymentToken)
	// Estimate gas using web3
	estimated_gas := estimateCreationGas(
		safeCreationTx.ethereumClient, sender, safeCreationTx.funder, safeSetupData, safeCreationTx.paymentToken)

	gas := max(calculated_gas, estimated_gas)
	payment := calculateRefundPayment(gas, gasPrice, safeCreationTx.fixedCreationCost, safeCreationTx.paymentTokenEthValue)

	return gas, gasPrice, payment, nil
}

func getInitialSetupSafeData(
	owners []common.Address,
	threshold int64,
	paymentToken common.Address,
	paymentReceiver common.Address,
) ([]byte, error) {
	/* construct the initialization data needed for the proxy to initialize the Safe */
	safeAbi, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	method := safeAbi.Methods["setup"].Name
	initializer, err := safeAbi.Pack(
		method,
		owners,
		big.NewInt(threshold),
		eth.NULL_ADDRESS, // Contract address for optional delegate call
		make([]byte, 1),  // Data payload for optional delegate call
		eth.NULL_ADDRESS,
		paymentToken,
		common.Big0,
		paymentReceiver,
	)
	if err != nil {
		return nil, err
	}
	return initializer, nil
}

func calculateCreationGas(
	owners []common.Address,
	safeSetupData []byte,
	paymentToken common.Address,
) int64 {
	/* Calculate gas manually, based on tests of previosly deployed safes */
	baseGas := 60580 // Transaction standard gas

	// If we already have the token, we don't have to pay for storage, so it will be just 5K instead of 20K.
	// The other 1K is for overhead of making the call
	paymentTokenGas := 0
	if paymentToken != eth.NULL_ADDRESS {
		paymentTokenGas = 55000
	}

	dataGas := eth.GAS_CALL_DATA_BYTE * len(safeSetupData) // Data gas
	gasPerOwner := 18020                                   // Magic number calculated by testing and averaging owners
	return int64(baseGas + dataGas + paymentTokenGas + 270000 + (len(owners) * gasPerOwner))
}

func estimateCreationGas(
	ethereumClient *eth.EthereumClient,
	sender common.Address,
	funder common.Address,
	safeSetupData []byte,
	paymentToken common.Address,
) int64 {
	_, _, estimatedEIP1559Gas, gasPrice, err := safe_utils.GetDefaultTxParams(ethereumClient, sender)
	if err != nil {
		return 0
	}
	estimatedGas, err := ethereumClient.EstimateGas(
		sender,
		&eth.NULL_ADDRESS,
		0,
		gasPrice,
		estimatedEIP1559Gas.GasFeeCap, estimatedEIP1559Gas.GasTipCap, common.Big0, safeSetupData, nil, nil, nil)
	if err != nil {
		return 0
	}

	// We estimate the refund as a new tx
	if paymentToken == eth.NULL_ADDRESS {
		// Same cost to send 1 ether than 1000
		estimatedEthTransferGas, err := ethereumClient.EstimateGas(
			sender,
			&funder,
			1e18,
			gasPrice,
			estimatedEIP1559Gas.GasFeeCap, estimatedEIP1559Gas.GasTipCap, common.Big0, nil, nil, nil, nil)
		if err != nil {
			return 0
		}
		estimatedGas += estimatedEthTransferGas
	} else {
		erc20Abi, err := contracts.ERC20MetaData.GetAbi()
		if err != nil {
			return int64(estimatedGas)
		}
		method := erc20Abi.Methods["transfer"].Name
		transfer, err := erc20Abi.Pack(method, funder, 1)
		if err != nil {
			return int64(estimatedGas)
		}
		estimatedERC20TransferGas, err := ethereumClient.EstimateGas(
			sender,
			&paymentToken,
			0,
			gasPrice,
			estimatedEIP1559Gas.GasFeeCap, estimatedEIP1559Gas.GasTipCap, common.Big0, transfer, nil, nil, nil)
		if err != nil {
			return int64(estimatedGas)
		}
		estimatedGas += estimatedERC20TransferGas
	}
	return int64(estimatedGas)
}

func calculateRefundPayment(
	gas int64,
	gasPrice int64,
	fixedCreationCost int,
	paymentTokenEthValue float64,
) uint64 {
	if fixedCreationCost == 0 {
		// Payment will be safe deploy cost + transfer fees for sending ether to the deployer
		basePayment := (gas + 23000) * gasPrice
		// Calculate payment for tokens using the conversion (if used)
		return uint64(math.Ceil(float64(basePayment) / paymentTokenEthValue))
	} else {
		return uint64(fixedCreationCost)
	}
}
