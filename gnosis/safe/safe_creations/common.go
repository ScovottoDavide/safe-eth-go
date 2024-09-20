package safecreations

import (
	"math"
	"math/big"

	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth"
	"github.com/ScovottoDavide/safe-eth-go/gnosis/eth/contracts"
	"github.com/ethereum/go-ethereum/common"
)

func getInitialSetupSafeData(
	owners []common.Address,
	threshold int64,
	payment uint64,
	fallbackHandler common.Address,
	paymentToken common.Address,
	paymentReceiver common.Address,
) ([]byte, error) {

	refundPayment := big.NewInt(0)
	if paymentReceiver != eth.NULL_ADDRESS {
		refundPayment = big.NewInt(int64(payment))
	}

	/* construct the initialization data needed for the proxy to initialize the Safe */
	safeAbi, err := contracts.GnosisSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	method := safeAbi.Methods["setup"].Name
	initializer, err := safeAbi.Pack(
		method,
		owners,
		big.NewInt(threshold),
		eth.NULL_ADDRESS, // Contract address for optional delegate call
		make([]byte, 0),  // Data payload for optional delegate call
		fallbackHandler,
		paymentToken,
		refundPayment,
		paymentReceiver,
	)
	if err != nil {
		return nil, err
	}
	return initializer, nil
}

func calculatePayment(
	fixedCreationCost int,
	paymentTokenEthValue float64,
	gas int64,
	gasPrice int64,
) uint64 {
	if fixedCreationCost == 0 {
		// Payment will be safe deploy cost + transfer fees for sending ether to the deployer
		basePayment := (gas) * gasPrice
		// Calculate payment for tokens using the conversion (if used)
		return uint64(math.Ceil(float64(basePayment) / paymentTokenEthValue))
	} else {
		return uint64(fixedCreationCost)
	}
}
