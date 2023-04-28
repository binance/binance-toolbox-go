package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginIsolatedMarginTier()
}

func MarginIsolatedMarginTier() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginIsolatedMarginTierService - /sapi/v1/margin/isolatedMarginTier
	marginIsolatedMarginTier, err := client.NewMarginIsolatedMarginTierService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedMarginTier))
}
