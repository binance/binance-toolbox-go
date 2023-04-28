package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetDetailOnSubAccountMarginAccount()
}

func GetDetailOnSubAccountMarginAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get Detail on Sub-account's Margin Account (For Master Account) - /sapi/v1/sub-account/margin/account
	getDetailOnSubAccountMarginAccount, err := client.NewGetDetailOnSubAccountMarginAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getDetailOnSubAccountMarginAccount))
}
