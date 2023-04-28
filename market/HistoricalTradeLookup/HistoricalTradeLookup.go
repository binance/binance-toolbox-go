package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	HistoricalTradeLookup()
}

func HistoricalTradeLookup() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, "", binance_toolbox.BASE_URL)

	historicalTradeLookup, err := client.NewHistoricalTradeLookupService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(historicalTradeLookup))
}
