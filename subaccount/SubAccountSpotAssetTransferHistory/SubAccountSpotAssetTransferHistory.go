package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	SubAccountSpotAssetTransferHistory()
}

func SubAccountSpotAssetTransferHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Sub-account Spot Asset Transfer History (For Master Account) - /sapi/v1/sub-account/sub/transfer/history
	subaccountSpotAssetTransferHistory, err := client.NewQuerySubAccountSpotAssetTransferHistoryService().FromEmail("from@email.com").
		ToEmail("to@email.com").StartTime(1234567891011).EndTime(1).Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountSpotAssetTransferHistory))
}
