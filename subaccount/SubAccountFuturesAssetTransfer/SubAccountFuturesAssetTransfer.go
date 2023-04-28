package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	SubAccountFuturesAssetTransfer()
}

func SubAccountFuturesAssetTransfer() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Sub-account Futures Asset Transfer (For Master Account) - /sapi/v1/sub-account/futures/internalTransfer
	subaccountFuturesAssetTransfer, err := client.NewSubAccountFuturesAssetTransferService().FromEmail("from@email.com").
		ToEmail("to@email.com").FuturesType(1).Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountFuturesAssetTransfer))
}
