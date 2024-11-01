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

	safeTx := realSendEthFromSafeTx()

	safeTxRaw, err := safeTx.Raw()
	if err != nil {
		t.Fatalf(err.Error())
	}

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

func TestSafeTxExecute(t *testing.T) {
	tearDownTest := setupTest(t)
	defer tearDownTest(t)

	/* Get a sendEth Safe tx */
	safeTx := realSendEthFromSafeTx()
	/**/

	/* sign the multisig safe tx */
	owner0_sk, _ := eth.GetCryptoPrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	owner1_sk, _ := eth.GetCryptoPrivateKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	safeTx.Sign(owner0_sk)
	safeTx.Sign(owner1_sk)
	/**/

	/* Send 10 eth to the Safe account if necessary */
	desiredSafeBalance := eth.ToWei(10, 18)
	safeBalance, err := safeTx.EthereumClient.GetBalance(safeTx.SafeAddress)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if safeBalance.Uint64() < desiredSafeBalance.Uint64() {
		from, _ := eth.AddressFromPrivKey(owner0_sk.D.Bytes())
		amount_to_transfer := new(big.Int).Sub(desiredSafeBalance, safeBalance)
		gas_price, err := safeTx.EthereumClient.GasPrice()
		if err != nil {
			t.Fatalf(err.Error())
		}
		estimateGas, err := safeTx.EthereumClient.EstimateGas(*from, &safeTx.SafeAddress, 0, gas_price, nil, nil, amount_to_transfer, nil, nil, nil, nil)
		if err != nil {
			t.Fatalf(err.Error())
		}
		_, err = safeTx.EthereumClient.SendEthTo(
			"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			&safeTx.SafeAddress,
			gas_price, amount_to_transfer,
			estimateGas,
		)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	/**/

	/* Get balances before executing the SafeTx*/
	receiverBeforeBalance, err := safeTx.EthereumClient.GetBalance(safeTx.To)
	if err != nil {
		t.Fatalf(err.Error())
	}
	safeBeforeBalance, err := safeTx.EthereumClient.GetBalance(safeTx.SafeAddress)
	if err != nil {
		t.Fatalf(err.Error())
	}
	/**/

	/* Execute the Safe Tx */
	recommendedGas := safeTx.RecommendedGas()
	tx, err := safeTx.Execute(owner0_sk, &recommendedGas, safeTx.GasPrice, nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	/**/

	/* Wait for Tx confirmation and get usedGas */
	ch := safeTx.EthereumClient.WaitTxConfirmed(tx.Hash())
	isPending := <-ch
	if isPending {
		t.Fatalf("could not find tx receipt")
	}
	receipt, err := safeTx.EthereumClient.GetReceipt(tx.Hash().Hex())
	if err != nil {
		t.Fatalf(err.Error())
	}
	gasUsedBySafe := receipt.GasUsed
	/**/

	/* Check final balances are rigth */
	receiverAfterBalance, err := safeTx.EthereumClient.GetBalance(safeTx.To)
	if err != nil {
		t.Fatalf(err.Error())
	}
	safeAfterBalance, err := safeTx.EthereumClient.GetBalance(safeTx.SafeAddress)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if receiverAfterBalance.Uint64()-receiverBeforeBalance.Uint64() != eth.ToWei(1, 18).Uint64() {
		t.Fatalf("receiver balance not what expected")
	}
	if int(safeAfterBalance.Uint64())-int(safeBeforeBalance.Uint64()) == -(int(gasUsedBySafe) + int(eth.ToWei(1, 18).Int64())) {
		t.Fatalf("safe balance not what expected")
	}
	/**/
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

func realSendEthFromSafeTx() *safetx.SafeTx {
	to := common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f") // just a default anvil address
	value := eth.ToWei(1, 18)
	data := make([]byte, 0)
	operation := uint8(0)
	safeTxGasEst, err := safe_.EstimateTxGas(to, value.Uint64(), data, operation)
	if err != nil {
		return nil
	}
	baseTxGas, err := safe_.EstimateTxBaseGas(to, value.Uint64(), data, int(operation), eth.NULL_ADDRESS, safeTxGasEst)
	if err != nil {
		return nil
	}
	gas_price, err := safe_.EthereumClient.GasPrice()
	if err != nil {
		return nil
	}
	return safetx.New(
		safe_.EthereumClient,
		*safe_.SafeAddress,
		to,
		value, // send 1 eth
		data,
		operation,
		big.NewInt(int64(safeTxGasEst)),
		big.NewInt(int64(baseTxGas)),
		gas_price,
		eth.NULL_ADDRESS,
		eth.NULL_ADDRESS,
		nil,
		nil,
		"1.3.0",
	)
}
