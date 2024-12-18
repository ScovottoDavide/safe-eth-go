package eth

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math/big"
	"reflect"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case int:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func StringToAddress(address string) (*common.Address, error) {
	if IsValidAddress(address) {
		res := common.HexToAddress(address)
		return &res, nil
	}
	return nil, errors.New("invalid address format")
}

func TryAnyToAddress(iaddress interface{}) (*common.Address, error) {
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
	return address, nil
}

func GetCryptoPrivateKey(iprivateKey interface{}) (*ecdsa.PrivateKey, error) {
	privateKey := new(ecdsa.PrivateKey)
	err := error(nil)
	switch v := iprivateKey.(type) {
	case string:
		privateKey, err = crypto.ToECDSA(hexutil.MustDecode(v))
		if err != nil {
			return nil, err
		}
	case []byte:
		privateKey, err = crypto.ToECDSA(v)
		if err != nil {
			return nil, err
		}
	}
	return privateKey, nil
}

func IntToBigInt(value int) *big.Int {
	return big.NewInt(int64(value))
}

func MakeContractAddress(txFrom common.Address, txNonce uint64) common.Address {
	return crypto.CreateAddress(txFrom, txNonce)
}

func MakeContractAddress2(txFrom common.Address, saltNonce [32]byte, inithash []byte) common.Address {
	return crypto.CreateAddress2(txFrom, saltNonce, inithash)
}

func IsTransactionSuccessful(receipt *types.Receipt) bool {
	return receipt.Status == types.ReceiptStatusSuccessful
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

func RandomAddress() (*common.Address, error) {
	s_key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return AddressFromPrivKey(s_key.X.Bytes())
}
