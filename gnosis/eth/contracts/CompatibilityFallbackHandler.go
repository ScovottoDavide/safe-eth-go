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

// CompatibiliyFallbackHandlerMetaData contains all meta data concerning the CompatibiliyFallbackHandler contract.
var CompatibiliyFallbackHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHashForSafe\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610eca806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ce5760003560e01c80636ac247841161008c578063bc197c8111610066578063bc197c8114610205578063bd61951d14610227578063f23a6e611461023a578063ffa1ad741461025a57600080fd5b80636ac2478414610194578063a3f4df7e146101a7578063b2494df3146101f057600080fd5b806223de29146100d357806301ffc9a7146100ed5780630a1028c414610115578063150b7a02146101365780631626ba7e1461016e57806320c13b0b14610181575b600080fd5b6100eb6100e136600461082e565b5050505050505050565b005b6101006100fb366004610bbc565b61027e565b60405190151581526020015b60405180910390f35b610128610123366004610c5c565b6102d0565b60405190815260200161010c565b610155610144366004610974565b630a85bd0160e11b95945050505050565b6040516001600160e01b0319909116815260200161010c565b61015561017c366004610b8c565b6102dc565b61015561018f366004610bf4565b6103a9565b6101286101a2366004610c96565b610531565b6101e36040518060400160405280601881526020017f44656661756c742043616c6c6261636b2048616e646c6572000000000000000081525081565b60405161010c9190610ddd565b6101f8610658565b60405161010c9190610d57565b6101556102133660046108db565b63bc197c8160e01b98975050505050505050565b6101e3610235366004610a5d565b6106e2565b6101556102483660046109e4565b63f23a6e6160e01b9695505050505050565b6101e3604051806040016040528060058152602001640312e302e360dc1b81525081565b60006001600160e01b03198216630271189760e51b14806102af57506001600160e01b03198216630a85bd0160e11b145b806102ca57506001600160e01b031982166301ffc9a760e01b145b92915050565b60006102ca3383610531565b60408051602080820186905282518083039091018152818301928390526320c13b0b60e01b9092526000913391839183916320c13b0b916103239189908990604401610df0565b60206040518083038186803b15801561033b57600080fd5b505afa15801561034f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103739190610bd8565b90506001600160e01b031981166320c13b0b60e01b1461039457600061039d565b630b135d3f60e11b5b925050505b9392505050565b60008033905060006103f18288888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061053192505050565b9050836104b957604051635ae6bd3760e01b8152600481018290526001600160a01b03831690635ae6bd379060240160206040518083038186803b15801561043857600080fd5b505afa15801561044c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104709190610b74565b6104b45760405162461bcd60e51b815260206004820152601160248201527012185cda081b9bdd08185c1c1c9bdd9959607a1b604482015260640160405180910390fd5b61051e565b60405163934f3a1160e01b81526001600160a01b0383169063934f3a11906104ed9084908b908b908b908b90600401610da4565b60006040518083038186803b15801561050557600080fd5b505afa158015610519573d6000803e3d6000fd5b505050505b506320c13b0b60e01b9695505050505050565b6000807f60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca60001b838051906020012060405160200161057a929190918252602082015260400190565b604051602081830303815290604052805190602001209050601960f81b600160f81b856001600160a01b031663f698da256040518163ffffffff1660e01b815260040160206040518083038186803b1580156105d557600080fd5b505afa1580156105e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060d9190610b74565b6040516001600160f81b031993841660208201529290911660218301526022820152604281018290526062016040516020818303038152906040528051906020012091505092915050565b604051636617c22960e11b815260016004820152600a60248201526060903390600090829063cc2f84529060440160006040518083038186803b15801561069e57600080fd5b505afa1580156106b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106da9190810190610aaf565b509392505050565b606060405163b4faba0960e01b8152600436036004808301376020600036836000335af1505060203d036040519150808201604052806020833e506000516103a257805160208201fd5b805161073781610e66565b919050565b60008083601f84011261074d578182fd5b5081356001600160401b03811115610763578182fd5b6020830191508360208260051b850101111561077e57600080fd5b9250929050565b60008083601f840112610796578182fd5b5081356001600160401b038111156107ac578182fd5b60208301915083602082850101111561077e57600080fd5b600082601f8301126107d4578081fd5b81356001600160401b038111156107ed576107ed610e50565b610800601f8201601f1916602001610e20565b818152846020838601011115610814578283fd5b816020850160208301379081016020019190915292915050565b60008060008060008060008060c0898b031215610849578384fd5b883561085481610e66565b9750602089013561086481610e66565b9650604089013561087481610e66565b95506060890135945060808901356001600160401b0380821115610896578586fd5b6108a28c838d01610785565b909650945060a08b01359150808211156108ba578384fd5b506108c78b828c01610785565b999c989b5096995094979396929594505050565b60008060008060008060008060a0898b0312156108f6578384fd5b883561090181610e66565b9750602089013561091181610e66565b965060408901356001600160401b038082111561092c578586fd5b6109388c838d0161073c565b909850965060608b0135915080821115610950578586fd5b61095c8c838d0161073c565b909650945060808b01359150808211156108ba578384fd5b60008060008060006080868803121561098b578081fd5b853561099681610e66565b945060208601356109a681610e66565b93506040860135925060608601356001600160401b038111156109c7578182fd5b6109d388828901610785565b969995985093965092949392505050565b60008060008060008060a087890312156109fc578182fd5b8635610a0781610e66565b95506020870135610a1781610e66565b9450604087013593506060870135925060808701356001600160401b03811115610a3f578283fd5b610a4b89828a01610785565b979a9699509497509295939492505050565b600080600060408486031215610a71578283fd5b8335610a7c81610e66565b925060208401356001600160401b03811115610a96578283fd5b610aa286828701610785565b9497909650939450505050565b60008060408385031215610ac1578182fd5b82516001600160401b0380821115610ad7578384fd5b818501915085601f830112610aea578384fd5b8151602082821115610afe57610afe610e50565b8160051b9250610b0f818401610e20565b8281528181019085830185870184018b1015610b29578889fd5b8896505b84871015610b575780519550610b4286610e66565b85835260019690960195918301918301610b2d565b509650610b67905087820161072c565b9450505050509250929050565b600060208284031215610b85578081fd5b5051919050565b600080600060408486031215610ba0578081fd5b8335925060208401356001600160401b03811115610a96578182fd5b600060208284031215610bcd578081fd5b81356103a281610e7e565b600060208284031215610be9578081fd5b81516103a281610e7e565b60008060008060408587031215610c09578182fd5b84356001600160401b0380821115610c1f578384fd5b610c2b88838901610785565b90965094506020870135915080821115610c43578384fd5b50610c5087828801610785565b95989497509550505050565b600060208284031215610c6d578081fd5b81356001600160401b03811115610c82578182fd5b610c8e848285016107c4565b949350505050565b60008060408385031215610ca8578182fd5b8235610cb381610e66565b915060208301356001600160401b03811115610ccd578182fd5b610cd9858286016107c4565b9150509250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008151808452815b81811015610d3157602081850181015186830182015201610d15565b81811115610d425782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015610d985783516001600160a01b031683529284019291840191600101610d73565b50909695505050505050565b858152606060208201526000610dbe606083018688610ce3565b8281036040840152610dd1818587610ce3565b98975050505050505050565b6020815260006103a26020830184610d0c565b604081526000610e036040830186610d0c565b8281036020840152610e16818587610ce3565b9695505050505050565b604051601f8201601f191681016001600160401b0381118282101715610e4857610e48610e50565b604052919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610e7b57600080fd5b50565b6001600160e01b031981168114610e7b57600080fdfea2646970667358221220b18ba8b243df06036984dd39603087e6075205969b74c35b223de1615425565a64736f6c63430008040033",
}

// CompatibiliyFallbackHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use CompatibiliyFallbackHandlerMetaData.ABI instead.
var CompatibiliyFallbackHandlerABI = CompatibiliyFallbackHandlerMetaData.ABI

// CompatibiliyFallbackHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompatibiliyFallbackHandlerMetaData.Bin instead.
var CompatibiliyFallbackHandlerBin = CompatibiliyFallbackHandlerMetaData.Bin

// DeployCompatibiliyFallbackHandler deploys a new Ethereum contract, binding an instance of CompatibiliyFallbackHandler to it.
func DeployCompatibiliyFallbackHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompatibiliyFallbackHandler, error) {
	parsed, err := CompatibiliyFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompatibiliyFallbackHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompatibiliyFallbackHandler{CompatibiliyFallbackHandlerCaller: CompatibiliyFallbackHandlerCaller{contract: contract}, CompatibiliyFallbackHandlerTransactor: CompatibiliyFallbackHandlerTransactor{contract: contract}, CompatibiliyFallbackHandlerFilterer: CompatibiliyFallbackHandlerFilterer{contract: contract}}, nil
}

// CompatibiliyFallbackHandler is an auto generated Go binding around an Ethereum contract.
type CompatibiliyFallbackHandler struct {
	CompatibiliyFallbackHandlerCaller     // Read-only binding to the contract
	CompatibiliyFallbackHandlerTransactor // Write-only binding to the contract
	CompatibiliyFallbackHandlerFilterer   // Log filterer for contract events
}

// CompatibiliyFallbackHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompatibiliyFallbackHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibiliyFallbackHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompatibiliyFallbackHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibiliyFallbackHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompatibiliyFallbackHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibiliyFallbackHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompatibiliyFallbackHandlerSession struct {
	Contract     *CompatibiliyFallbackHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CompatibiliyFallbackHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompatibiliyFallbackHandlerCallerSession struct {
	Contract *CompatibiliyFallbackHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// CompatibiliyFallbackHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompatibiliyFallbackHandlerTransactorSession struct {
	Contract     *CompatibiliyFallbackHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// CompatibiliyFallbackHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompatibiliyFallbackHandlerRaw struct {
	Contract *CompatibiliyFallbackHandler // Generic contract binding to access the raw methods on
}

// CompatibiliyFallbackHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompatibiliyFallbackHandlerCallerRaw struct {
	Contract *CompatibiliyFallbackHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// CompatibiliyFallbackHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompatibiliyFallbackHandlerTransactorRaw struct {
	Contract *CompatibiliyFallbackHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompatibiliyFallbackHandler creates a new instance of CompatibiliyFallbackHandler, bound to a specific deployed contract.
func NewCompatibiliyFallbackHandler(address common.Address, backend bind.ContractBackend) (*CompatibiliyFallbackHandler, error) {
	contract, err := bindCompatibiliyFallbackHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompatibiliyFallbackHandler{CompatibiliyFallbackHandlerCaller: CompatibiliyFallbackHandlerCaller{contract: contract}, CompatibiliyFallbackHandlerTransactor: CompatibiliyFallbackHandlerTransactor{contract: contract}, CompatibiliyFallbackHandlerFilterer: CompatibiliyFallbackHandlerFilterer{contract: contract}}, nil
}

// NewCompatibiliyFallbackHandlerCaller creates a new read-only instance of CompatibiliyFallbackHandler, bound to a specific deployed contract.
func NewCompatibiliyFallbackHandlerCaller(address common.Address, caller bind.ContractCaller) (*CompatibiliyFallbackHandlerCaller, error) {
	contract, err := bindCompatibiliyFallbackHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibiliyFallbackHandlerCaller{contract: contract}, nil
}

