package eth

// NETWORK
var (
	Unknown         = Network{-1}
	Olympic         = Network{0}
	Mainnet         = Network{1}
	Gochain_Testnet = Network{31337}
)

var networks = map[int]Network{
	Unknown.chainId:         Unknown,
	Olympic.chainId:         Olympic,
	Mainnet.chainId:         Mainnet,
	Gochain_Testnet.chainId: Gochain_Testnet,
}

func isSupported(chain_id int) bool {
	_, exists := networks[chain_id]
	return exists
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

var tx_speed_map = map[int]TxSpeed{
	Slowest.speed:  Slowest,
	VerySlow.speed: VerySlow,
	Slow.speed:     Slow,
	Normal.speed:   Normal,
	Fast.speed:     Fast,
	VeryFast.speed: VeryFast,
	Fastest.speed:  Fastest,
}

func isTxSpeedAvailable(txSpeed TxSpeed) bool {
	_, exists := tx_speed_map[txSpeed.speed]
	return exists
}
