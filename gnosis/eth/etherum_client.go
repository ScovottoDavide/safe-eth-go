package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

func (ethereumClient *EthereumClient) GasPrice() (*big.Int, error) {
	return ethereumClient.ethereumClient.SuggestGasPrice(context.Background())
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
		Reward:  feeHistory.Reward[0][0],
		BaseFee: feeHistory.BaseFee[len(feeHistory.BaseFee)-1],
	}
	return estimatedGas, nil
}

func (ethereumClient *EthereumClient) SetEip1559Fee(msg *ethereum.CallMsg, txSpeed TxSpeed) {
	eip1559_estimated_gas, err := ethereumClient.EstimateFeeEip1559(txSpeed)
	if err != nil {
		panic(err)
	}
	msg.GasFeeCap = eip1559_estimated_gas.BaseFee
	msg.GasTipCap = eip1559_estimated_gas.Reward
}

func (ethereumClient *EthereumClient) GetBalance(iaddress interface{}) (*big.Int, error) {
	address := new(common.Address)
	err := error(nil)
	switch v := iaddress.(type) {
	case string:
		address, err = StringToAddress(v)
		if err != nil {
			return nil, err
		}
	case *common.Address:
		address = v
	case common.Address:
		*address = v
	}
	return ethereumClient.ethereumClient.BalanceAt(context.Background(), *address, nil)
}

func (ethereumClient *EthereumClient) GetTransaction(txHash string) (*types.Transaction, bool, error) {
	_, err := hexutil.Decode(txHash)
	if err != nil {
		return nil, false, err
	}
	ethTxHash := common.HexToHash(txHash)
	tx, isPending, err := ethereumClient.ethereumClient.TransactionByHash(context.Background(), ethTxHash)
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}

func (ethereumClient *EthereumClient) GetTransactions(txHashes []string) ([]*batcthedTransactionResult, error) {
	if len(txHashes) == 0 {
		return nil, errors.New("txHashes list is empty")
	}
	transactions := make([]*rpcTransaction, len(txHashes))
	reqs := make([]rpc.BatchElem, len(txHashes))
	for i, txHash := range txHashes {
		_, err := hexutil.Decode(txHash)
		if err != nil {
			return nil, err
		}
		ethTxHash := common.HexToHash(txHash)
		fmt.Println("hash: ", ethTxHash)
		reqs[i] = rpc.BatchElem{
			Method: "eth_getTransactionByHash",
			Args:   []interface{}{ethTxHash},
			Result: &transactions[i].Tx,
		}
	}
	if err := ethereumClient.ethereumClient.Client().BatchCallContext(context.Background(), reqs); err != nil {
		return nil, err
	}
	for i := range reqs {
		if reqs[i].Error != nil {
			return nil, reqs[i].Error
		}
		if transactions[i].Tx == nil {
			return nil, fmt.Errorf("%d, got null transaction for transaction with hash %s", i, common.HexToHash(txHashes[i]))
		}
	}

	result := make([]*batcthedTransactionResult, len(transactions))
	for _, tx := range transactions {
		tx_result := batcthedTransactionResult{
			Tx:        tx.Tx,
			IsPending: tx.BlockNumber == nil,
		}
		result = append(result, &tx_result)
		fmt.Println("chain id ", tx.Tx.ChainId())
		fmt.Println("cost (value + fees)", tx.Tx.Cost())
		fmt.Println("data ", tx.Tx.Data())
		fmt.Println("gas ", tx.Tx.Gas())
		fmt.Println("gas price ", tx.Tx.GasPrice())
		fmt.Println("hash ", tx.Tx.Hash())
		fmt.Println("value ", tx.Tx.Value())
	}
	return result, nil
}

func (ethereumClient *EthereumClient) SendUnsignedTransaction(
	iprivateKey interface{},
	tx *types.Transaction,
) (common.Hash, error) {
	privateKey, err := GetCryptoPrivateKey(iprivateKey)
	if err != nil {
		return common.Hash{}, err
	}

	R, S, V := tx.RawSignatureValues()
	if (V.Int64() != 0) || (R.Int64() != 0) || (S.Int64() != 0) {
		return common.Hash{}, fmt.Errorf("transaction seems to have at least one signature value (v, r, s) filled")
	}
	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return common.Hash{}, err
	}

	signedTx, err := types.SignTx(tx, types.NewCancunSigner(big.NewInt(int64(chainId))), privateKey)
	if err != nil {
		return common.Hash{}, err
	}
	err = ethereumClient.ethereumClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, err
	}
	return signedTx.Hash(), nil
}

