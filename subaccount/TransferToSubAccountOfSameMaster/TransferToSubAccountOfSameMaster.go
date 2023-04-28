package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	TransferToSubAccountOfSameMaster()
}

func TransferToSubAccountOfSameMaster() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Transfer to Sub-account of Same Master (For Sub-account) - /sapi/v1/sub-account/transfer/subToSub
	transferToSubAccountOfSameMaster, err := client.NewTransferToSubAccountOfSameMasterService().ToEmail("from@email.com").Asset("BTC").
		Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(transferToSubAccountOfSameMaster))
}
