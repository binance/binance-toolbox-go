package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AccountStatus()
}

func AccountStatus() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// AccountStatusService - /sapi/v1/account/status
	accountStatus, err := client.NewAccountStatusService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountStatus))
}
