package contracts

import (
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Caller struct {
	Contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type Transactor struct {
	Contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type Filterer struct {
	Contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

type Contract struct {
	Caller     // Read-only binding to the contract
	Transactor // Write-only binding to the contract
	Filterer   // Log filterer for contract events
}

func GetSafeV130Contract(backend bind.ContractBackend) (*Contract, error) {
	safeV130 := network.NetworkToSafeAddress[network.Gochain_Testnet]
	contractAbi, err := safeV130.GetAbi()
	if err != nil {
		return nil, err
	}
	contract, err := bindContract(safeV130.Address, backend, backend, backend, contractAbi)
	if err != nil {
		return nil, err
	}
	return &Contract{
		Caller:     Caller{Contract: contract},
		Transactor: Transactor{Contract: contract},
		Filterer:   Filterer{Contract: contract},
	}, nil
}

func bindContract(
	address common.Address,
	caller bind.ContractCaller,
	transactor bind.ContractTransactor,
	filterer bind.ContractFilterer,
	contractAbi *abi.ABI,
) (*bind.BoundContract, error) {
	return bind.NewBoundContract(address, *contractAbi, caller, transactor, filterer), nil
}
