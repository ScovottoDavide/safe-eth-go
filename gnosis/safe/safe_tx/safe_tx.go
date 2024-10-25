package safetx

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"slices"
	"sort"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	safesignature "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_signature"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type SafeTx struct {
	EthereumClient *eth.EthereumClient
	SafeAddress    common.Address
	SafeContract   *contracts.GnosisSafe
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signatures     []byte
	SafeNonce      *big.Int
	SafeVersion    string
	ChainId        int
	Tx             *types.Transaction // set if `SafeTx` is executed/prepared
	TxHash         *common.Hash       // set if `SafeTx` is executed
	Signers        []common.Address
}

func New(
	ethereumClient *eth.EthereumClient,
	safeAddress common.Address,
	to common.Address,
	value *big.Int,
	data []byte,
	operation uint8,
	safeTxGas *big.Int,
	baseGas *big.Int,
	gasPrice *big.Int,
	gasToken common.Address,
	refundReceiver common.Address,
	safeNonce *big.Int,
	signatures []byte,
	safeVersion string,
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
	if safeVersion == "" {
		safeVersion_, err := safeContract.VERSION(nil)
		safeVersion = safeVersion_
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
		SafeVersion:    safeVersion,
		ChainId:        chainId,
		Tx:             nil,
		TxHash:         nil,
		Signers:        nil,
	}
}

// only for Version >= 1.3.0
func (safeTx *SafeTx) EIP712StructuredData() apitypes.TypedData {
	types := apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"SafeTx": []apitypes.Type{
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
		VerifyingContract: safeTx.SafeAddress.Hex(),
		ChainId:           math.NewHexOrDecimal256(int64(safeTx.ChainId)),
	}
	message := apitypes.TypedDataMessage{
		"to":             safeTx.To.Hex(),
		"value":          safeTx.Value.String(),
		"data":           safeTx.Data,
		"operation":      fmt.Sprintf("%d", safeTx.Operation),
		"safeTxGas":      fmt.Sprintf("%#d", safeTx.SafeTxGas),
		"baseGas":        fmt.Sprintf("%#d", safeTx.BaseGas),
		"gasPrice":       safeTx.GasPrice.String(),
		"gasToken":       safeTx.GasToken.Hex(),
		"refundReceiver": safeTx.RefundReceiver.Hex(),
		"nonce":          fmt.Sprintf("%d", safeTx.SafeNonce.Uint64()),
	}
	typedData := apitypes.TypedData{
		Types:       types,
		PrimaryType: "SafeTx",
		Domain:      dataDomain,
		Message:     message,
	}
	return typedData
}

func (safeTx *SafeTx) SafeTxHash() (common.Hash, error) {
	typedData := safeTx.EIP712StructuredData()
	messageHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return *new(common.Hash), err
	}
	return common.BytesToHash(messageHash), nil
}

// Get signers of the Safe Tx by parsing the signatures
func (safeTx *SafeTx) GetSignersFromSignatures() []common.Address {
	if len(safeTx.Signatures) == 0 {
		return nil
	}
	safe_tx_hash, err := safeTx.SafeTxHash()
	if err != nil {
		return nil
	}
	signatures := safesignature.ParseSignature(
		safeTx.Signatures,
		safe_tx_hash.Bytes(),
	)
	signers := make([]common.Address, len(signatures))
	for i := 0; i < len(signatures); i++ {
		switch signatures[i].(type) {
		case safesignature.SafeSignatureContract:
			signature := signatures[i].(safesignature.SafeSignatureContract)
			signers[i] = *signature.Owner()
		case safesignature.SafeSignatureApprovedHash:
			signature := signatures[i].(safesignature.SafeSignatureApprovedHash)
			signers[i] = *signature.Owner()
		case safesignature.SafeSignatureEOA:
			signature := signatures[i].(safesignature.SafeSignatureEOA)
			signers[i] = *signature.Owner()
		case *safesignature.SafeSignatureEthSign:
			signature := signatures[i].(safesignature.SafeSignatureEthSign)
			signers[i] = *signature.Owner()
		}
	}

	return signers
}

func (safeTx *SafeTx) SortedSigners() []common.Address {
	sort.Slice(safeTx.Signers, func(i, j int) bool {
		return safeTx.Signers[i].Big().Uint64() < safeTx.Signers[j].Big().Uint64()
	})
	return safeTx.Signers
}

