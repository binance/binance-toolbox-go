package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	WithdrawAssetsFromTheManagedSubAccount()
}

func WithdrawAssetsFromTheManagedSubAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewWithdrawAssetsFromTheManagedSubAccountService().FromEmail("email@email.com").
		Asset("BTC").Amount(1.5).TransferDate(123132123).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}
