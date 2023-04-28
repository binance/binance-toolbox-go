package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	CancelReplace()
}

func CancelReplace() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Cancel an Existing Order and Send a New Order (TRADE) - POST /api/v3/order/cancelReplace
	cancelReplace, err := client.NewCancelReplaceService().
		Symbol("BTCUSDT").Side("BUY").OrderType("LIMIT").CancelReplaceMode("STOP_ON_FAILURE").
		TimeInForce("GTC").Quantity(0.001).Price(20000.0).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(cancelReplace))
}
