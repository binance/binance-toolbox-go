package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginIsolatedAccountTransfer()
}

func MarginIsolatedAccountTransfer() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginIsolatedAccountTransferService - /sapi/v1/margin/isolated/transfer
	marginIsolatedAccountTransfer, err := client.NewMarginIsolatedAccountTransferService().Asset("USDT").
		Symbol("BTCUSDT").TransFrom("SPOT").TransTo("ISOLATED_MARGIN").Amount(100).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedAccountTransfer))
}
