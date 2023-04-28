package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	UiKlines()
}

func UiKlines() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// UiKlines
	uiKlines, err := client.NewUIKlinesService().
		Symbol("BTCUSDT").Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(uiKlines))
}
