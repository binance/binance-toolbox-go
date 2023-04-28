package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	UserUniversalTransfer()
}

func UserUniversalTransfer() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// UserUniversalTransferService - /sapi/v1/asset/transfer
	userUniversalTransfer, err := client.NewUserUniversalTransferService().
		TransferType("MAIN_UMFUTURE").Asset("USDT").Amount(20.50).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userUniversalTransfer))
}
