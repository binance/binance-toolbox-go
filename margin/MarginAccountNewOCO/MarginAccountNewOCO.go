package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginAccountNewOCO()
}

func MarginAccountNewOCO() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginAccountNewOCOService - /sapi/v1/margin/order/oco
	marginAccountNewOCO, err := client.NewMarginAccountNewOCOService().Symbol("BTCUSDT").
		Side("BUY").Quantity(0.01).Price(20000).StopPrice(18000).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOCO))
}
