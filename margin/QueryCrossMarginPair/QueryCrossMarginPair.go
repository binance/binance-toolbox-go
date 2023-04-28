package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryCrossMarginPair()
}

func QueryCrossMarginPair() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// QueryCrossMarginPairService - /sapi/v1/margin/pair
	queryCrossMarginPair, err := client.NewQueryCrossMarginPairService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryCrossMarginPair))
}
