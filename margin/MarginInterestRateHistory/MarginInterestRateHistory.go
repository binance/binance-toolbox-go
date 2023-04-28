package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginInterestRateHistory()
}

func MarginInterestRateHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginInterestRateHistoryService - /sapi/v1/margin/interestRateHistory
	marginInterestRateHistory, err := client.NewMarginInterestRateHistoryService().Asset("USDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginInterestRateHistory))
}
