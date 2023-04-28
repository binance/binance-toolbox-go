package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetMyTrades()
}

func GetMyTrades() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Binance Get trades for a specific account and symbol (USER_DATA) - GET /api/v3/myTrades
	getMyTradesService, err := client.NewGetMyTradesService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getMyTradesService))
}
