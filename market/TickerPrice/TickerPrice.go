package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	TickerPrice()
}

func TickerPrice() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// TickerPrice
	tickerPrice, err := client.NewTickerPriceService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(tickerPrice))
}
