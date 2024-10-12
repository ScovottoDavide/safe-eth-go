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
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func TestSignaturesAndSigner(t *testing.T) {
	tearDownTest := setupTest(t)
	defer tearDownTest(t)

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

	owner0_sk, _ := eth.GetCryptoPrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	owner1_sk, _ := eth.GetCryptoPrivateKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	// TODO: check expected signers order is as expected
	// owner0_addr, _ := eth.AddressFromPrivKey(owner0_sk.D.Bytes())
	// owner1_addr, _ := eth.AddressFromPrivKey(owner1_sk.D.Bytes())

	owner0_sig, err := safeTx.Sign(owner0_sk)
	if err != nil {
		t.Error(err)
	}
	owner1_sig, err := safeTx.Sign(owner1_sk)
	if err != nil {
		t.Error(err)
	}

	t.Log("owner0 sig ", hexutil.Encode(owner0_sig))
	t.Log("owner1 sig ", hexutil.Encode(owner1_sig))

	t.Log("SafeTx.Signatures len ", len(safeTx.Signatures))
	t.Log("SafeTx.Signatures: ", hexutil.Encode(safeTx.Signatures))
	t.Log("SafeTx.Signers: ", safeTx.Signers)

	if len(safeTx.Signatures) != 65*2 {
		t.Errorf("we should have 2 signatures (130 bytes)")
	}
	if len(safeTx.Signers) != 2 {
		t.Errorf("signers array len should be 2")
	}
}
