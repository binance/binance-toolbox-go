package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AggTradesList()
}

func AggTradesList() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// AggTradesList
	aggTradesList, err := client.NewAggTradesListService().
		Symbol("BTCUSDT").Limit(20).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(aggTradesList))
}
