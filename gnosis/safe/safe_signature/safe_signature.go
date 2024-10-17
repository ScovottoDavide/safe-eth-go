package safe_signature

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const SignatureSize = 65

type SafeSignatureType int

const (
	CONTRACT_SIGNATURE SafeSignatureType = 0
	APPROVED_HASH      SafeSignatureType = 1
	EOA                SafeSignatureType = 2
	ETH_SIGN           SafeSignatureType = 3
)

func SigTypeFromV(v uint64) SafeSignatureType {
	if v == 0 {
		return CONTRACT_SIGNATURE
	} else if v == 1 {
		return APPROVED_HASH
	} else if v > 30 {
		return ETH_SIGN
	} else {
		return EOA
	}
}

type safeSignature_ struct {
	safeTxHash []byte
	v, r, s    *big.Int
}

type SafeSignatureContract struct {
	safeSignature     *safeSignature_
	contractSignature []byte
}

type SafeSignatureApprovedHash struct {
	safeSignature *safeSignature_
}

type SafeSignatureEthSign struct {
	safeSignature *safeSignature_
}

type SafeSignatureEOA struct {
	safeSignature *safeSignature_
}

type SafeSignature interface{}

func ParseSignature(
	signatures []byte,
	safeTxHash []byte,
) []SafeSignature {
	if signatures == nil {
		return nil
	}

	dataPosition := len(signatures) // For contract signatures, to stop parsing at data position
	var safeSignatures []SafeSignature
	// var safe_signature T
	for i := 0; i < len(signatures); i += SignatureSize {
		if i >= dataPosition { // If contract signature data position is reached, stop
			break
		}
		v, r, s, err := SplitAndValidateSignature(signatures, i/SignatureSize)
		if err != nil {
			return nil
		}
		signatureType := SigTypeFromV(v.Uint64())
		var safe_signature SafeSignature
		if signatureType == CONTRACT_SIGNATURE {
			s_int := int(s.Uint64())
			if s_int < dataPosition {
				dataPosition = s_int
			}
			contract_signature_len := int(new(big.Int).SetBytes(signatures[s_int : s_int+32]).Uint64()) // Len size is 32 bytes
			contract_signature := signatures[s_int+32 : s_int+32+contract_signature_len]                // Skip array size (32 bytes)
			fmt.Println(contract_signature)
			safe_signature = SafeSignatureContract{
				&safeSignature_{
					safeTxHash: safeTxHash,
					v:          v,
					r:          r,
					s:          s,
				}, contract_signature,
			}
			//safeSignatures = append(safeSignatures, T(safe_signature))
		} else if signatureType == APPROVED_HASH {
			safe_signature = SafeSignatureApprovedHash{
				&safeSignature_{
					safeTxHash: safeTxHash,
					v:          v,
					r:          r,
					s:          s,
				},
			}
			//safeSignatures = append(safeSignatures, safe_signature)
		} else if signatureType == ETH_SIGN {
			safe_signature = SafeSignatureEthSign{
				&safeSignature_{
					safeTxHash: safeTxHash,
					v:          v,
					r:          r,
					s:          s,
				},
			}
			//safeSignatures = append(safeSignatures, safe_signature)
		} else if signatureType == EOA {
			safe_signature = SafeSignatureEOA{
				&safeSignature_{
					safeTxHash: safeTxHash,
					v:          v,
					r:          r,
					s:          s,
				},
			}
		}
		safeSignatures = append(safeSignatures, safe_signature)
	}
	return safeSignatures
}

func signatureSplit(signatures []byte, pos int) (*big.Int, *big.Int, *big.Int, error) {
	signaturePos := SignatureSize * pos
	if len(signatures[signaturePos:signaturePos+SignatureSize]) < 65 {
		return nil, nil, nil, fmt.Errorf("invalid signatures provided for splitSignature function")
	}

	r := new(big.Int).SetBytes(signatures[signaturePos : 32+signaturePos])
	s := new(big.Int).SetBytes(signatures[32+signaturePos : 64+signaturePos])
	v := new(big.Int).SetBytes(signatures[64+signaturePos : crypto.RecoveryIDOffset+signaturePos+1])
	return v, r, s, nil
}

