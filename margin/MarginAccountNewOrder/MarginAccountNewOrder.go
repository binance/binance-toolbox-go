package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginAccountNewOrder()
}

func MarginAccountNewOrder() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginAccountNewOrderService - /sapi/v1/margin/order
	marginAccountNewOrder, err := client.NewMarginAccountNewOrderService().Symbol("BTCUSDT").
		Side("BUY").OrderType("MARKET").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOrder))
}
