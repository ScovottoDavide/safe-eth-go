package safecreations

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth_crypto "github.com/ethereum/go-ethereum/crypto"
)

type SafeCreationTx2 struct {
	ethereumClient       *eth.EthereumClient // Web3 instance
	owners               []common.Address    // Owners of the Safe
	threshold            int64               // Minimum number of users required to operate the Safe
	masterCopy           common.Address      // Safe master copy address
	fallbackHandler      common.Address      // Handler for fallback calls to the Safe
	funder               common.Address      // Address to refund when the Safe is created. Address(0) if no need to refund
	paymentToken         common.Address      // Payment token instead of paying the funder with ether. If None Ether will be used
	paymentTokenEthValue float64             // Value of payment token per 1 Ether
	fixedCreationCost    int                 // Fixed creation cost of Safe (Wei)
	ExpectedSafeAddress2 common.Address      // safe address calculated from sender's current nonce (CREATE2_OP)
	saltNonce            int64               // salt nonce used for create2 safe address pre-deploy calculation
	Payment              uint64              // refund to be payed after creation to 'funder' address
	CreationGas          uint64
}

func NewSafeCreationTx2(
	ethereumClient *eth.EthereumClient,
	owners []common.Address,
	threshold int64,
	masterCopy common.Address,
	fallbackHandler common.Address,
	funder common.Address,
	paymentToken common.Address,
	paymentTokenEthValue float64,
	fixedCreationCost int,
	saltNonce int64,
) (*SafeCreationTx2, error) {
	if threshold <= 0 || int(threshold) > len(owners) {
		return nil, errors.New("threshold cannot be negative or greter the number of Safe owners")
	}

	return &SafeCreationTx2{
		ethereumClient:       ethereumClient,
		owners:               owners,
		threshold:            threshold,
		masterCopy:           masterCopy,
		fallbackHandler:      fallbackHandler,
		funder:               funder,
		paymentToken:         paymentToken,
		paymentTokenEthValue: paymentTokenEthValue,
		fixedCreationCost:    fixedCreationCost,
		ExpectedSafeAddress2: eth.NULL_ADDRESS,
		saltNonce:            saltNonce,
		Payment:              0,
		CreationGas:          0,
	}, nil
}

func (safeCreationTx2 *SafeCreationTx2) EstimateSafeCreation2() error {
	gasPrice, err := safeCreationTx2.ethereumClient.GasPrice()
	if err != nil {
		return err
	}
	// This initializer will be passed to the proxy and will be called right after proxy is deployed
	safeSetupData := safeCreationTx2.GetInitializer2()
	if safeSetupData == nil {
		return errors.New("error getting the safe setup initializer")
	}

	// Calculate gas based on experience of previous deployments of the safe
	calculated_gas := safeCreationTx2.calculateCreationGas(safeSetupData)
	// Estimate gas using web3
	estimated_gas := safeCreationTx2.estimateCreationGas2(safeSetupData)

	gas := max(calculated_gas, estimated_gas)
	payment := calculatePayment(safeCreationTx2.fixedCreationCost, safeCreationTx2.paymentTokenEthValue, int64(gas), gasPrice.Int64())

	safeCreationTx2.CreationGas = gas
	safeCreationTx2.Payment = payment
	return nil
}

