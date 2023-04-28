package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	DeleteIPListForSubAccountAPIKey()
}

func DeleteIPListForSubAccountAPIKey() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Delete IP List For a Sub-account API Key (For Master Account) - /sapi/v1/sub-account/subaccountApi/ipRestriction/ipList
	deleteIPListForSubAccountAPIKey, err := client.NewDeleteIPListForSubAccountAPIKeyService().Email("email@email.com").
		SubAccountApiKey("123123").IpAddress("127.0.0.1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(deleteIPListForSubAccountAPIKey))
}
