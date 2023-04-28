package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginIsolatedAccountLimit()
}

func MarginIsolatedAccountLimit() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginIsolatedAccountLimitService - /sapi/v1/margin/isolated/accountLimit
	marginIsolatedAccountLimit, err := client.NewMarginIsolatedAccountLimitService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedAccountLimit))
}
