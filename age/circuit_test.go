package age

import (
	"math/big"
	"testing"

	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/brevis-network/brevis-sdk/test"
	"github.com/ethereum/go-ethereum/common"
)

func TestCircuit(t *testing.T) {
	app, err := sdk.NewBrevisApp()
	check(err)

	contractAddress := common.HexToAddress("0xc944E90C64B2c07662A292be6244BDf05Cda44a7")

	app.AddStorage(sdk.StorageData{
		BlockNum: big.NewInt(19342663),
		Address:  contractAddress,
		Key:      common.HexToHash("0x55ccb1b16b10b19d498a335426da71059f3255a84a320fe81c2a761e2cc095d0"),
		Value:    common.HexToHash("0x0000000000000000000000000000000000000000000000252248deb6e6940000"),
	})

	guest := &AppCircuit{}
	guestAssignment := &AppCircuit{}

	circuitInput, err := app.BuildCircuitInput(guest)
	check(err)

	test.ProverSucceeded(t, guest, guestAssignment, circuitInput)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
