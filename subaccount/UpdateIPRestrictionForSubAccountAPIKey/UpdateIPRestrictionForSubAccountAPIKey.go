package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	UpdateIPRestrictionForSubAccountAPIKey()
}

func UpdateIPRestrictionForSubAccountAPIKey() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Update IP Restriction for Sub-Account API key (For Master Account) - /sapi/v2/sub-account/subaccountApi/ipRestriction
	updateIPRestrictionForSubAccountAPIKey, err := client.NewUpdateIPRestrictionForSubAccountAPIKeyService().Email("email@email.com").
		SubAccountApiKey("123123").Status("").IpAddress("127.0.0.1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(updateIPRestrictionForSubAccountAPIKey))
}
