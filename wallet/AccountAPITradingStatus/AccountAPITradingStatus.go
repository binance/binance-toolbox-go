package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AccountApiTradingStatus()
}

func AccountApiTradingStatus() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// AccountApiTradingStatusService - /sapi/v1/account/apiTradingStatus
	accountApiTradingStatus, err := client.NewAccountApiTradingStatusService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountApiTradingStatus))
}
