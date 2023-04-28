package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	DepositAssetsIntoManagedSubAccount()
}

func DepositAssetsIntoManagedSubAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Deposit Assets Into The Managed Sub-account（For Investor Master Account） - /sapi/v1/sub-account/managed-subaccount/deposit
	depositAssetsIntoManagedSubAccount, err := client.NewDepositAssetsIntoManagedSubAccountService().ToEmail("to@email.com").
		Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(depositAssetsIntoManagedSubAccount))
}
