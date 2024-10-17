package safetx_test

import (
	"log"
	"math/big"
	"slices"
	"testing"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
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

	testSafeAddr, ethClient, sKey := testcommon.CreateTestSafe()

	safe_ = safe.New(*testSafeAddr, ethClient, sKey)

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
			From: *testcommon.Sender(),
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

	safeTx := newDefaultSafeTx()

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

	safeTx := newDefaultSafeTx()

	owner0_sk, _ := eth.GetCryptoPrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	owner1_sk, _ := eth.GetCryptoPrivateKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	owner0_addr, _ := eth.AddressFromPrivKey(owner0_sk.D.Bytes())
	owner1_addr, _ := eth.AddressFromPrivKey(owner1_sk.D.Bytes())
	expectedSignersOrder := []common.Address{*owner1_addr, *owner0_addr}

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

	t.Log("SafeTx.Signatures: ", hexutil.Encode(safeTx.Signatures))
	t.Log("SafeTx.Signers: ", safeTx.Signers)
	t.Log("Expected Signers order: ", expectedSignersOrder)

	if len(safeTx.Signatures) != 65*2 {
		t.Errorf("we should have 2 signatures (130 bytes)")
	}
	if len(safeTx.Signers) != 2 {
		t.Errorf("signers array len should be 2")
	}
	if !slices.Equal(expectedSignersOrder, safeTx.Signers) {
		t.Errorf("signers are not ordered as expected")
	}
}

func TestGetSignersFromSignatures(t *testing.T) {
	tearDownTest := setupTest(t)
	defer tearDownTest(t)

	safeTx := newDefaultSafeTx()

	owner0_sk, _ := eth.GetCryptoPrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	owner1_sk, _ := eth.GetCryptoPrivateKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	safeTx.Sign(owner0_sk)
	safeTx.Sign(owner1_sk)

	signersFromSigs := safeTx.GetSignersFromSignatures()

	if !slices.Equal(signersFromSigs, safeTx.Signers) {
		t.Errorf("signers are extracted from signatures are not as expected")
	}
}

func TestRawTx(t *testing.T) {
	tearDownTest := setupTest(t)
	defer tearDownTest(t)

	safeTx := newDefaultSafeTx()

	safeTxRaw, err := safeTx.Raw()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// t.Log(common.Bytes2Hex(safeTxRaw))
	// t.Log(len(safeTxRaw))

	abi, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		t.Fatalf(err.Error())
	}

	method := abi.Methods["execTransaction"]
	unpackedArgs, err := method.Inputs.Unpack(safeTxRaw[4:])
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(unpackedArgs...)
}

func newDefaultSafeTx() *safetx.SafeTx {
	return safetx.New(
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
}

// 6a76120200000000000000000000000000000000000000000000000000000000
// 0000000000000000000000005ac255889882aab35a2aa939679e3f3d4cea221e00
// 000000000000000000000000000000000000000000000000000000004f892b00000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000001e240000000000000000000000000000000000000000000000000000000000000007a000000000000000000000000000000000000000000000000000000000000303900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
