package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetSubAccountDepositAddress()
}

func GetSubAccountDepositAddress() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get Sub-account Deposit Address (For Master Account) - /sapi/v1/capital/deposit/subAddress
	getSubAccountDepositAddress, err := client.NewGetSubAccountDepositAddressService().Email("from@email.com").
		Coin("BTC").Network("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSubAccountDepositAddress))
}
