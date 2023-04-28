package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	RecentTradesList()
}

func RecentTradesList() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// RecentTradesList
	recentTradesList, err := client.NewRecentTradesListService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(recentTradesList))
}