func (safeCreationTx2 *SafeCreationTx2) PredictSafeAddress_CREATE2() {
	// first estimate gas creation in order to fill up the Payment param
	err := safeCreationTx2.EstimateSafeCreation2()
	if err != nil {
		return
	}

	chainId, err := safeCreationTx2.ethereumClient.GetChainId()
	if err != nil {
		return
	}
	proxyFactoryAddress := network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address
	proxyFactory, err := contracts.NewGnosisSafeProxyFactory(
		proxyFactoryAddress,
		safeCreationTx2.ethereumClient.GetGEthClient(),
	)
	if err != nil {
		return
	}

	uint256, _ := abi.NewType("uint256", "uint256", nil)
	bytesTy, _ := abi.NewType("bytes", "bytes", nil)

	arguments := abi.Arguments{
		{
			Type: bytesTy,
		},
		{
			Type: uint256,
		},
	}

	salt, err := arguments.Pack(eth_crypto.Keccak256(safeCreationTx2.GetInitializer2()), big.NewInt(safeCreationTx2.saltNonce))
	//salt, err := rlp.EncodeToBytes([]interface{}{eth_crypto.Keccak256(safeCreationTx2.GetInitializer2()), uint64(safeCreationTx2.saltNonce)})
	fmt.Println("Salt: ", hexutil.Encode(salt))

	if err != nil {
		return
	}
	salt32 := eth_crypto.Keccak256Hash(salt)

	creationCode, err := proxyFactory.ProxyCreationCode(nil)
	if err != nil {
		return
	}

	arguments = abi.Arguments{
		{
			Type: bytesTy,
		},
		{
			Type: uint256,
		},
	}
	masterCopyInt := new(big.Int)
	masterCopyInt.SetString(safeCreationTx2.masterCopy.String(), 16)
	deploymentData, err := arguments.Pack(creationCode, masterCopyInt)
	inithash := eth_crypto.Keccak256Hash(deploymentData)
	// inithash, err := rlp.EncodeToBytes([]interface{}{creationCode, masterCopyInt})
	if err != nil {
		return
	}

	safeCreationTx2.ExpectedSafeAddress2 = common.BytesToAddress(
		eth_crypto.Keccak256(
			[]byte{0xff},
			proxyFactoryAddress.Bytes(),
			salt32[:],
			inithash.Bytes())[12:])
}

func (safeCreationTx2 *SafeCreationTx2) GetInitializer2() []byte {
	initializer, err := getInitialSetupSafeData(
		safeCreationTx2.owners,
		safeCreationTx2.threshold,
		safeCreationTx2.Payment,
		safeCreationTx2.fallbackHandler,
		safeCreationTx2.paymentToken,
		safeCreationTx2.funder,
	)
	if err != nil {
		return nil
	}
	return initializer
}

func (safeCreationTx2 *SafeCreationTx2) calculateCreationGas(safeSetupData []byte) uint64 {
	baseGas := 205000 // Transaction base gas

	// If we already have the token, we don't have to pay for storage, so it will be just 5K instead of 20K.
	// The other 1K is for overhead of making the call
	paymentTokenGas := 0
	if safeCreationTx2.paymentToken != eth.NULL_ADDRESS {
		paymentTokenGas = 55000
	}
	dataGas := eth.GAS_CALL_DATA_BYTE * len(safeSetupData) //# Data gas
	gasPerOwner := 20000                                   // Magic number calculated by testing and averaging owners
	return uint64(baseGas + dataGas + paymentTokenGas + len(safeCreationTx2.owners)*gasPerOwner)
}

func (safeCreationTx2 *SafeCreationTx2) estimateCreationGas2(safeSetupData []byte) uint64 {
	randSender, err := eth.RandomAddress()
	if err != nil {
		return 0
	}
	chainId, err := safeCreationTx2.ethereumClient.GetChainId()
	if err != nil {
		return 0
	}
	proxyFactoryAddress := network.NetworkToSafeProxyFactoryAddress[network.GetNetwork(chainId)].Address
	proxyFactoryAbi, err := contracts.GnosisSafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return 0
	}
	method := proxyFactoryAbi.Methods["createProxyWithNonce"].Name
	input, err := proxyFactoryAbi.Pack(
		method,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		safeSetupData,
		big.NewInt(safeCreationTx2.saltNonce),
	)
	if err != nil {
		return 0
	}
	estimatedGas, err := safeCreationTx2.ethereumClient.EstimateGas(
		*randSender,
		&proxyFactoryAddress,
		0,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		input,
		nil, nil, nil,
	)
	if err != nil {
		return 0
	}

	if safeCreationTx2.paymentToken == eth.NULL_ADDRESS {
		estimatedGasTransfer, err := safeCreationTx2.ethereumClient.EstimateGas(
			eth.NULL_ADDRESS,
			&safeCreationTx2.funder,
			0,
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			nil,
			nil, nil, nil,
		)
		if err != nil {
			return 0
		}
		estimatedGas += estimatedGasTransfer
	} else {
		estimatedGas += 55000
	}
	return estimatedGas
}
