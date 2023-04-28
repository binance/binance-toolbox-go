package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QuerySubAccountAssets()
}

func QuerySubAccountAssets() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Sub-account Assets (For Master Account) - /sapi/v3/sub-account/assets
	querySubAccountAssets, err := client.NewQuerySubAccountAssetsService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(querySubAccountAssets))
}
