// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GnosisSafeMetaData contains all meta data concerning the GnosisSafe contract.
var GnosisSafeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"requiredTxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600160048190555061611c80620000296000396000f3fe6080604052600436106101dc5760003560e01c8063affed0e011610102578063e19a9dd911610095578063f08a032311610064578063f08a0323146107de578063f698da2514610807578063f8dc5dd914610832578063ffa1ad741461085b57610231565b8063e19a9dd914610724578063e318b52b1461074d578063e75235b814610776578063e86637db146107a157610231565b8063cc2f8452116100d1578063cc2f845214610657578063d4d9bdcd14610695578063d8d11f78146106be578063e009cfde146106fb57610231565b8063affed0e01461059d578063b4faba09146105c8578063b63e800d146105f1578063c4ca3a9c1461061a57610231565b80635624b25b1161017a5780636a761202116101495780636a761202146104dc5780637d8329741461050c578063934f3a1114610549578063a0e67e2b1461057257610231565b80635624b25b146104105780635ae6bd371461044d578063610b59251461048a578063694e80c3146104b357610231565b80632f54bf6e116101b65780632f54bf6e1461032d5780633408e4701461036a578063468721a7146103955780635229073f146103d257610231565b80630d582f131461029e57806312fb68e0146102c75780632d9ad53d146102f057610231565b36610231573373ffffffffffffffffffffffffffffffffffffffff167f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d3460405161022791906156aa565b60405180910390a2005b34801561023d57600080fd5b5060007f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d560001b905080548061027257600080f35b36600080373360601b365260008060143601600080855af13d6000803e80610299573d6000fd5b3d6000f35b3480156102aa57600080fd5b506102c560048036038101906102c091906142c2565b610886565b005b3480156102d357600080fd5b506102ee60048036038101906102e99190614794565b610c00565b005b3480156102fc57600080fd5b506103176004803603810190610312919061412f565b611207565b6040516103249190615110565b60405180910390f35b34801561033957600080fd5b50610354600480360381019061034f919061412f565b6112d9565b6040516103619190615110565b60405180910390f35b34801561037657600080fd5b5061037f6113ab565b60405161038c91906156aa565b60405180910390f35b3480156103a157600080fd5b506103bc60048036038101906103b79190614586565b6113b8565b6040516103c99190615110565b60405180910390f35b3480156103de57600080fd5b506103f960048036038101906103f49190614586565b61156b565b60405161040792919061512b565b60405180910390f35b34801561041c57600080fd5b5061043760048036038101906104329190614879565b6115a1565b60405161044491906152ef565b60405180910390f35b34801561045957600080fd5b50610474600480360381019061046f91906146ec565b61165f565b60405161048191906156aa565b60405180910390f35b34801561049657600080fd5b506104b160048036038101906104ac919061412f565b611677565b005b3480156104bf57600080fd5b506104da60048036038101906104d59190614850565b61198d565b005b6104f660048036038101906104f1919061437e565b611a61565b6040516105039190615110565b60405180910390f35b34801561051857600080fd5b50610533600480360381019061052e9190614232565b611e22565b60405161054091906156aa565b60405180910390f35b34801561055557600080fd5b50610570600480360381019061056b9190614715565b611e47565b005b34801561057e57600080fd5b50610587611ea3565b60405161059491906150be565b60405180910390f35b3480156105a957600080fd5b506105b26120a7565b6040516105bf91906156aa565b60405180910390f35b3480156105d457600080fd5b506105ef60048036038101906105ea919061426e565b6120ad565b005b3480156105fd57600080fd5b5061061860048036038101906106139190614601565b6120cf565b005b34801561062657600080fd5b50610641600480360381019061063c91906142fe565b612222565b60405161064e91906156aa565b60405180910390f35b34801561066357600080fd5b5061067e600480360381019061067991906142c2565b6122ee565b60405161068c9291906150e0565b60405180910390f35b3480156106a157600080fd5b506106bc60048036038101906106b791906146ec565b61253b565b005b3480156106ca57600080fd5b506106e560048036038101906106e0919061448e565b6126a7565b6040516106f2919061515b565b60405180910390f35b34801561070757600080fd5b50610722600480360381019061071d9190614158565b6126d4565b005b34801561073057600080fd5b5061074b6004803603810190610746919061412f565b6129e9565b005b34801561075957600080fd5b50610774600480360381019061076f9190614194565b612a57565b005b34801561078257600080fd5b5061078b612fd1565b60405161079891906156aa565b60405180910390f35b3480156107ad57600080fd5b506107c860048036038101906107c3919061448e565b612fdb565b6040516107d591906152ef565b60405180910390f35b3480156107ea57600080fd5b506108056004803603810190610800919061412f565b61309d565b005b34801561081357600080fd5b5061081c6130e8565b604051610829919061515b565b60405180910390f35b34801561083e57600080fd5b50610859600480360381019061085491906141e3565b613144565b005b34801561086757600080fd5b506108706134d6565b60405161087d9190615348565b60405180910390f35b61088e61350f565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580156108f85750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b801561093057503073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b61096f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109669061544a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a34906155ea565b60405180910390fd5b60026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060036000815480929190610bad90615ac8565b91905055507f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2682604051610be19190614f71565b60405180910390a18060045414610bfc57610bfb8161198d565b5b5050565b610c1460418261357f90919063ffffffff16565b82511015610c57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4e9061566a565b60405180910390fd5b6000808060008060005b868110156111fb57610c7388826135c4565b80945081955082965050505060008460ff161415610edd578260001c9450610ca560418861357f90919063ffffffff16565b8260001c1015610cea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce19061562a565b60405180910390fd5b8751610d0360208460001c6135f390919063ffffffff16565b1115610d44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3b9061540a565b60405180910390fd5b60006020838a01015190508851610d7a82610d6c60208760001c6135f390919063ffffffff16565b6135f390919063ffffffff16565b1115610dbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db29061536a565b60405180910390fd5b60606020848b010190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168773ffffffffffffffffffffffffffffffffffffffff166320c13b0b8d846040518363ffffffff1660e01b8152600401610e27929190615311565b60206040518083038186803b158015610e3f57600080fd5b505afa158015610e53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e779190614827565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610ed6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ecd906153aa565b60405180910390fd5b50506110a6565b60018460ff161415610fbe578260001c94508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610f7a57506000600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008c81526020019081526020016000205414155b610fb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb0906155ca565b60405180910390fd5b6110a5565b601e8460ff1611156110525760018a604051602001610fdd9190614f30565b6040516020818303038152906040528051906020012060048661100091906158e5565b85856040516000815260200160405260405161101f94939291906152aa565b6020604051602081039080840390855afa158015611041573d6000803e3d6000fd5b5050506020604051035194506110a4565b60018a8585856040516000815260200160405260405161107594939291906152aa565b6020604051602081039080840390855afa158015611097573d6000803e3d6000fd5b5050506020604051035194505b5b5b8573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1611801561116d5750600073ffffffffffffffffffffffffffffffffffffffff16600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b80156111a65750600173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614155b6111e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111dc9061538a565b60405180910390fd5b84955080806111f390615ac8565b915050610c61565b50505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff16600173ffffffffffffffffffffffffffffffffffffffff16141580156112d25750600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b6000600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580156113a45750600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b6000804690508091505090565b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580156114835750600073ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b6114c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b99061558a565b60405180910390fd5b6114cf858585855a61361b565b9050801561151f573373ffffffffffffffffffffffffffffffffffffffff167f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb860405160405180910390a2611563565b3373ffffffffffffffffffffffffffffffffffffffff167facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37560405160405180910390a25b949350505050565b6000606061157b868686866113b8565b915060405160203d0181016040523d81523d6000602083013e8091505094509492505050565b606060006020836115b29190615857565b67ffffffffffffffff8111156115f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156116235781602001600182028036833780820191505090505b50905060005b838110156116545780850154806020830260208501015250808061164c90615ac8565b915050611629565b508091505092915050565b60076020528060005260406000206000915090505481565b61167f61350f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156116e95750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b611728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171f9061564a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146117f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117ed9061560a565b60405180910390fd5b60016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440816040516119829190614f71565b60405180910390a150565b61199561350f565b6003548111156119da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d1906153ea565b60405180910390fd5b6001811015611a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a159061556a565b60405180910390fd5b806004819055507f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93600454604051611a5691906156aa565b60405180910390a150565b6000806000611a7b8e8e8e8e8e8e8e8e8e8e600554612fdb565b905060056000815480929190611a9090615ac8565b919050555080805190602001209150611aaa828286611e47565b506000611ab56136c1565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b6d578073ffffffffffffffffffffffffffffffffffffffff166375f0bb528f8f8f8f8f8f8f8f8f8f8f336040518d63ffffffff1660e01b8152600401611b3a9c9b9a99989796959493929190614fb5565b600060405180830381600087803b158015611b5457600080fd5b505af1158015611b68573d6000803e3d6000fd5b505050505b6101f4611ba86109c48b611b8191906157d0565b603f60408d611b909190615857565b611b9a9190615826565b6136f290919063ffffffff16565b611bb291906157d0565b5a1015611bf4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611beb9061554a565b60405180910390fd5b60005a9050611c668f8f8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e60008d14611c52578e611c61565b6109c45a611c6091906158b1565b5b61361b565b9350611c7b5a8261370c90919063ffffffff16565b90508380611c8a575060008a14155b80611c96575060008814155b611cd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ccc9061552a565b60405180910390fd5b600080891115611cef57611cec828b8b8b8b613735565b90505b8415611d33577f442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e8482604051611d2692919061524a565b60405180910390a1611d6d565b7f23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d238482604051611d6492919061524a565b60405180910390a15b5050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611e11578073ffffffffffffffffffffffffffffffffffffffff16639327136883856040518363ffffffff1660e01b8152600401611dde929190615221565b600060405180830381600087803b158015611df857600080fd5b505af1158015611e0c573d6000803e3d6000fd5b505050505b50509b9a5050505050505050505050565b6008602052816000526040600020602052806000526040600020600091509150505481565b6000600454905060008111611e91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e889061550a565b60405180910390fd5b611e9d84848484610c00565b50505050565b6060600060035467ffffffffffffffff811115611ee9577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015611f175781602001602082028036833780820191505090505b50905060008060026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461209e5780838381518110611fef577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050818061209690615ac8565b925050611f81565b82935050505090565b60055481565b600080825160208401855af4806000523d6020523d600060403e60403d016000fd5b61211a8a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050896138d5565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16146121585761215784613d09565b5b6121a68787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050613d38565b60008211156121c0576121be82600060018685613735565b505b3373ffffffffffffffffffffffffffffffffffffffff167f141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a88b8b8b8b8960405161220e959493929190615070565b60405180910390a250505050505050505050565b6000805a9050612279878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050865a61361b565b61228257600080fd5b60005a8261229091906158b1565b9050806040516020016122a39190614f56565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122e59190615348565b60405180910390fd5b606060008267ffffffffffffffff811115612332577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156123605781602001602082028036833780820191505090505b509150600080600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156124335750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561243e57508482105b1561252c578084838151811061247d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050818061252490615ac8565b9250506123c9565b80925081845250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff16600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561260a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612601906153ca565b60405180910390fd5b6001600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16817ff2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c60405160405180910390a350565b60006126bc8c8c8c8c8c8c8c8c8c8c8c612fdb565b8051906020012090509b9a5050505050505050505050565b6126dc61350f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156127465750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b612785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161277c9061564a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612852576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612849906154aa565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276816040516129dd9190614f71565b60405180910390a15050565b6129f161350f565b60007f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c860001b90508181557f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa282604051612a4b9190614f71565b60405180910390a15050565b612a5f61350f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015612ac95750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015612b0157503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b612b40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b379061544a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612c0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c05906155ea565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015612c785750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b612cb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cae9061544a565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612d84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d7b9061568a565b60405180910390fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf82604051612f8d9190614f71565b60405180910390a17f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2681604051612fc49190614f71565b60405180910390a1505050565b6000600454905090565b606060007fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d860001b8d8d8d8d604051613015929190614f17565b60405180910390208c8c8c8c8c8c8c60405160200161303e9b9a99989796959493929190615176565b604051602081830303815290604052805190602001209050601960f81b600160f81b6130686130e8565b8360405160200161307c9493929190614ec9565b6040516020818303038152906040529150509b9a5050505050505050505050565b6130a561350f565b6130ae81613d09565b7f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0816040516130dd9190614f71565b60405180910390a150565b60007f47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a7946921860001b6131166113ab565b3060405160200161312993929190615273565b60405160208183030381529060405280519060200120905090565b61314c61350f565b80600160035461315c91906158b1565b101561319d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613194906153ea565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580156132075750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b613246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161323d9061544a565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613313576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161330a9061568a565b60405180910390fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003600081548092919061348290615a6d565b91905055507ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf826040516134b69190614f71565b60405180910390a180600454146134d1576134d08161198d565b5b505050565b6040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461357d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613574906155aa565b60405180910390fd5b565b60008083141561359257600090506135be565b600082846135a09190615857565b90508284826135af9190615826565b146135b957600080fd5b809150505b92915050565b60008060008360410260208101860151925060408101860151915060ff60418201870151169350509250925092565b600080828461360291906157d0565b90508381101561361157600080fd5b8091505092915050565b6000600180811115613656577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b83600181111561368f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156136a8576000808551602087018986f490506136b8565b600080855160208701888a87f190505b95945050505050565b6000807f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c860001b9050805491505090565b6000818310156137025781613704565b825b905092915050565b60008282111561371b57600080fd5b6000828461372991906158b1565b90508091505092915050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146137725782613774565b325b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415613859576137de3a86106137bb573a6137bd565b855b6137d0888a6135f390919063ffffffff16565b61357f90919063ffffffff16565b91508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050613854576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161384b9061548a565b60405180910390fd5b6138cb565b61387e85613870888a6135f390919063ffffffff16565b61357f90919063ffffffff16565b915061388b848284613f0d565b6138ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138c1906154ca565b60405180910390fd5b5b5095945050505050565b60006004541461391a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139119061542a565b60405180910390fd5b815181111561395e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613955906153ea565b60405180910390fd5b60018110156139a2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139999061556a565b60405180910390fd5b60006001905060005b8351811015613c755760008482815181106139ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015613a635750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015613a9b57503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015613ad357508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b613b12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b099061544a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613be0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bd7906155ea565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550809250508080613c6d90615ac8565b9150506139ab565b506001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550825160038190555081600481905550505050565b60007f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d560001b90508181555050565b600073ffffffffffffffffffffffffffffffffffffffff1660016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613e07576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613dfe906154ea565b60405180910390fd5b6001806000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614613f0957613ec98260008360015a61361b565b613f08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613eff9061546a565b60405180910390fd5b5b5050565b60008063a9059cbb8484604051602401613f28929190614f8c565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050602060008251602084016000896127105a03f13d60008114613f985760208114613fa05760009350613fab565b819350613fab565b600051158215171593505b5050509392505050565b6000613fc8613fc3846156ea565b6156c5565b905082815260208101848484011115613fe057600080fd5b613feb848285615a2b565b509392505050565b60008135905061400281616063565b92915050565b6000813590506140178161607a565b92915050565b60008083601f84011261402f57600080fd5b8235905067ffffffffffffffff81111561404857600080fd5b60208301915083602082028301111561406057600080fd5b9250929050565b60008135905061407681616091565b92915050565b60008151905061408b816160a8565b92915050565b60008083601f8401126140a357600080fd5b8235905067ffffffffffffffff8111156140bc57600080fd5b6020830191508360018202830111156140d457600080fd5b9250929050565b600082601f8301126140ec57600080fd5b81356140fc848260208601613fb5565b91505092915050565b600081359050614114816160bf565b92915050565b600081359050614129816160cf565b92915050565b60006020828403121561414157600080fd5b600061414f84828501613ff3565b91505092915050565b6000806040838503121561416b57600080fd5b600061417985828601613ff3565b925050602061418a85828601613ff3565b9150509250929050565b6000806000606084860312156141a957600080fd5b60006141b786828701613ff3565b93505060206141c886828701613ff3565b92505060406141d986828701613ff3565b9150509250925092565b6000806000606084860312156141f857600080fd5b600061420686828701613ff3565b935050602061421786828701613ff3565b92505060406142288682870161411a565b9150509250925092565b6000806040838503121561424557600080fd5b600061425385828601613ff3565b925050602061426485828601614067565b9150509250929050565b6000806040838503121561428157600080fd5b600061428f85828601613ff3565b925050602083013567ffffffffffffffff8111156142ac57600080fd5b6142b8858286016140db565b9150509250929050565b600080604083850312156142d557600080fd5b60006142e385828601613ff3565b92505060206142f48582860161411a565b9150509250929050565b60008060008060006080868803121561431657600080fd5b600061432488828901613ff3565b95505060206143358882890161411a565b945050604086013567ffffffffffffffff81111561435257600080fd5b61435e88828901614091565b9350935050606061437188828901614105565b9150509295509295909350565b60008060008060008060008060008060006101408c8e0312156143a057600080fd5b60006143ae8e828f01613ff3565b9b505060206143bf8e828f0161411a565b9a505060408c013567ffffffffffffffff8111156143dc57600080fd5b6143e88e828f01614091565b995099505060606143fb8e828f01614105565b975050608061440c8e828f0161411a565b96505060a061441d8e828f0161411a565b95505060c061442e8e828f0161411a565b94505060e061443f8e828f01613ff3565b9350506101006144518e828f01614008565b9250506101208c013567ffffffffffffffff81111561446f57600080fd5b61447b8e828f016140db565b9150509295989b509295989b9093969950565b60008060008060008060008060008060006101408c8e0312156144b057600080fd5b60006144be8e828f01613ff3565b9b505060206144cf8e828f0161411a565b9a505060408c013567ffffffffffffffff8111156144ec57600080fd5b6144f88e828f01614091565b9950995050606061450b8e828f01614105565b975050608061451c8e828f0161411a565b96505060a061452d8e828f0161411a565b95505060c061453e8e828f0161411a565b94505060e061454f8e828f01613ff3565b9350506101006145618e828f01613ff3565b9250506101206145738e828f0161411a565b9150509295989b509295989b9093969950565b6000806000806080858703121561459c57600080fd5b60006145aa87828801613ff3565b94505060206145bb8782880161411a565b935050604085013567ffffffffffffffff8111156145d857600080fd5b6145e4878288016140db565b92505060606145f587828801614105565b91505092959194509250565b6000806000806000806000806000806101008b8d03121561462157600080fd5b60008b013567ffffffffffffffff81111561463b57600080fd5b6146478d828e0161401d565b9a509a5050602061465a8d828e0161411a565b985050604061466b8d828e01613ff3565b97505060608b013567ffffffffffffffff81111561468857600080fd5b6146948d828e01614091565b965096505060806146a78d828e01613ff3565b94505060a06146b88d828e01613ff3565b93505060c06146c98d828e0161411a565b92505060e06146da8d828e01614008565b9150509295989b9194979a5092959850565b6000602082840312156146fe57600080fd5b600061470c84828501614067565b91505092915050565b60008060006060848603121561472a57600080fd5b600061473886828701614067565b935050602084013567ffffffffffffffff81111561475557600080fd5b614761868287016140db565b925050604084013567ffffffffffffffff81111561477e57600080fd5b61478a868287016140db565b9150509250925092565b600080600080608085870312156147aa57600080fd5b60006147b887828801614067565b945050602085013567ffffffffffffffff8111156147d557600080fd5b6147e1878288016140db565b935050604085013567ffffffffffffffff8111156147fe57600080fd5b61480a878288016140db565b925050606061481b8782880161411a565b91505092959194509250565b60006020828403121561483957600080fd5b60006148478482850161407c565b91505092915050565b60006020828403121561486257600080fd5b60006148708482850161411a565b91505092915050565b6000806040838503121561488c57600080fd5b600061489a8582860161411a565b92505060206148ab8582860161411a565b9150509250929050565b60006148c183836148dc565b60208301905092915050565b6148d68161592b565b82525050565b6148e581615919565b82525050565b6148f481615919565b82525050565b60006149068385615770565b93506149118261571b565b8060005b8581101561494a5761492782846157b9565b61493188826148b5565b975061493c83615756565b925050600181019050614915565b5085925050509392505050565b600061496282615735565b61496c8185615770565b935061497783615725565b8060005b838110156149a857815161498f88826148b5565b975061499a83615763565b92505060018101905061497b565b5085935050505092915050565b6149be8161593d565b82525050565b6149d56149d082615949565b615b11565b82525050565b6149e481615975565b82525050565b6149fb6149f682615975565b615b1b565b82525050565b6000614a0d8385615781565b9350614a1a838584615a2b565b614a2383615beb565b840190509392505050565b6000614a3a8385615792565b9350614a47838584615a2b565b82840190509392505050565b6000614a5e82615740565b614a688185615781565b9350614a78818560208601615a3a565b614a8181615beb565b840191505092915050565b614a95816159f5565b82525050565b614aa481615a19565b82525050565b6000614ab58261574b565b614abf818561579d565b9350614acf818560208601615a3a565b614ad881615beb565b840191505092915050565b6000614af060058361579d565b9150614afb82615bfc565b602082019050919050565b6000614b1360058361579d565b9150614b1e82615c25565b602082019050919050565b6000614b36601c836157ae565b9150614b4182615c4e565b601c82019050919050565b6000614b5960058361579d565b9150614b6482615c77565b602082019050919050565b6000614b7c60058361579d565b9150614b8782615ca0565b602082019050919050565b6000614b9f60058361579d565b9150614baa82615cc9565b602082019050919050565b6000614bc260058361579d565b9150614bcd82615cf2565b602082019050919050565b6000614be560058361579d565b9150614bf082615d1b565b602082019050919050565b6000614c0860058361579d565b9150614c1382615d44565b602082019050919050565b6000614c2b60058361579d565b9150614c3682615d6d565b602082019050919050565b6000614c4e60058361579d565b9150614c5982615d96565b602082019050919050565b6000614c7160058361579d565b9150614c7c82615dbf565b602082019050919050565b6000614c9460058361579d565b9150614c9f82615de8565b602082019050919050565b6000614cb760058361579d565b9150614cc282615e11565b602082019050919050565b6000614cda60058361579d565b9150614ce582615e3a565b602082019050919050565b6000614cfd60058361579d565b9150614d0882615e63565b602082019050919050565b6000614d2060058361579d565b9150614d2b82615e8c565b602082019050919050565b6000614d4360058361579d565b9150614d4e82615eb5565b602082019050919050565b6000614d6660058361579d565b9150614d7182615ede565b602082019050919050565b6000614d8960058361579d565b9150614d9482615f07565b602082019050919050565b6000614dac60058361579d565b9150614db782615f30565b602082019050919050565b6000614dcf60058361579d565b9150614dda82615f59565b602082019050919050565b6000614df260058361579d565b9150614dfd82615f82565b602082019050919050565b6000614e1560058361579d565b9150614e2082615fab565b602082019050919050565b6000614e3860058361579d565b9150614e4382615fd4565b602082019050919050565b6000614e5b60058361579d565b9150614e6682615ffd565b602082019050919050565b6000614e7e60058361579d565b9150614e8982616026565b602082019050919050565b614e9d816159de565b82525050565b614eb4614eaf826159de565b615b25565b82525050565b614ec3816159e8565b82525050565b6000614ed582876149c4565b600182019150614ee582866149c4565b600182019150614ef582856149ea565b602082019150614f0582846149ea565b60208201915081905095945050505050565b6000614f24828486614a2e565b91508190509392505050565b6000614f3b82614b29565b9150614f4782846149ea565b60208201915081905092915050565b6000614f628284614ea3565b60208201915081905092915050565b6000602082019050614f8660008301846148eb565b92915050565b6000604082019050614fa160008301856148eb565b614fae6020830184614e94565b9392505050565b600061016082019050614fcb600083018f6148eb565b614fd8602083018e614e94565b8181036040830152614feb818c8e614a01565b9050614ffa606083018b614a9b565b615007608083018a614e94565b61501460a0830189614e94565b61502160c0830188614e94565b61502e60e08301876148eb565b61503c6101008301866148cd565b81810361012083015261504f8185614a53565b905061505f6101408301846148eb565b9d9c50505050505050505050505050565b6000608082019050818103600083015261508b8187896148fa565b905061509a6020830186614e94565b6150a760408301856148eb565b6150b460608301846148eb565b9695505050505050565b600060208201905081810360008301526150d88184614957565b905092915050565b600060408201905081810360008301526150fa8185614957565b905061510960208301846148eb565b9392505050565b600060208201905061512560008301846149b5565b92915050565b600060408201905061514060008301856149b5565b81810360208301526151528184614a53565b90509392505050565b600060208201905061517060008301846149db565b92915050565b60006101608201905061518c600083018e6149db565b615199602083018d6148eb565b6151a6604083018c614e94565b6151b3606083018b6149db565b6151c0608083018a614a9b565b6151cd60a0830189614e94565b6151da60c0830188614e94565b6151e760e0830187614e94565b6151f56101008301866148eb565b6152036101208301856148eb565b615211610140830184614e94565b9c9b505050505050505050505050565b600060408201905061523660008301856149db565b61524360208301846149b5565b9392505050565b600060408201905061525f60008301856149db565b61526c6020830184614e94565b9392505050565b600060608201905061528860008301866149db565b6152956020830185614e94565b6152a26040830184614a8c565b949350505050565b60006080820190506152bf60008301876149db565b6152cc6020830186614eba565b6152d960408301856149db565b6152e660608301846149db565b95945050505050565b600060208201905081810360008301526153098184614a53565b905092915050565b6000604082019050818103600083015261532b8185614a53565b9050818103602083015261533f8184614a53565b90509392505050565b600060208201905081810360008301526153628184614aaa565b905092915050565b6000602082019050818103600083015261538381614ae3565b9050919050565b600060208201905081810360008301526153a381614b06565b9050919050565b600060208201905081810360008301526153c381614b4c565b9050919050565b600060208201905081810360008301526153e381614b6f565b9050919050565b6000602082019050818103600083015261540381614b92565b9050919050565b6000602082019050818103600083015261542381614bb5565b9050919050565b6000602082019050818103600083015261544381614bd8565b9050919050565b6000602082019050818103600083015261546381614bfb565b9050919050565b6000602082019050818103600083015261548381614c1e565b9050919050565b600060208201905081810360008301526154a381614c41565b9050919050565b600060208201905081810360008301526154c381614c64565b9050919050565b600060208201905081810360008301526154e381614c87565b9050919050565b6000602082019050818103600083015261550381614caa565b9050919050565b6000602082019050818103600083015261552381614ccd565b9050919050565b6000602082019050818103600083015261554381614cf0565b9050919050565b6000602082019050818103600083015261556381614d13565b9050919050565b6000602082019050818103600083015261558381614d36565b9050919050565b600060208201905081810360008301526155a381614d59565b9050919050565b600060208201905081810360008301526155c381614d7c565b9050919050565b600060208201905081810360008301526155e381614d9f565b9050919050565b6000602082019050818103600083015261560381614dc2565b9050919050565b6000602082019050818103600083015261562381614de5565b9050919050565b6000602082019050818103600083015261564381614e08565b9050919050565b6000602082019050818103600083015261566381614e2b565b9050919050565b6000602082019050818103600083015261568381614e4e565b9050919050565b600060208201905081810360008301526156a381614e71565b9050919050565b60006020820190506156bf6000830184614e94565b92915050565b60006156cf6156e0565b90506156db8282615a97565b919050565b6000604051905090565b600067ffffffffffffffff82111561570557615704615bbc565b5b61570e82615beb565b9050602081019050919050565b6000819050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b60006157c86020840184613ff3565b905092915050565b60006157db826159de565b91506157e6836159de565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561581b5761581a615b2f565b5b828201905092915050565b6000615831826159de565b915061583c836159de565b92508261584c5761584b615b5e565b5b828204905092915050565b6000615862826159de565b915061586d836159de565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156158a6576158a5615b2f565b5b828202905092915050565b60006158bc826159de565b91506158c7836159de565b9250828210156158da576158d9615b2f565b5b828203905092915050565b60006158f0826159e8565b91506158fb836159e8565b92508282101561590e5761590d615b2f565b5b828203905092915050565b6000615924826159be565b9050919050565b6000615936826159be565b9050919050565b60008115159050919050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b60008190506159b98261604f565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000615a0082615a07565b9050919050565b6000615a12826159be565b9050919050565b6000615a24826159ab565b9050919050565b82818337600083830152505050565b60005b83811015615a58578082015181840152602081019050615a3d565b83811115615a67576000848401525b50505050565b6000615a78826159de565b91506000821415615a8c57615a8b615b2f565b5b600182039050919050565b615aa082615beb565b810181811067ffffffffffffffff82111715615abf57615abe615bbc565b5b80604052505050565b6000615ad3826159de565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415615b0657615b05615b2f565b5b600182019050919050565b6000819050919050565b6000819050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4753303233000000000000000000000000000000000000000000000000000000600082015250565b7f4753303236000000000000000000000000000000000000000000000000000000600082015250565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b7f4753303234000000000000000000000000000000000000000000000000000000600082015250565b7f4753303330000000000000000000000000000000000000000000000000000000600082015250565b7f4753323031000000000000000000000000000000000000000000000000000000600082015250565b7f4753303232000000000000000000000000000000000000000000000000000000600082015250565b7f4753323030000000000000000000000000000000000000000000000000000000600082015250565b7f4753323033000000000000000000000000000000000000000000000000000000600082015250565b7f4753303030000000000000000000000000000000000000000000000000000000600082015250565b7f4753303131000000000000000000000000000000000000000000000000000000600082015250565b7f4753313033000000000000000000000000000000000000000000000000000000600082015250565b7f4753303132000000000000000000000000000000000000000000000000000000600082015250565b7f4753313030000000000000000000000000000000000000000000000000000000600082015250565b7f4753303031000000000000000000000000000000000000000000000000000000600082015250565b7f4753303133000000000000000000000000000000000000000000000000000000600082015250565b7f4753303130000000000000000000000000000000000000000000000000000000600082015250565b7f4753323032000000000000000000000000000000000000000000000000000000600082015250565b7f4753313034000000000000000000000000000000000000000000000000000000600082015250565b7f4753303331000000000000000000000000000000000000000000000000000000600082015250565b7f4753303235000000000000000000000000000000000000000000000000000000600082015250565b7f4753323034000000000000000000000000000000000000000000000000000000600082015250565b7f4753313032000000000000000000000000000000000000000000000000000000600082015250565b7f4753303231000000000000000000000000000000000000000000000000000000600082015250565b7f4753313031000000000000000000000000000000000000000000000000000000600082015250565b7f4753303230000000000000000000000000000000000000000000000000000000600082015250565b7f4753323035000000000000000000000000000000000000000000000000000000600082015250565b600281106160605761605f615b8d565b5b50565b61606c81615919565b811461607757600080fd5b50565b6160838161592b565b811461608e57600080fd5b50565b61609a81615975565b81146160a557600080fd5b50565b6160b18161597f565b81146160bc57600080fd5b50565b600281106160cc57600080fd5b50565b6160d8816159de565b81146160e357600080fd5b5056fea26469706673582212204efbc10a3922e4b014d3e2dd8aaee15b483391b8b90741b34c9e84cd8bfd1a2264736f6c63430008040033",
}

// GnosisSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeMetaData.ABI instead.
var GnosisSafeABI = GnosisSafeMetaData.ABI

// GnosisSafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GnosisSafeMetaData.Bin instead.
var GnosisSafeBin = GnosisSafeMetaData.Bin

// DeployGnosisSafe deploys a new Ethereum contract, binding an instance of GnosisSafe to it.
func DeployGnosisSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GnosisSafe, error) {
	parsed, err := GnosisSafeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GnosisSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GnosisSafe{GnosisSafeCaller: GnosisSafeCaller{contract: contract}, GnosisSafeTransactor: GnosisSafeTransactor{contract: contract}, GnosisSafeFilterer: GnosisSafeFilterer{contract: contract}}, nil
}

// GnosisSafe is an auto generated Go binding around an Ethereum contract.
type GnosisSafe struct {
	GnosisSafeCaller     // Read-only binding to the contract
	GnosisSafeTransactor // Write-only binding to the contract
	GnosisSafeFilterer   // Log filterer for contract events
}

// GnosisSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeSession struct {
	Contract     *GnosisSafe       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeCallerSession struct {
	Contract *GnosisSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GnosisSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeTransactorSession struct {
	Contract     *GnosisSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GnosisSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeRaw struct {
	Contract *GnosisSafe // Generic contract binding to access the raw methods on
}

// GnosisSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeCallerRaw struct {
	Contract *GnosisSafeCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeTransactorRaw struct {
	Contract *GnosisSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafe creates a new instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafe(address common.Address, backend bind.ContractBackend) (*GnosisSafe, error) {
	contract, err := bindGnosisSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafe{GnosisSafeCaller: GnosisSafeCaller{contract: contract}, GnosisSafeTransactor: GnosisSafeTransactor{contract: contract}, GnosisSafeFilterer: GnosisSafeFilterer{contract: contract}}, nil
}

// NewGnosisSafeCaller creates a new read-only instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeCaller, error) {
	contract, err := bindGnosisSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeCaller{contract: contract}, nil
}

// NewGnosisSafeTransactor creates a new write-only instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeTransactor, error) {
	contract, err := bindGnosisSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeTransactor{contract: contract}, nil
}

// NewGnosisSafeFilterer creates a new log filterer instance of GnosisSafe, bound to a specific deployed contract.
func NewGnosisSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeFilterer, error) {
	contract, err := bindGnosisSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeFilterer{contract: contract}, nil
}

