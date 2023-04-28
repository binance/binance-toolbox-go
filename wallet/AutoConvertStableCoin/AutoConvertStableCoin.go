package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AutoConvertStableCoin()
}

func AutoConvertStableCoin() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// AutoConvertStableCoinService - /sapi/v1/capital/contract/convertible-coins
	autoConvertStableCoin, err := client.NewAutoConvertStableCoinService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(autoConvertStableCoin))
}
