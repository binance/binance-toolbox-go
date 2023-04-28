package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	EnableLeverageTokenForSubAccount()
}

func EnableLeverageTokenForSubAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Enable Leverage Token for Sub-account (For Master Account) - /sapi/v1/sub-account/blvt/enable
	enableLeverageTokenForSubAccount, err := client.NewEnableLeverageTokenForSubAccountService().Email("email@email.com").
		EnableBlvt(true).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(enableLeverageTokenForSubAccount))
}