func SignMessageHash(hash common.Hash, sKey *ecdsa.PrivateKey) ([]byte, error) {
	signature, err := crypto.Sign(hash.Bytes(), sKey)
	if err != nil {
		return nil, err
	}

	signature[crypto.RecoveryIDOffset] += 27
	return signature, nil
}

func SplitAndValidateSignature(signature []byte, pos int) (*big.Int, *big.Int, *big.Int, error) {
	v, r, s, err := signatureSplit(signature, pos)
	if err != nil {
		return nil, nil, nil, err
	}

	if !crypto.ValidateSignatureValues(byte(big.NewInt(v.Int64()-27).Int64()), r, s, false) {
		return nil, nil, nil, err
	}
	return v, r, s, nil
}

func (signature *SafeSignatureEOA) Owner() *common.Address {
	_sig := SigFromVRS(signature.safeSignature.v, signature.safeSignature.r, signature.safeSignature.s)
	_sig[crypto.RecoveryIDOffset] -= 27
	pubKey, err := crypto.SigToPub(signature.safeSignature.safeTxHash, _sig)
	if err != nil {
		return nil
	}
	addr := crypto.PubkeyToAddress(*pubKey)
	return &addr
}
func (signaure *SafeSignatureEOA) IsValid() bool {
	return true
}
func (signaure *SafeSignatureEOA) SignatureType() SafeSignatureType {
	return EOA
}

func (signature *SafeSignatureEthSign) Owner() *common.Address {
	_sig := SigFromVRS(signature.safeSignature.v, signature.safeSignature.r, signature.safeSignature.s)
	_sig[crypto.RecoveryIDOffset] -= 27
	pubKey, err := crypto.SigToPub(signature.safeSignature.safeTxHash, _sig)
	if err != nil {
		return nil
	}
	addr := crypto.PubkeyToAddress(*pubKey)
	return &addr
}
func (signaure *SafeSignatureEthSign) IsValid() bool {
	return true
}
func (signaure *SafeSignatureEthSign) SignatureType() SafeSignatureType {
	return ETH_SIGN
}

func (signature *SafeSignatureApprovedHash) Owner() *common.Address {
	// TODO: Test uint_to_address
	return uintToAddress(signature.safeSignature.r)
}
func (signaure *SafeSignatureApprovedHash) SignatureType() SafeSignatureType {
	return APPROVED_HASH
}
func (signature *SafeSignatureApprovedHash) IsValid(
	ethClient *eth.EthereumClient,
	safeContract *contracts.GnosisSafe,
) bool {
	res, err := safeContract.ApprovedHashes(
		nil, *signature.Owner(), common.BytesToHash(signature.safeSignature.safeTxHash))
	if err != nil {
		return false
	}
	return res.Uint64() == 1
}

func (signature *SafeSignatureContract) Owner() *common.Address {
	// TODO: Test uint_to_address
	return uintToAddress(signature.safeSignature.r)
}
func (signaure *SafeSignatureContract) SignatureType() SafeSignatureType {
	return CONTRACT_SIGNATURE
}
func (signaure *SafeSignatureContract) IsValid() bool {
	return false
}

// Convert a Solidity `uint` value to a checksummed `address`, removing invalid padding bytes if present
func uintToAddress(value *big.Int) *common.Address {
	uintTy, err := abi.NewType("uint", "uint", nil)
	if err != nil {
		return nil
	}
	addressTy, err := abi.NewType("address", "address", nil)
	if err != nil {
		return nil
	}
	arguments := abi.Arguments{
		{
			Type: uintTy,
		},
	}
	encoded, err := arguments.Pack(value)
	if err != nil {
		return nil
	}
	encoded_without_padding := append([]byte{0x00 * 12}, encoded[20:]...)
	arguments = abi.Arguments{
		{
			Type: addressTy,
		},
	}
	decoded, err := arguments.Unpack(encoded_without_padding)
	if err != nil {
		return nil
	}
	address_b := common.BytesToAddress(decoded[0].([]byte))
	return &address_b
}

func SigFromVRS(v, r, s *big.Int) []byte {
	if v == nil || r == nil || s == nil {
		return nil
	}
	return append(r.Bytes(), append(s.Bytes(), v.Bytes()...)...)
}
