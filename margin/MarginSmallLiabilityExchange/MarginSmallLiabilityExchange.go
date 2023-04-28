package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginSmallLiabilityExchange()
}

func MarginSmallLiabilityExchange() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginSmallLiabilityExchangeService - /sapi/v1/margin/exchange-small-liability
	marginSmallLiabilityExchange, err := client.NewMarginSmallLiabilityExchangeService().
		AssetNames("BTC,ETH,BNB").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginSmallLiabilityExchange))
}