func (ethereumClient *EthereumClient) BuildTransaction(
	From common.Address, // the sender of the 'transaction'
	To *common.Address, // the destination contract (nil for contract creation)
	Gas uint64, // if 0, the call executes with near-infinite gas
	GasPrice *big.Int, // wei <-> gas exchange ratio
	GasFeeCap *big.Int, // EIP-1559 fee cap per gas.
	GasTipCap *big.Int, // EIP-1559 tip per gas.
	Value *big.Int, // amount of wei sent along with the call
	Data []byte, // input data, usually an ABI-encoded contract method invocation
	AccessList types.AccessList, // EIP-2930 access list.
) (*types.Transaction, error) {
	txType := types.DynamicFeeTxType
	if GasFeeCap == nil && GasTipCap == nil {
		txType = types.LegacyTxType
	} else if (GasFeeCap == nil && GasTipCap != nil) || (GasFeeCap != nil && GasTipCap == nil) {
		return nil, fmt.Errorf("one of GasFeeCap and GasTipCap is nil")
	}

	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return nil, err
	}
	nonce, err := ethereumClient.GetNonceForAccount(From, "pending")
	if err != nil {
		return nil, err
	}
	switch txType {
	case types.DynamicFeeTxType:
		transaction := types.NewTx(&types.DynamicFeeTx{
			ChainID:    IntToBigInt(chainId),
			Nonce:      nonce,
			GasTipCap:  GasTipCap,
			GasFeeCap:  GasFeeCap,
			Gas:        Gas,
			To:         To,
			Value:      Value,
			Data:       Data,
			AccessList: AccessList,
		})
		return transaction, nil
	case types.LegacyTxType:
		transaction := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       To,
			Value:    Value,
			GasPrice: GasPrice,
			Gas:      Gas,
			Data:     Data,
		})
		return transaction, nil
	}
	return nil, fmt.Errorf("detected invalid transaction type. transaction build failed")
}

// func (ethereumClient *EthereumClient) SendSignedTransaction(
// 	tx *types.TxData,
// 	retry bool,
// 	blockIdentifier string,
// ) {

// }

func (ethereumClient *EthereumClient) SendEthTo(
	privateKey string,
	to *common.Address,
	gasPrice *big.Int,
	value *big.Int,
	gas uint64,
) (common.Hash, error) {
	address, err := AddressFromPrivKey(hexutil.MustDecode(privateKey))
	if err != nil {
		return common.Hash{}, err
	}
	nonce, err := ethereumClient.GetNonceForAccount(*address, "latest")
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       to,
		Value:    value,
		Data:     nil,
		// V:        nil, R: nil, S: nil,
	})

	return ethereumClient.SendUnsignedTransaction(privateKey, tx)
}

//func (etheruemClient *EthereumClient) RawBatchRequest() {
//	return
//}

// Returns a channel that blocks until the transaction is confirmed
func (ethereumClient *EthereumClient) WaitTxConfirmed(hash common.Hash) <-chan bool {
	ch := make(chan bool)
	go func() {
		for {
			_, pending, _ := ethereumClient.GetTransaction(hash.Hex())
			if !pending {
				ch <- pending
			}

			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}

func (ethereumClient *EthereumClient) GetTransactionReceipt(txHash common.Hash) *types.Receipt {
	client := ethereumClient.ethereumClient
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil
	}
	return receipt
}

func (ethereumClient *EthereumClient) DeployAndInitializeContract(
	iprivateKey interface{},
	constructorAndInitializerData multipleTxData,
	checkReceipt bool,
	txSpeed TxSpeed,
	value big.Int,
) (*common.Hash, *types.Transaction, *common.Address, error) {
	privateKey, err := GetCryptoPrivateKey(iprivateKey)
	if err != nil {
		return nil, nil, nil, err
	}
	if constructorAndInitializerData.constructorData == nil {
		return nil, nil, nil, fmt.Errorf("missing constructorData parameter")
	}

	deployerAddress, err := AddressFromPrivKey(privateKey.D.Bytes())
	if err != nil {
		return nil, nil, nil, err
	}
	nonce, err := ethereumClient.GetNonceForAccount(*deployerAddress, "pending")
	if err != nil {
		return nil, nil, nil, err
	}
	gasPrice, err := ethereumClient.GasPrice()
	if err != nil {
		return nil, nil, nil, err
	}
	estimatedEIP1559Gas := &EIP1559EstimatedGas{Reward: nil, BaseFee: nil}
	if ethereumClient.IsEip1559Supported() {
		estimatedEIP1559Gas, err = ethereumClient.EstimateFeeEip1559(txSpeed)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	/*
	* params to return
	 */
	newContractTx := new(types.Transaction)
	newContractTxHash := new(common.Hash)
	contractAddress := new(common.Address)
	/**/

	for index, data := range constructorAndInitializerData.toArray() {
		estimatedGas, err := ethereumClient.EstimateGas(
			*deployerAddress,
			nil,
			0,
			gasPrice,
			estimatedEIP1559Gas.BaseFee,
			estimatedEIP1559Gas.Reward,
			&value,
			data,
			nil,
			nil,
			nil,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		newContractTx, err = ethereumClient.BuildTransaction(
			*deployerAddress,
			nil,
			estimatedGas,
			gasPrice,
			estimatedEIP1559Gas.BaseFee,
			estimatedEIP1559Gas.Reward,
			&value,
			data,
			nil,
		)
		if err != nil {
			return nil, nil, nil, err
		}

		txHash, err := ethereumClient.SendUnsignedTransaction(
			privateKey, newContractTx,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		if index == 0 { // deploying contract
			*contractAddress = MakeContractAddress(*deployerAddress, nonce)
			newContractTxHash = &txHash
		}

		if checkReceipt {
			isPending := <-ethereumClient.WaitTxConfirmed(txHash)
			if !isPending {
				return nil, nil, nil, fmt.Errorf("checkReceipt::Transaction should not be pending")
			}
			receipt := ethereumClient.GetTransactionReceipt(txHash)
			if receipt == nil {
				return nil, nil, nil, fmt.Errorf("checkReceipt::Got nil receipt")
			}
			if !isTransactionSuccessful(receipt) {
				return nil, nil, nil, fmt.Errorf("checkReceipt::Got Unsucessful Receipt Status")
			}
		}
	}
	return newContractTxHash, newContractTx, contractAddress, nil
}
