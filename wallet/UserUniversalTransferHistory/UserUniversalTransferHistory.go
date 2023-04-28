package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	UserUniversalTransferHistory()
}

func UserUniversalTransferHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// UserUniversalTransferHistoryService - /sapi/v1/asset/transfer
	userUniversalTransferHistory, err := client.NewUserUniversalTransferHistoryService().
		TransferType("MAIN_UMFUTURE").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userUniversalTransferHistory))
}
