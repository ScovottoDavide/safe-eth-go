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

// GnosisSafeProxyFactoryMetaData contains all meta data concerning the GnosisSafeProxyFactory contract.
var GnosisSafeProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"ProxyCreation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"calculateCreateProxyWithNonceAddress\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createProxy\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIProxyCreationCallback\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"createProxyWithCallback\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyRuntimeCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// GnosisSafeProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeProxyFactoryMetaData.ABI instead.
var GnosisSafeProxyFactoryABI = GnosisSafeProxyFactoryMetaData.ABI

// GnosisSafeProxyFactory is an auto generated Go binding around an Ethereum contract.
type GnosisSafeProxyFactory struct {
	GnosisSafeProxyFactoryCaller     // Read-only binding to the contract
	GnosisSafeProxyFactoryTransactor // Write-only binding to the contract
	GnosisSafeProxyFactoryFilterer   // Log filterer for contract events
}

// GnosisSafeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeProxyFactorySession struct {
	Contract     *GnosisSafeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GnosisSafeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeProxyFactoryCallerSession struct {
	Contract *GnosisSafeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GnosisSafeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeProxyFactoryTransactorSession struct {
	Contract     *GnosisSafeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GnosisSafeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryRaw struct {
	Contract *GnosisSafeProxyFactory // Generic contract binding to access the raw methods on
}

// GnosisSafeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryCallerRaw struct {
	Contract *GnosisSafeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryTransactorRaw struct {
	Contract *GnosisSafeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeProxyFactory creates a new instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactory(address common.Address, backend bind.ContractBackend) (*GnosisSafeProxyFactory, error) {
	contract, err := bindGnosisSafeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactory{GnosisSafeProxyFactoryCaller: GnosisSafeProxyFactoryCaller{contract: contract}, GnosisSafeProxyFactoryTransactor: GnosisSafeProxyFactoryTransactor{contract: contract}, GnosisSafeProxyFactoryFilterer: GnosisSafeProxyFactoryFilterer{contract: contract}}, nil
}

// NewGnosisSafeProxyFactoryCaller creates a new read-only instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeProxyFactoryCaller, error) {
	contract, err := bindGnosisSafeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryCaller{contract: contract}, nil
}

// NewGnosisSafeProxyFactoryTransactor creates a new write-only instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeProxyFactoryTransactor, error) {
	contract, err := bindGnosisSafeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryTransactor{contract: contract}, nil
}

// NewGnosisSafeProxyFactoryFilterer creates a new log filterer instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeProxyFactoryFilterer, error) {
	contract, err := bindGnosisSafeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryFilterer{contract: contract}, nil
}

// bindGnosisSafeProxyFactory binds a generic wrapper to an already deployed contract.
func bindGnosisSafeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GnosisSafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafeProxyFactory.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) ProxyCreationCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyCreationCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerSession) ProxyCreationCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyCreationCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCaller) ProxyRuntimeCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _GnosisSafeProxyFactory.contract.Call(opts, &out, "proxyRuntimeCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) ProxyRuntimeCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyRuntimeCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerSession) ProxyRuntimeCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyRuntimeCode(&_GnosisSafeProxyFactory.CallOpts)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CalculateCreateProxyWithNonceAddress(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "calculateCreateProxyWithNonceAddress", _singleton, initializer, saltNonce)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CalculateCreateProxyWithNonceAddress(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CalculateCreateProxyWithNonceAddress(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CalculateCreateProxyWithNonceAddress(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CalculateCreateProxyWithNonceAddress(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address singleton, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, singleton common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxy", singleton, data)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address singleton, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxy(singleton common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxy(&_GnosisSafeProxyFactory.TransactOpts, singleton, data)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address singleton, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxy(singleton common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxy(&_GnosisSafeProxyFactory.TransactOpts, singleton, data)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxyWithCallback(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxyWithCallback", _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithCallback(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithCallback(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithNonce(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithNonce(&_GnosisSafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// GnosisSafeProxyFactoryProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the GnosisSafeProxyFactory contract.
type GnosisSafeProxyFactoryProxyCreationIterator struct {
	Event *GnosisSafeProxyFactoryProxyCreation // Event containing the contract specifics and raw log

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
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeProxyFactoryProxyCreation)
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
		it.Event = new(GnosisSafeProxyFactoryProxyCreation)
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
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeProxyFactoryProxyCreation represents a ProxyCreation event raised by the GnosisSafeProxyFactory contract.
type GnosisSafeProxyFactoryProxyCreation struct {
	Proxy     common.Address
	Singleton common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address singleton)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) FilterProxyCreation(opts *bind.FilterOpts) (*GnosisSafeProxyFactoryProxyCreationIterator, error) {

	logs, sub, err := _GnosisSafeProxyFactory.contract.FilterLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryProxyCreationIterator{contract: _GnosisSafeProxyFactory.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address singleton)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *GnosisSafeProxyFactoryProxyCreation) (event.Subscription, error) {

	logs, sub, err := _GnosisSafeProxyFactory.contract.WatchLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeProxyFactoryProxyCreation)
				if err := _GnosisSafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address proxy, address singleton)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) ParseProxyCreation(log types.Log) (*GnosisSafeProxyFactoryProxyCreation, error) {
	event := new(GnosisSafeProxyFactoryProxyCreation)
	if err := _GnosisSafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