func (safeTx *SafeTx) Raw() ([]byte, error) {
	safeABI, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	method := safeABI.Methods["execTransaction"].Name
	execTransactionRaw, err := safeABI.Pack(
		method,
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		safeTx.Signatures,
	)
	if err != nil {
		return nil, err
	}
	return execTransactionRaw, nil
}

// Call executes a `ExecTransaction` transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain. Can be used to simulate the outcome of the
// prepared SafeTx.
func (safeTx *SafeTx) Call(senderAddr common.Address, txGas uint64) error {
	safeTxRaw, err := safeTx.Raw()
	if err != nil {
		return err
	}

	res, err := safeTx.EthereumClient.CallContract(
		senderAddr,
		&safeTx.To,
		txGas,
		safeTx.Value,
		safeTxRaw,
	)
	if err != nil {
		return err
	}
	fmt.Println("execTransaction call result: ", res)
	return nil
}

// Recommended gas to use on the ethereum_tx
func (safeTx *SafeTx) RecommendedGas() uint64 {
	recommendedGas := new(big.Int)
	recommendedGas = recommendedGas.Add(safeTx.BaseGas, safeTx.SafeTxGas)
	return recommendedGas.Uint64() + 75000
}

func (safeTx *SafeTx) Execute(
	privateKey *ecdsa.PrivateKey, // Sender private key
	txGas *uint64, // Gas for the external tx. If not, `(safeTxGas + baseGas) * 2` will be used
	txGasPrice *big.Int, // Gas price of the external tx. If not, `gas_price` will be used
	txNonce *big.Int, // Force nonce for `tx_sender`
	eip1559Speed *network.TxSpeed, // If provided, use EIP1559 transaction
) error {
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(safeTx.ChainId)))
	if err != nil {
		return err
	}
	var gasLimit uint64
	if txGas == nil {
		gasLimit = safeTx.RecommendedGas()
	}

	eip1559EstimateGas := new(eth.EIP1559EstimatedGas)
	if eip1559Speed != nil && safeTx.EthereumClient.IsEip1559Supported() {
		eip1559EstimateGas, err = safeTx.EthereumClient.EstimateFeeEip1559(*eip1559Speed)
		if err != nil {
			return err
		}
		txGasPrice = nil
	} else {
		eip1559EstimateGas = nil
	}

	safeTx.SafeContract.ExecTransaction(
		&bind.TransactOpts{
			Signer:    signer.Signer,
			Nonce:     txNonce,
			GasLimit:  gasLimit,
			GasPrice:  txGasPrice,
			GasFeeCap: eip1559EstimateGas.GasFeeCap,
			GasTipCap: eip1559EstimateGas.GasTipCap,
		},
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		txGasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		safeTx.Signatures,
	)

	// Set signatures empty after executing the tx. `Nonce` is increased even if it fails,
	// so signatures are not valid anymore
	safeTx.Signatures = nil
	return nil
}

// Signs the Safe Transaction and adds (in order) the signature to the SafeTx::Signature byte array and updates
// the SafeTx::Signers common.Address array
func (safeTx *SafeTx) Sign(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	_safe_hash, err := safeTx.SafeTxHash()
	if err != nil {
		return nil, err
	}
	signature, err := safesignature.SignMessageHash(_safe_hash, privateKey)
	if err != nil {
		return nil, err
	}

	address, err := eth.AddressFromPrivKey(privateKey.D.Bytes())
	if err != nil {
		return nil, err
	}
	if !slices.Contains(safeTx.signersToHex(), address.Hex()) {
		safeTx.Signers = append(safeTx.Signers, *address)
		newOwnerPos := slices.Index(safeTx.SortedSigners(), *address)
		safeTx.Signatures = append(safeTx.Signatures[:65*newOwnerPos], append(signature, safeTx.Signatures[65*newOwnerPos:]...)...)
	}
	return signature, nil
}

func (safeTx *SafeTx) Unsign() {

}

func (safeTx *SafeTx) signersToHex() []string {
	var hexSigners []string
	for i := 0; i < len(safeTx.Signers); i++ {
		hexSigners = append(hexSigners, safeTx.Signers[i].Hex())
	}
	return hexSigners
}
