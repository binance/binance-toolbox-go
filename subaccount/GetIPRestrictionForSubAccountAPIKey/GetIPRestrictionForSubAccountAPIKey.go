package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetIPRestrictionForSubAccountAPIKey()
}

func GetIPRestrictionForSubAccountAPIKey() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get IP Restriction for a Sub-account API Key (For Master Account) - /sapi/v1/sub-account/subaccountApi/ipRestriction
	getIPRestrictionForSubAccountAPIKey, err := client.NewGetIPRestrictionForSubAccountAPIKeyService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getIPRestrictionForSubAccountAPIKey))
}
