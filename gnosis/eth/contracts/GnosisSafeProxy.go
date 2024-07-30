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

// GnosisSafeProxyMetaData contains all meta data concerning the GnosisSafeProxy contract.
var GnosisSafeProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
}

// GnosisSafeProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use GnosisSafeProxyMetaData.ABI instead.
var GnosisSafeProxyABI = GnosisSafeProxyMetaData.ABI

// GnosisSafeProxy is an auto generated Go binding around an Ethereum contract.
type GnosisSafeProxy struct {
	GnosisSafeProxyCaller     // Read-only binding to the contract
	GnosisSafeProxyTransactor // Write-only binding to the contract
	GnosisSafeProxyFilterer   // Log filterer for contract events
}

// GnosisSafeProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeProxySession struct {
	Contract     *GnosisSafeProxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GnosisSafeProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeProxyCallerSession struct {
	Contract *GnosisSafeProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GnosisSafeProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeProxyTransactorSession struct {
	Contract     *GnosisSafeProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GnosisSafeProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeProxyRaw struct {
	Contract *GnosisSafeProxy // Generic contract binding to access the raw methods on
}

// GnosisSafeProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeProxyCallerRaw struct {
	Contract *GnosisSafeProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeProxyTransactorRaw struct {
	Contract *GnosisSafeProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeProxy creates a new instance of GnosisSafeProxy, bound to a specific deployed contract.
func NewGnosisSafeProxy(address common.Address, backend bind.ContractBackend) (*GnosisSafeProxy, error) {
	contract, err := bindGnosisSafeProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxy{GnosisSafeProxyCaller: GnosisSafeProxyCaller{contract: contract}, GnosisSafeProxyTransactor: GnosisSafeProxyTransactor{contract: contract}, GnosisSafeProxyFilterer: GnosisSafeProxyFilterer{contract: contract}}, nil
}

// NewGnosisSafeProxyCaller creates a new read-only instance of GnosisSafeProxy, bound to a specific deployed contract.
func NewGnosisSafeProxyCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeProxyCaller, error) {
	contract, err := bindGnosisSafeProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyCaller{contract: contract}, nil
}

// NewGnosisSafeProxyTransactor creates a new write-only instance of GnosisSafeProxy, bound to a specific deployed contract.
func NewGnosisSafeProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeProxyTransactor, error) {
	contract, err := bindGnosisSafeProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyTransactor{contract: contract}, nil
}

// NewGnosisSafeProxyFilterer creates a new log filterer instance of GnosisSafeProxy, bound to a specific deployed contract.
func NewGnosisSafeProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeProxyFilterer, error) {
	contract, err := bindGnosisSafeProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFilterer{contract: contract}, nil
}

// bindGnosisSafeProxy binds a generic wrapper to an already deployed contract.
func bindGnosisSafeProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GnosisSafeProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxy *GnosisSafeProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxy.Contract.GnosisSafeProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxy *GnosisSafeProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.GnosisSafeProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxy *GnosisSafeProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.GnosisSafeProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxy *GnosisSafeProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxy *GnosisSafeProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxy *GnosisSafeProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GnosisSafeProxy *GnosisSafeProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GnosisSafeProxy *GnosisSafeProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.Fallback(&_GnosisSafeProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GnosisSafeProxy *GnosisSafeProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GnosisSafeProxy.Contract.Fallback(&_GnosisSafeProxy.TransactOpts, calldata)
}
