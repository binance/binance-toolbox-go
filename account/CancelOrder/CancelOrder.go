package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	CancelOrder()
}

func CancelOrder() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Binance Cancel Order endpoint - DELETE /api/v3/order
	cancelOrder, err := client.NewCancelOrderService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(cancelOrder))
}