// NewCompatibiliyFallbackHandlerTransactor creates a new write-only instance of CompatibiliyFallbackHandler, bound to a specific deployed contract.
func NewCompatibiliyFallbackHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*CompatibiliyFallbackHandlerTransactor, error) {
	contract, err := bindCompatibiliyFallbackHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibiliyFallbackHandlerTransactor{contract: contract}, nil
}

// NewCompatibiliyFallbackHandlerFilterer creates a new log filterer instance of CompatibiliyFallbackHandler, bound to a specific deployed contract.
func NewCompatibiliyFallbackHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*CompatibiliyFallbackHandlerFilterer, error) {
	contract, err := bindCompatibiliyFallbackHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompatibiliyFallbackHandlerFilterer{contract: contract}, nil
}

// bindCompatibiliyFallbackHandler binds a generic wrapper to an already deployed contract.
func bindCompatibiliyFallbackHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompatibiliyFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompatibiliyFallbackHandler.Contract.CompatibiliyFallbackHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.CompatibiliyFallbackHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.CompatibiliyFallbackHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompatibiliyFallbackHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) NAME() (string, error) {
	return _CompatibiliyFallbackHandler.Contract.NAME(&_CompatibiliyFallbackHandler.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) NAME() (string, error) {
	return _CompatibiliyFallbackHandler.Contract.NAME(&_CompatibiliyFallbackHandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) VERSION() (string, error) {
	return _CompatibiliyFallbackHandler.Contract.VERSION(&_CompatibiliyFallbackHandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) VERSION() (string, error) {
	return _CompatibiliyFallbackHandler.Contract.VERSION(&_CompatibiliyFallbackHandler.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) GetMessageHash(opts *bind.CallOpts, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "getMessageHash", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.GetMessageHash(&_CompatibiliyFallbackHandler.CallOpts, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.GetMessageHash(&_CompatibiliyFallbackHandler.CallOpts, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) GetMessageHashForSafe(opts *bind.CallOpts, safe common.Address, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "getMessageHashForSafe", safe, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.GetMessageHashForSafe(&_CompatibiliyFallbackHandler.CallOpts, safe, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.GetMessageHashForSafe(&_CompatibiliyFallbackHandler.CallOpts, safe, message)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) GetModules() ([]common.Address, error) {
	return _CompatibiliyFallbackHandler.Contract.GetModules(&_CompatibiliyFallbackHandler.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) GetModules() ([]common.Address, error) {
	return _CompatibiliyFallbackHandler.Contract.GetModules(&_CompatibiliyFallbackHandler.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "isValidSignature", _dataHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.IsValidSignature(&_CompatibiliyFallbackHandler.CallOpts, _dataHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.IsValidSignature(&_CompatibiliyFallbackHandler.CallOpts, _dataHash, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "isValidSignature0", _data, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.IsValidSignature0(&_CompatibiliyFallbackHandler.CallOpts, _data, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.IsValidSignature0(&_CompatibiliyFallbackHandler.CallOpts, _data, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC1155BatchReceived(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC1155BatchReceived(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC1155Received(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC1155Received(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC721Received(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CompatibiliyFallbackHandler.Contract.OnERC721Received(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CompatibiliyFallbackHandler.Contract.SupportsInterface(&_CompatibiliyFallbackHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CompatibiliyFallbackHandler.Contract.SupportsInterface(&_CompatibiliyFallbackHandler.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _CompatibiliyFallbackHandler.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _CompatibiliyFallbackHandler.Contract.TokensReceived(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _CompatibiliyFallbackHandler.Contract.TokensReceived(&_CompatibiliyFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerTransactor) Simulate(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.contract.Transact(opts, "simulate", targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.Simulate(&_CompatibiliyFallbackHandler.TransactOpts, targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibiliyFallbackHandler *CompatibiliyFallbackHandlerTransactorSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibiliyFallbackHandler.Contract.Simulate(&_CompatibiliyFallbackHandler.TransactOpts, targetContract, calldataPayload)
}
