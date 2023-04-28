package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	EnableFastWithdrawSwitchService()
}

func EnableFastWithdrawSwitchService() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// EnableFastWithdrawSwitchService - /sapi/v1/account/enableFastWithdrawSwitch
	res, err := client.NewEnableFastWithdrawSwitchService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(res))
}
