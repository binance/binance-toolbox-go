package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	DisableFastWithdrawSwitch()
}

func DisableFastWithdrawSwitch() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// DisableFastWithdrawSwitchService  - /sapi/v1/account/disableFastWithdrawSwitch
	disableFastWithdrawSwitch, err := client.NewDisableFastWithdrawSwitchService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(disableFastWithdrawSwitch))
}
