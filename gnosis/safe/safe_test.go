package safe

import (
	"fmt"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
)

func TestSafeCall(t *testing.T) {
	uri := eth.NewURI("http://127.0.0.1:8545")
	ethClient, err := eth.EthereumClientInit(uri)
	if err != nil {
		panic(err)
	}
	safe := New(ethClient)
	if safe == nil {
		t.Fatalf("Got nil Safe object")
	}

	version, err := safe.Version()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(version)
}
