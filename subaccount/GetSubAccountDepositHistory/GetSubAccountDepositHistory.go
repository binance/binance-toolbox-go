package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetSubAccountDepositHistory()
}

func GetSubAccountDepositHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get Sub-account Deposit History (For Master Account) - /sapi/v1/capital/deposit/subHisrec
	getSubAccountDepositHistory, err := client.NewGetSubAccountDepositHistoryService().Email("from@email.com").
		Coin("BTC").Status(1).StartTime(1234567891011).EndTime(1234567891011).Limit(10).Offset(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSubAccountDepositHistory))
}
