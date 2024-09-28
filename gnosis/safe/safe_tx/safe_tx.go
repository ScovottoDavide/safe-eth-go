package safe_tx

import (
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	safesignature "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_signature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type SafeTx struct {
	EthereumClient *eth.EthereumClient
	SafeAddress    common.Address
	SafeContract   *contracts.GnosisSafe
	To             common.Address
	Value          big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      big.Int
	BaseGas        big.Int
	GasPrice       big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signatures     []byte
	SafeNonce      *big.Int
	SafeVersion    string
	ChainId        int
	Tx             *types.Transaction // set if `SafeTx` is executed
	TxHash         *common.Hash       // set if `SafeTx` is executed
}

func New(
	ethereumClient *eth.EthereumClient,
	safeAddress common.Address,
	to common.Address,
	value big.Int,
	data []byte,
	operation uint8,
	safeTxGas big.Int,
	baseGas big.Int,
	gasPrice big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	signatures []byte,
	safeNonce *big.Int,
	safeVersion *string,
) *SafeTx {
	chainId, err := ethereumClient.GetChainId()
	if err != nil {
		return nil
	}
	safeContract, err := contracts.NewGnosisSafe(safeAddress, ethereumClient.GetGEthClient())
	if err != nil {
		return nil
	}
	if safeNonce == nil {
		safeNonce, err = safeContract.Nonce(nil)
		if err != nil {
			return nil
		}
	}
	if safeVersion == nil {
		safeVersion_, err := safeContract.VERSION(nil)
		safeVersion = &safeVersion_
		if err != nil {
			return nil
		}
	}
	return &SafeTx{
		EthereumClient: ethereumClient,
		SafeAddress:    safeAddress,
		SafeContract:   safeContract,
		To:             to,
		Value:          value,
		Data:           data,
		Operation:      operation,
		SafeTxGas:      safeTxGas,
		BaseGas:        baseGas,
		GasPrice:       gasPrice,
		GasToken:       gasToken,
		RefundReceiver: refundReceiver,
		Signatures:     signatures,
		SafeNonce:      safeNonce,
		SafeVersion:    *safeVersion,
		ChainId:        chainId,
		Tx:             nil,
		TxHash:         nil,
	}
}

// only for Version >= 1.3.0
func (safeTx *SafeTx) EIP712StructuredData() apitypes.TypedData {
	types := apitypes.Types{
		"EIP712Domain": {
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"SafeTx": {
			{Name: "to", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "data", Type: "bytes"},
			{Name: "operation", Type: "uint8"},
			{Name: "safeTxGas", Type: "uint256"},
			{Name: "baseGas", Type: "uint256"},
			{Name: "gasPrice", Type: "uint256"},
			{Name: "gasToken", Type: "address"},
			{Name: "refundReceiver", Type: "address"},
			{Name: "nonce", Type: "uint256"},
		},
	}
	dataDomain := apitypes.TypedDataDomain{
		ChainId:           math.NewHexOrDecimal256(int64(safeTx.ChainId)),
		VerifyingContract: safeTx.SafeAddress.Hex(),
	}
	message := apitypes.TypedDataMessage{
		"to":             safeTx.To,
		"value":          safeTx.Value,
		"data":           safeTx.Data,
		"operation":      safeTx.Operation,
		"safeTxGas":      safeTx.SafeTxGas,
		"baseGas":        safeTx.BaseGas,
		"gasPrice":       safeTx.GasPrice,
		"gasToken":       safeTx.GasToken,
		"refundReceiver": safeTx.RefundReceiver,
		"nonce":          safeTx.SafeNonce,
	}
	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "SafeTx",
		Domain:      dataDomain,
		Message:     message,
	}
	return typedData
}

func (safeTx *SafeTx) SafeTxHash() (hexutil.Bytes, error) {
	typedData := safeTx.EIP712StructuredData()
	messageHash, err := typedData.HashStruct("SafeTx", typedData.Message)
	if err != nil {
		return nil, err
	}
	return messageHash, nil
}

func (safeTx *SafeTx) Signers() []common.Address {
	if len(safeTx.Signatures) == 0 {
		return nil
	}
	// TODO: parse signatures, get the owner for each one and return the array
	safe_tx_hash, err := safeTx.SafeTxHash()
	if err != nil {
		return nil
	}
	signatures := safesignature.ParseSignatures[safesignature.SafeSignatureContract](
		safeTx.Signatures,
		safe_tx_hash,
	)
	return nil
}
