package safe_signature

import (
	"fmt"
	"math/big"

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
	signature  []byte
	safeTxHash []byte
	v, r, s    *big.Int
}

type SafeSignatureContract struct {
	safeSignature     *safeSignature_
	contractSignature []byte
}

type SafeSignatureApprovedHash struct {
	safeSignature     *safeSignature_
	contractSignature []byte
}

type SafeSignatureEthSign struct {
	safeSignature     *safeSignature_
	contractSignature []byte
}

type SafeSignatureEOA struct {
	safeSignature     *safeSignature_
	contractSignature []byte
}

type SafeSignature interface {
	SafeSignatureContract | SafeSignatureApprovedHash | SafeSignatureEthSign | SafeSignatureEOA
}

func ParseSignatures[T SafeSignature](
	signatures []byte,
	safeTxHash []byte,
) []T {
	if signatures == nil {
		return nil
	}

	dataPosition := len(signatures) // For contract signatures, to stop parsing at data position
	var safeSignatures []T
	// var safe_signature T
	for i := 0; i < len(signatures); i += SignatureSize {
		if i >= dataPosition { // If contract signature data position is reached, stop
			break
		}
		signature := signatures[i : i+SignatureSize]
		v, r, s, err := signatureSplit(signature, 0)
		if err != nil {
			return nil
		}
		if !crypto.ValidateSignatureValues(v.Bytes()[0], r, s, false) {
			return nil
		}
		signatureType := SigTypeFromV(v.Uint64())
		if signatureType == CONTRACT_SIGNATURE {
			s_int := int(s.Uint64())
			if s_int < dataPosition {
				dataPosition = s_int
			}
			contract_signature_len := int(new(big.Int).SetBytes(signatures[s_int : s_int+32]).Uint64()) // Len size is 32 bytes
			contract_signature := signatures[s_int+32 : s_int+32+contract_signature_len]                // Skip array size (32 bytes)
			fmt.Println(contract_signature)
			safe_signature := SafeSignatureContract{
				&safeSignature_{
					signature:  signature,
					safeTxHash: safeTxHash,
				}, contract_signature,
			}
			safeSignatures = append(safeSignatures, T(safe_signature))
		} else if signatureType == APPROVED_HASH {
			safe_signature := SafeSignatureApprovedHash{
				&safeSignature_{
					signature:  signature,
					safeTxHash: safeTxHash,
				}, nil,
			}
			safeSignatures = append(safeSignatures, T(safe_signature))
		} else if signatureType == ETH_SIGN {
			safe_signature := SafeSignatureEthSign{
				&safeSignature_{
					signature:  signature,
					safeTxHash: safeTxHash,
				}, nil,
			}
			safeSignatures = append(safeSignatures, T(safe_signature))
		} else if signatureType == EOA {
			safe_signature := SafeSignatureEOA{
				&safeSignature_{
					signature:  signature,
					safeTxHash: safeTxHash,
				}, nil,
			}
			safeSignatures = append(safeSignatures, T(safe_signature))
		}
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
	v := new(big.Int).SetBytes(signatures[:64+signaturePos])
	return v, r, s, nil
}

func (signature *SafeSignatureEOA) Owner() (*common.Address, error) {
	pubKey, err := crypto.SigToPub(signature.safeSignature.safeTxHash, signature.safeSignature.signature)
	if err != nil {
		return nil, err
	}
	addr := crypto.PubkeyToAddress(*pubKey)
	return &addr, nil
}

func (signature *SafeSignatureEthSign) Owner() (*common.Address, error) {
	pubKey, err := crypto.SigToPub(signature.safeSignature.safeTxHash, signature.safeSignature.signature)
	if err != nil {
		return nil, err
	}
	addr := crypto.PubkeyToAddress(*pubKey)
	return &addr, nil
}

func (signature *SafeSignatureApprovedHash) Owner() (*common.Address, error) {
	// TODO: uint_to_address
	return nil, nil
}
