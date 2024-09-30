package testcommon

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/network"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/safe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const HARDHAT_S_KEY0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const owner2 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const owner3 = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

var uri = eth.NewURI("http://127.0.0.1:8545")
var EthClient, _ = eth.EthereumClientInit(uri)
var chainId, _ = EthClient.GetChainId()
var PrivateKey, _ = eth.GetCryptoPrivateKey(HARDHAT_S_KEY0)
var Sender, _ = eth.AddressFromPrivKey(hexutil.MustDecode(HARDHAT_S_KEY0))

func CreateTestSafe() *common.Address {
	return deploySafe(Sender, EthClient, PrivateKey)
}

func deploySafe(sender *common.Address, ethClient *eth.EthereumClient, privateKey *ecdsa.PrivateKey) *common.Address {
	var owners []common.Address
	owners = append(owners, *sender)
	owners = append(owners, common.HexToAddress(owner2))
	owners = append(owners, common.HexToAddress(owner3))
	txSent, err := safe.Create(
		ethClient,
		*sender,
		privateKey,
		network.NetworkToMasterCopyAddress[network.GetNetwork(chainId)].Address,
		owners,
		2,
		eth.NULL_ADDRESS,
		common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
		eth.NULL_ADDRESS,
		0,
		eth.NULL_ADDRESS,
	)
	if err != nil {
		return nil
	}
	if txSent.ContractAddress == eth.NULL_ADDRESS {
		return nil
	}
	ch := ethClient.WaitTxConfirmed(txSent.TxHash)
	isPending := <-ch
	if isPending {
		fmt.Printf("unexpected pending tx %s", txSent.TxHash.Hex())
	}
	receipt, err := ethClient.GetReceipt(txSent.TxHash.Hex())
	if err != nil {
		return nil
	}
	fmt.Println("Used gas: ", receipt.GasUsed)
	// fmt.Println("Safe creation payment: ", receipt.GasUsed*receipt.EffectiveGasPrice.Uint64())
	fmt.Println(txSent.ContractAddress.Hex())
	return &txSent.ContractAddress
}
