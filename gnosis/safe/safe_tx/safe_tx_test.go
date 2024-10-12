package safetx_test

import (
	"log"
	"math/big"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/safe"
	safetx "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/safe_tx"
	testcommon "github.com/ScovottoDavide/safe-eth-go/gnosis/safe/test_common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var safe_ *safe.Safe

func setupTest(_ *testing.T) func(t *testing.T) {
	log.Println("setup test")

	ethClient := testcommon.EthClient
	safeSigner := testcommon.PrivateKey
	testSafe := testcommon.CreateTestSafe()

	safe_ = safe.New(*testSafe, ethClient, safeSigner)

	return func(t *testing.T) {
		log.Println("teardown test")
	}
}

func TestSafeTxHash(t *testing.T) {
	tearDownTest := setupTest(t)
	defer tearDownTest(t)

	// Expected hash must be the same calculated by `getTransactionHash` of the contract
	expectedSafeTxHash, err := safe_.SafeContract.GetTransactionHash(
		&bind.CallOpts{
			From: *testcommon.Sender,
		},
		common.HexToAddress("0x5AC255889882aaB35A2aa939679E3F3d4Cea221E"),
		big.NewInt(5212459),
		make([]byte, 0),
		1,
		big.NewInt(123456),
		big.NewInt(122),
		big.NewInt(12345),
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		big.NewInt(10789),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	safeTx := safetx.New(
		safe_.EthereumClient,
		*safe_.SafeAddress,
		common.HexToAddress("0x5AC255889882aaB35A2aa939679E3F3d4Cea221E"),
		big.NewInt(5212459),
		make([]byte, 0),
		1,
		big.NewInt(123456),
		big.NewInt(122),
		big.NewInt(12345),
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		big.NewInt(10789),
		nil,
		"1.3.0",
	)
	safeTxHash, err := safeTx.SafeTxHash()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if safeTxHash.Cmp(common.BytesToHash(expectedSafeTxHash[:])) != 0 {
		t.Fatalf("inconsistent safe tx hashes -> contract hash (%s), SafeTx hash (%s)",
			common.BytesToHash(expectedSafeTxHash[:]).String(), safeTxHash.String())
	}
}

func TestSignaturesAndSigner() {

}
