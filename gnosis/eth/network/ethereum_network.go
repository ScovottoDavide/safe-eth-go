package network

import (
	"github.com/ethereum/go-ethereum/common"
)

// NETWORK
var (
	Unknown         = Network{-1}
	Olympic         = Network{0}
	Mainnet         = Network{1}
	Gochain_Testnet = Network{31337}
)

var networks = map[int]Network{
	Unknown.ChainId:         Unknown,
	Olympic.ChainId:         Olympic,
	Mainnet.ChainId:         Mainnet,
	Gochain_Testnet.ChainId: Gochain_Testnet,
}

func IsSupported(chain_id int) bool {
	_, exists := networks[chain_id]
	return exists
}

func GetNetwork(chainId int) Network {
	return networks[chainId]
}

// TRANSACTION SPEED
var (
	Slowest  = TxSpeed{0}
	VerySlow = TxSpeed{10}
	Slow     = TxSpeed{25}
	Normal   = TxSpeed{50}
	Fast     = TxSpeed{75}
	VeryFast = TxSpeed{90}
	Fastest  = TxSpeed{100}
)

var TX_SPEED_MAP = map[int]TxSpeed{
	Slowest.Speed:  Slowest,
	VerySlow.Speed: VerySlow,
	Slow.Speed:     Slow,
	Normal.Speed:   Normal,
	Fast.Speed:     Fast,
	VeryFast.Speed: VeryFast,
	Fastest.Speed:  Fastest,
}

func IsTxSpeedAvailable(txSpeed TxSpeed) bool {
	_, exists := TX_SPEED_MAP[txSpeed.Speed]
	return exists
}

var NetworkToSafeAddress = map[Network]GnosisSafeContract{
	Mainnet: {
		Address: common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"),
	},
	Gochain_Testnet: {
		Address: common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"),
	},
}

var NetworkToSafeProxyFactoryAddress = map[Network]GnosisSafeContract{
	Mainnet: {
		Address: common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
	},
	Gochain_Testnet: {
		Address: common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
	},
}