// bindGnosisSafe binds a generic wrapper to an already deployed contract.
func bindGnosisSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GnosisSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafe *GnosisSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafe.Contract.GnosisSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafe *GnosisSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.Contract.GnosisSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafe *GnosisSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafe.Contract.GnosisSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafe *GnosisSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafe *GnosisSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafe *GnosisSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafe.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeSession) VERSION() (string, error) {
	return _GnosisSafe.Contract.VERSION(&_GnosisSafe.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_GnosisSafe *GnosisSafeCallerSession) VERSION() (string, error) {
	return _GnosisSafe.Contract.VERSION(&_GnosisSafe.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.ApprovedHashes(&_GnosisSafe.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.ApprovedHashes(&_GnosisSafe.CallOpts, arg0, arg1)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeCaller) CheckNSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafe.Contract.CheckNSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_GnosisSafe *GnosisSafeCallerSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _GnosisSafe.Contract.CheckNSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeCaller) CheckSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafe.Contract.CheckSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_GnosisSafe *GnosisSafeCallerSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _GnosisSafe.Contract.CheckSignatures(&_GnosisSafe.CallOpts, dataHash, data, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeSession) DomainSeparator() ([32]byte, error) {
	return _GnosisSafe.Contract.DomainSeparator(&_GnosisSafe.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_GnosisSafe *GnosisSafeCallerSession) DomainSeparator() ([32]byte, error) {
	return _GnosisSafe.Contract.DomainSeparator(&_GnosisSafe.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.EncodeTransactionData(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_GnosisSafe *GnosisSafeCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.EncodeTransactionData(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) GetChainId() (*big.Int, error) {
	return _GnosisSafe.Contract.GetChainId(&_GnosisSafe.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) GetChainId() (*big.Int, error) {
	return _GnosisSafe.Contract.GetChainId(&_GnosisSafe.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafe.Contract.GetModulesPaginated(&_GnosisSafe.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_GnosisSafe *GnosisSafeCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _GnosisSafe.Contract.GetModulesPaginated(&_GnosisSafe.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeSession) GetOwners() ([]common.Address, error) {
	return _GnosisSafe.Contract.GetOwners(&_GnosisSafe.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GnosisSafe *GnosisSafeCallerSession) GetOwners() ([]common.Address, error) {
	return _GnosisSafe.Contract.GetOwners(&_GnosisSafe.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.GetStorageAt(&_GnosisSafe.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_GnosisSafe *GnosisSafeCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _GnosisSafe.Contract.GetStorageAt(&_GnosisSafe.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) GetThreshold() (*big.Int, error) {
	return _GnosisSafe.Contract.GetThreshold(&_GnosisSafe.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) GetThreshold() (*big.Int, error) {
	return _GnosisSafe.Contract.GetThreshold(&_GnosisSafe.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafe.Contract.GetTransactionHash(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafe *GnosisSafeCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafe.Contract.GetTransactionHash(&_GnosisSafe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsModuleEnabled(&_GnosisSafe.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_GnosisSafe *GnosisSafeCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsModuleEnabled(&_GnosisSafe.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeSession) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsOwner(&_GnosisSafe.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_GnosisSafe *GnosisSafeCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _GnosisSafe.Contract.IsOwner(&_GnosisSafe.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) Nonce() (*big.Int, error) {
	return _GnosisSafe.Contract.Nonce(&_GnosisSafe.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) Nonce() (*big.Int, error) {
	return _GnosisSafe.Contract.Nonce(&_GnosisSafe.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCaller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GnosisSafe.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.SignedMessages(&_GnosisSafe.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_GnosisSafe *GnosisSafeCallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _GnosisSafe.Contract.SignedMessages(&_GnosisSafe.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.AddOwnerWithThreshold(&_GnosisSafe.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.AddOwnerWithThreshold(&_GnosisSafe.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ApproveHash(&_GnosisSafe.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ApproveHash(&_GnosisSafe.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ChangeThreshold(&_GnosisSafe.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ChangeThreshold(&_GnosisSafe.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.DisableModule(&_GnosisSafe.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.DisableModule(&_GnosisSafe.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.EnableModule(&_GnosisSafe.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.EnableModule(&_GnosisSafe.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransaction(&_GnosisSafe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransaction(&_GnosisSafe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModule(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModule(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_GnosisSafe *GnosisSafeTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.ExecTransactionFromModuleReturnData(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RemoveOwner(&_GnosisSafe.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RemoveOwner(&_GnosisSafe.TransactOpts, prevOwner, owner, _threshold)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeTransactor) RequiredTxGas(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "requiredTxGas", to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RequiredTxGas(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_GnosisSafe *GnosisSafeTransactorSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _GnosisSafe.Contract.RequiredTxGas(&_GnosisSafe.TransactOpts, to, value, data, operation)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetFallbackHandler(&_GnosisSafe.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetFallbackHandler(&_GnosisSafe.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetGuard(&_GnosisSafe.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SetGuard(&_GnosisSafe.TransactOpts, guard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Setup(&_GnosisSafe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Setup(&_GnosisSafe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SimulateAndRevert(&_GnosisSafe.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SimulateAndRevert(&_GnosisSafe.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SwapOwner(&_GnosisSafe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_GnosisSafe *GnosisSafeTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GnosisSafe.Contract.SwapOwner(&_GnosisSafe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Fallback(&_GnosisSafe.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafe.Contract.Fallback(&_GnosisSafe.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeSession) Receive() (*types.Transaction, error) {
	return _GnosisSafe.Contract.Receive(&_GnosisSafe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GnosisSafe *GnosisSafeTransactorSession) Receive() (*types.Transaction, error) {
	return _GnosisSafe.Contract.Receive(&_GnosisSafe.TransactOpts)
}

// GnosisSafeAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the GnosisSafe contract.
type GnosisSafeAddedOwnerIterator struct {
	Event *GnosisSafeAddedOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeAddedOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeAddedOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeAddedOwner represents a AddedOwner event raised by the GnosisSafe contract.
type GnosisSafeAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*GnosisSafeAddedOwnerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeAddedOwnerIterator{contract: _GnosisSafe.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeAddedOwner) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeAddedOwner)
				if err := _GnosisSafe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseAddedOwner(log types.Log) (*GnosisSafeAddedOwner, error) {
	event := new(GnosisSafeAddedOwner)
	if err := _GnosisSafe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the GnosisSafe contract.
type GnosisSafeApproveHashIterator struct {
	Event *GnosisSafeApproveHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeApproveHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeApproveHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeApproveHash represents a ApproveHash event raised by the GnosisSafe contract.
type GnosisSafeApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*GnosisSafeApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeApproveHashIterator{contract: _GnosisSafe.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *GnosisSafeApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeApproveHash)
				if err := _GnosisSafe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseApproveHash(log types.Log) (*GnosisSafeApproveHash, error) {
	event := new(GnosisSafeApproveHash)
	if err := _GnosisSafe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the GnosisSafe contract.
type GnosisSafeChangedFallbackHandlerIterator struct {
	Event *GnosisSafeChangedFallbackHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedFallbackHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeChangedFallbackHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedFallbackHandler represents a ChangedFallbackHandler event raised by the GnosisSafe contract.
type GnosisSafeChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts) (*GnosisSafeChangedFallbackHandlerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedFallbackHandlerIterator{contract: _GnosisSafe.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedFallbackHandler) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedFallbackHandler)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedFallbackHandler(log types.Log) (*GnosisSafeChangedFallbackHandler, error) {
	event := new(GnosisSafeChangedFallbackHandler)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the GnosisSafe contract.
type GnosisSafeChangedGuardIterator struct {
	Event *GnosisSafeChangedGuard // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedGuard)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeChangedGuard)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedGuard represents a ChangedGuard event raised by the GnosisSafe contract.
type GnosisSafeChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedGuard(opts *bind.FilterOpts) (*GnosisSafeChangedGuardIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedGuardIterator{contract: _GnosisSafe.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedGuard) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedGuard)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedGuard(log types.Log) (*GnosisSafeChangedGuard, error) {
	event := new(GnosisSafeChangedGuard)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the GnosisSafe contract.
type GnosisSafeChangedThresholdIterator struct {
	Event *GnosisSafeChangedThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeChangedThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeChangedThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeChangedThreshold represents a ChangedThreshold event raised by the GnosisSafe contract.
type GnosisSafeChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*GnosisSafeChangedThresholdIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeChangedThresholdIterator{contract: _GnosisSafe.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *GnosisSafeChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeChangedThreshold)
				if err := _GnosisSafe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_GnosisSafe *GnosisSafeFilterer) ParseChangedThreshold(log types.Log) (*GnosisSafeChangedThreshold, error) {
	event := new(GnosisSafeChangedThreshold)
	if err := _GnosisSafe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the GnosisSafe contract.
type GnosisSafeDisabledModuleIterator struct {
	Event *GnosisSafeDisabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeDisabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeDisabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeDisabledModule represents a DisabledModule event raised by the GnosisSafe contract.
type GnosisSafeDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*GnosisSafeDisabledModuleIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeDisabledModuleIterator{contract: _GnosisSafe.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeDisabledModule) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeDisabledModule)
				if err := _GnosisSafe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) ParseDisabledModule(log types.Log) (*GnosisSafeDisabledModule, error) {
	event := new(GnosisSafeDisabledModule)
	if err := _GnosisSafe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the GnosisSafe contract.
type GnosisSafeEnabledModuleIterator struct {
	Event *GnosisSafeEnabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeEnabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeEnabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeEnabledModule represents a EnabledModule event raised by the GnosisSafe contract.
type GnosisSafeEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*GnosisSafeEnabledModuleIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeEnabledModuleIterator{contract: _GnosisSafe.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *GnosisSafeEnabledModule) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeEnabledModule)
				if err := _GnosisSafe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_GnosisSafe *GnosisSafeFilterer) ParseEnabledModule(log types.Log) (*GnosisSafeEnabledModule, error) {
	event := new(GnosisSafeEnabledModule)
	if err := _GnosisSafe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the GnosisSafe contract.
type GnosisSafeExecutionFailureIterator struct {
	Event *GnosisSafeExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFailure represents a ExecutionFailure event raised by the GnosisSafe contract.
type GnosisSafeExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*GnosisSafeExecutionFailureIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFailureIterator{contract: _GnosisSafe.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFailure)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFailure(log types.Log) (*GnosisSafeExecutionFailure, error) {
	event := new(GnosisSafeExecutionFailure)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleFailureIterator struct {
	Event *GnosisSafeExecutionFromModuleFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFromModuleFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeExecutionFromModuleFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFromModuleFailureIterator{contract: _GnosisSafe.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFromModuleFailure)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFromModuleFailure(log types.Log) (*GnosisSafeExecutionFromModuleFailure, error) {
	event := new(GnosisSafeExecutionFromModuleFailure)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleSuccessIterator struct {
	Event *GnosisSafeExecutionFromModuleSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionFromModuleSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeExecutionFromModuleSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the GnosisSafe contract.
type GnosisSafeExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*GnosisSafeExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionFromModuleSuccessIterator{contract: _GnosisSafe.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionFromModuleSuccess)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*GnosisSafeExecutionFromModuleSuccess, error) {
	event := new(GnosisSafeExecutionFromModuleSuccess)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the GnosisSafe contract.
type GnosisSafeExecutionSuccessIterator struct {
	Event *GnosisSafeExecutionSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeExecutionSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeExecutionSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeExecutionSuccess represents a ExecutionSuccess event raised by the GnosisSafe contract.
type GnosisSafeExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*GnosisSafeExecutionSuccessIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeExecutionSuccessIterator{contract: _GnosisSafe.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *GnosisSafeExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeExecutionSuccess)
				if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_GnosisSafe *GnosisSafeFilterer) ParseExecutionSuccess(log types.Log) (*GnosisSafeExecutionSuccess, error) {
	event := new(GnosisSafeExecutionSuccess)
	if err := _GnosisSafe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the GnosisSafe contract.
type GnosisSafeRemovedOwnerIterator struct {
	Event *GnosisSafeRemovedOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeRemovedOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeRemovedOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeRemovedOwner represents a RemovedOwner event raised by the GnosisSafe contract.
type GnosisSafeRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*GnosisSafeRemovedOwnerIterator, error) {

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeRemovedOwnerIterator{contract: _GnosisSafe.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *GnosisSafeRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeRemovedOwner)
				if err := _GnosisSafe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_GnosisSafe *GnosisSafeFilterer) ParseRemovedOwner(log types.Log) (*GnosisSafeRemovedOwner, error) {
	event := new(GnosisSafeRemovedOwner)
	if err := _GnosisSafe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the GnosisSafe contract.
type GnosisSafeSafeReceivedIterator struct {
	Event *GnosisSafeSafeReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSafeReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeSafeReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSafeReceived represents a SafeReceived event raised by the GnosisSafe contract.
type GnosisSafeSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*GnosisSafeSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSafeReceivedIterator{contract: _GnosisSafe.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *GnosisSafeSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSafeReceived)
				if err := _GnosisSafe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_GnosisSafe *GnosisSafeFilterer) ParseSafeReceived(log types.Log) (*GnosisSafeSafeReceived, error) {
	event := new(GnosisSafeSafeReceived)
	if err := _GnosisSafe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the GnosisSafe contract.
type GnosisSafeSafeSetupIterator struct {
	Event *GnosisSafeSafeSetup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeSafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSafeSetup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeSafeSetup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeSafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSafeSetup represents a SafeSetup event raised by the GnosisSafe contract.
type GnosisSafeSafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*GnosisSafeSafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSafeSetupIterator{contract: _GnosisSafe.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *GnosisSafeSafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSafeSetup)
				if err := _GnosisSafe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_GnosisSafe *GnosisSafeFilterer) ParseSafeSetup(log types.Log) (*GnosisSafeSafeSetup, error) {
	event := new(GnosisSafeSafeSetup)
	if err := _GnosisSafe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GnosisSafeSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the GnosisSafe contract.
type GnosisSafeSignMsgIterator struct {
	Event *GnosisSafeSignMsg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GnosisSafeSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeSignMsg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GnosisSafeSignMsg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GnosisSafeSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeSignMsg represents a SignMsg event raised by the GnosisSafe contract.
type GnosisSafeSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*GnosisSafeSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafe.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeSignMsgIterator{contract: _GnosisSafe.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *GnosisSafeSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _GnosisSafe.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeSignMsg)
				if err := _GnosisSafe.contract.UnpackLog(event, "SignMsg", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_GnosisSafe *GnosisSafeFilterer) ParseSignMsg(log types.Log) (*GnosisSafeSignMsg, error) {
	event := new(GnosisSafeSignMsg)
	if err := _GnosisSafe.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
