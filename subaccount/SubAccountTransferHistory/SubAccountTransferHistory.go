package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	SubAccountTransferHistory()
}

func SubAccountTransferHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Sub-account Transfer History (For Sub-account) - /sapi/v1/sub-account/transfer/subUserHistory
	subAccountTransferHistory, err := client.NewSubAccountTransferHistoryService().Asset("BTC").
		TransferType(1).StartTime(1234567891011).EndTime(1234567891011).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subAccountTransferHistory))
}
