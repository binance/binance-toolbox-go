package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginSmallLiabilityExchangeHistory()
}

func MarginSmallLiabilityExchangeHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginSmallLiabilityExchangeHistoryService - /sapi/v1/margin/exchange-small-liability-history
	marginSmallLiabilityExchangeHistory, err := client.NewMarginSmallLiabilityExchangeHistoryService().
		Current(1).Size(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginSmallLiabilityExchangeHistory))
}
