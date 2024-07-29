package network

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetAbi(t *testing.T) {
	safev130 := GnosisSafeV130{
		Address: common.Address{1},
		Path:    "./files/GnosisSafe_V1_3_0.abi",
	}
	_, err := safev130.GetAbi()
	if err != nil {
		t.Fatalf(err.Error())
	}
}
