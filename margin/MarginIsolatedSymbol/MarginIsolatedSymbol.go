package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginIsolatedSymbol()
}

func MarginIsolatedSymbol() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginIsolatedSymbolService - /sapi/v1/margin/isolated/pair
	marginIsolatedSymbol, err := client.NewMarginIsolatedSymbolService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedSymbol))
}
