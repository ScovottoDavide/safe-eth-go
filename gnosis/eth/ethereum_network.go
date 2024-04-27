package eth

type Network struct {
	chainId int
}

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
