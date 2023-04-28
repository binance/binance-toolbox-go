package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// ExchangeInfo
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(exchangeInfo))
}
