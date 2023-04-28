package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	FuturesTransferForSubAccount()
}

func FuturesTransferForSubAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Futures Transfer for Sub-account (For Master Account) - /sapi/v1/sub-account/futures/transfer
	futuresTransferForSubAccount, err := client.NewFuturesTransferForSubAccountService().Email("from@email.com").Asset("BTC").
		Amount(0.01).TransferType(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(futuresTransferForSubAccount))
}
