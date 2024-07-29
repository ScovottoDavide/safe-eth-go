package safe

import (
	"fmt"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// keccak256("fallback_manager.handler.address")
const FALLBACK_HANDLER_STORAGE_SLOT = 0x6C9A6C4A39284E37ED1CF53D337577D14212A4870FB976A4366C693B939918D5

// keccak256("guard_manager.guard.address")
const GUARD_STORAGE_SLOT = 0x4A204F620C8C5CCDCA3FD54D003BADD85BA500436A431F0CBDA4F558C93C34C8

// keccak256("SafeMessage(bytes message)");
var SAFE_MESSAGE_TYPEHASH = common.Hex2Bytes("60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca")

type Safe struct {
	ethereumClient *eth.EthereumClient
	address        *common.Address
	safeContract   *contracts.Contract
}

func (safe *Safe) String() string {
	return "Safe=" + safe.address.Hex()
}

func New(ethClient *eth.EthereumClient) *Safe {
	chainId, err := ethClient.GetChainId()
	if err != nil {
		return nil
	}
	safeAddress := network.NetworkToSafeAddress[network.GetNetwork(chainId)].Address
	return &Safe{
		ethereumClient: ethClient,
		address:        &safeAddress,
		safeContract:   getSafeContract(ethClient.GetGEthClient()),
	}
}

func getSafeContract(backend *ethclient.Client) *contracts.Contract {
	contract, err := contracts.GetSafeV130Contract(backend)
	if err != nil {
		return nil
	}
	return contract
}

func (safe *Safe) Version() (string, error) {
	if safe.safeContract == nil {
		return "", fmt.Errorf("no GnosisSafe contract instance")
	}
	var out []interface{}
	err := safe.safeContract.Caller.Contract.Call(nil, &out, "retrieve")

	if err != nil {
		return "", err
	}

	out0 := abi.ConvertType(out[0], new(string)).(string)

	return out0, err
}
