package safecreations

import (
	"errors"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	safe_utils "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/utils"
	"github.com/ethereum/go-ethereum/common"
)

type SafeCreationTx struct {
	ethereumClient       *eth.EthereumClient // Web3 instance
	owners               []common.Address    // Owners of the Safe
	threshold            int64               // Minimum number of users required to operate the Safe
	masterCopy           common.Address      // Safe master copy address
	fallbackHandler      common.Address      // Handler for fallback calls to the Safe
	funder               common.Address      // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken         common.Address      // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64             // Value of payment token per 1 Ether
	fixedCreationCost    int                 // Fixed creation cost of Safe (Wei)
	ExpectedSafeAddress  common.Address      // safe address calculated from sender's current nonce (CREATE_OP)
	Payment              uint64              // refund to be payed after creation to 'funder' address
	CreationGas          uint64
}

func NewSafeCreationTx(
	ethereumClient *eth.EthereumClient,
	owners []common.Address,
	threshold int64,
	masterCopy common.Address,
	fallbackHandler common.Address,
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
		masterCopy:           masterCopy,
		fallbackHandler:      fallbackHandler,
		funder:               funder,
		paymentToken:         paymentToken,
		paymentTokenEthValue: paymentTokenEthValue,
		fixedCreationCost:    fixedCreationCost,
		ExpectedSafeAddress:  eth.NULL_ADDRESS,
		Payment:              0,
		CreationGas:          0,
	}, nil
}

func (safeCreationTx *SafeCreationTx) PredictSafeAddress_CREATE() {
	chainId, err := safeCreationTx.ethereumClient.GetChainId()
	if err != nil {
		return
	}
	proxyFactoryAddress := network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address
	nonce, err := safeCreationTx.ethereumClient.GetNonceForAccount(proxyFactoryAddress, "pending")
	if err != nil {
		return
	}
	safeCreationTx.ExpectedSafeAddress = eth.MakeContractAddress(proxyFactoryAddress, nonce)
}

func (safeCreationTx *SafeCreationTx) EstimateSafeCreation() error {
	// Prepare Safe creation
	randSender, err := eth.RandomAddress()
	if err != nil {
		return err
	}

	gasPrice, err := safeCreationTx.ethereumClient.GasPrice()
	if err != nil {
		return err
	}
	// This initializer will be passed to the proxy and will be called right after proxy is deployed
	safeSetupData := safeCreationTx.GetInitializer()
	if safeSetupData == nil {
		return errors.New("error getting the safe setup initializer")
	}

	// Calculate gas based on experience of previous deployments of the safe
	calculated_gas := safeCreationTx.calculateCreationGas(safeSetupData)
	// Estimate gas using web3
	estimated_gas := safeCreationTx.estimateCreationGas(*randSender, safeSetupData)

	gas := max(calculated_gas, estimated_gas)
	if safeCreationTx.funder != eth.NULL_ADDRESS {
		payment := calculatePayment(safeCreationTx.fixedCreationCost, safeCreationTx.paymentTokenEthValue, int64(gas), gasPrice.Int64())
		safeCreationTx.Payment = payment
	}
	safeCreationTx.CreationGas = gas
	return nil
}

func (safeCreationTx *SafeCreationTx) GetInitializer() []byte {
	initializer, err := getInitialSetupSafeData(
		safeCreationTx.owners,
		safeCreationTx.threshold,
		safeCreationTx.Payment,
		safeCreationTx.fallbackHandler,
		safeCreationTx.paymentToken,
		safeCreationTx.funder,
	)
	if err != nil {
		return nil
	}
	return initializer
}

func (safeCreationTx *SafeCreationTx) calculateCreationGas(safeSetupData []byte) uint64 {
	/* Calculate gas manually, based on tests of previosly deployed safes */
	baseGas := 60580 // Transaction standard gas

	// If we already have the token, we don't have to pay for storage, so it will be just 5K instead of 20K.
	// The other 1K is for overhead of making the call
	paymentTokenGas := 0
	if safeCreationTx.paymentToken != eth.NULL_ADDRESS {
		paymentTokenGas = 55000
	}

	dataGas := eth.GAS_CALL_DATA_BYTE * len(safeSetupData) // Data gas
	gasPerOwner := 18020                                   // Magic number calculated by testing and averaging owners
	return uint64(baseGas + dataGas + paymentTokenGas + 270000 + (len(safeCreationTx.owners) * gasPerOwner))
}

func (safeCreationTx *SafeCreationTx) estimateCreationGas(sender common.Address, safeSetupData []byte) uint64 {
	_, _, estimatedEIP1559Gas, gasPrice, err := safe_utils.GetDefaultTxParams(
		safeCreationTx.ethereumClient, sender)
	if err != nil {
		return 0
	}
	estimatedGas, err := safeCreationTx.ethereumClient.EstimateGas(
		sender,
		&eth.NULL_ADDRESS,
		0,
		gasPrice,
		estimatedEIP1559Gas.GasFeeCap, estimatedEIP1559Gas.GasTipCap, common.Big0, safeSetupData, nil, nil, nil)
	if err != nil {
		return 0
	}

	// We estimate the refund as a new tx
	if safeCreationTx.paymentToken == eth.NULL_ADDRESS {
		// Same cost to send 1 ether than 1000
		estimatedEthTransferGas, err := safeCreationTx.ethereumClient.EstimateGas(
			sender,
			&safeCreationTx.funder,
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
			return uint64(estimatedGas)
		}
		method := erc20Abi.Methods["transfer"].Name
		transfer, err := erc20Abi.Pack(method, safeCreationTx.funder, 1)
		if err != nil {
			return uint64(estimatedGas)
		}
		estimatedERC20TransferGas, err := safeCreationTx.ethereumClient.EstimateGas(
			sender,
			&safeCreationTx.paymentToken,
			0,
			gasPrice,
			estimatedEIP1559Gas.GasFeeCap, estimatedEIP1559Gas.GasTipCap, common.Big0, transfer, nil, nil, nil)
		if err != nil {
			return uint64(estimatedGas)
		}
		estimatedGas += estimatedERC20TransferGas
	}
	return uint64(estimatedGas)
}
