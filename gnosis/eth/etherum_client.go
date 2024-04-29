package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// ethRpc "github.com/ethereum/go-ethereum/rpc"
)

type EthereumClient struct {
	ethereumNodeUrl URI
	ethereumClient  *ethclient.Client
}

func EthereumClientInit(uri *URI) (*EthereumClient, error) {
	client, err := ethclient.Dial(uri.address)
	if err != nil {
		return nil, err
	}

	chain_id, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	if !isSupported(int(chain_id.Int64())) {
		return nil, errors.New("chain not supported")
	}
	fmt.Println("Successfully connected to network with {address: ", uri.address, ", chainId: ", chain_id, "}")

	ethereum_client := &EthereumClient{
		ethereumNodeUrl: *uri,
		ethereumClient:  client,
	}
	return ethereum_client, nil
}

func (ethereumClient *EthereumClient) GetGEthClient() *ethclient.Client {
	return ethereumClient.ethereumClient
}

func (ethereumClient *EthereumClient) GetUri() URI {
	return ethereumClient.ethereumNodeUrl
}

func (ethereumClient *EthereumClient) CurrentBlockNumber() (uint64, error) {
	blockNum, err := ethereumClient.ethereumClient.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

func (ethereumClient *EthereumClient) GetChainId() (int, error) {
	chain_id, err := ethereumClient.ethereumClient.ChainID(context.Background())
	if err != nil {
		return Unknown.chainId, err
	}
	return int(chain_id.Int64()), nil
}

func (ethereumClient *EthereumClient) GetClientVersion() (string, error) {
	var response string
	err := ethereumClient.ethereumClient.Client().Call(&response, "web3_clientVersion")
	if err != nil {
		return "", err
	}
	return response, nil
}

func (ethereumClient *EthereumClient) IsEip1559Supported() bool {
	// hardhat node causes an os.exit() !!!
	_, err := ethereumClient.ethereumClient.FeeHistory(context.Background(), 1, nil, make([]float64, 50))
	return err == nil
}

func (ethereumClient *EthereumClient) GetNonceForAccount(account common.Address, blockIdentifier string) (uint64, error) {
	var result hexutil.Uint64
	err := ethereumClient.ethereumClient.Client().CallContext(
		context.Background(), &result, "eth_getTransactionCount", account, blockIdentifier,
	)
	return uint64(result), err
}

func AddressFromPrivKey(s_key_hex []byte) (*common.Address, error) {
	privateKey, err := crypto.ToECDSA(s_key_hex)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &address, nil
}

func (ethereumClient *EthereumClient) EstimateGas(
	From common.Address, // the sender of the 'transaction'
	To *common.Address, // the destination contract (nil for contract creation)
	Gas uint64, // if 0, the call executes with near-infinite gas
	GasPrice *big.Int, // wei <-> gas exchange ratio
	GasFeeCap *big.Int, // EIP-1559 fee cap per gas.
	GasTipCap *big.Int, // EIP-1559 tip per gas.
	Value *big.Int, // amount of wei sent along with the call
	Data []byte, // input data, usually an ABI-encoded contract method invocation
	AccessList types.AccessList, // EIP-2930 access list.
	// For BlobTxType
	BlobGasFeeCap *big.Int,
	BlobHashes []common.Hash,
) (uint64, error) {
	msg := ethereum.CallMsg{
		From:          From,
		To:            To,
		Gas:           Gas,
		GasPrice:      GasPrice,
		GasFeeCap:     GasFeeCap,
		GasTipCap:     GasTipCap,
		Value:         Value,
		Data:          Data,
		AccessList:    AccessList,
		BlobGasFeeCap: BlobGasFeeCap,
		BlobHashes:    BlobHashes,
	}
	estimatedGas, err := ethereumClient.ethereumClient.EstimateGas(context.Background(), msg)
	return estimatedGas, err
}

func EstimateDataGas(data []byte) uint64 {
	var gas uint64

	for _, d := range data {
		if d == byte(0) {
			gas += GAS_CALL_DATA_ZERO_BYTE
		} else {
			gas += GAS_CALL_DATA_BYTE
		}
	}

	return gas
}

func (ethereumClient *EthereumClient) EstimateFeeEip1559(txSpeed TxSpeed) (*EIP1559EstimatedGas, error) {
	var localSpeed TxSpeed

	if isTxSpeedAvailable(txSpeed) {
		localSpeed = tx_speed_map[txSpeed.speed]
	} else {
		return nil, errors.New("provided TxSpeed is unavailable")
	}

	feeHistory, err := ethereumClient.ethereumClient.FeeHistory(context.Background(), 1, nil, make([]float64, localSpeed.speed))
	if err != nil {
		return nil, err
	}

	estimatedGas := &EIP1559EstimatedGas{
		Reward:  *feeHistory.Reward[0][0],
		BaseFee: *feeHistory.BaseFee[len(feeHistory.BaseFee)-1],
	}
	return estimatedGas, nil
}

//func (etheruemClient *EthereumClient) RawBatchRequest() {
//	return
//}
