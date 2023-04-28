package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	NewOrder()
}

func NewOrder() {

	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Binance New Order endpoint - POST /api/v3/order
	newOrder, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side("BUY").Type("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(newOrder))
}
