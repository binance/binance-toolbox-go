package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	TransferToMaster()
}

func TransferToMaster() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Transfer to Master (For Sub-account) - /sapi/v1/sub-account/transfer/subToMaster
	transferToMaster, err := client.NewTransferToMasterService().Asset("BTC").
		Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(transferToMaster))
}
