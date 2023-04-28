package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryManagedSubAccountAssetDetails()
}

func QueryManagedSubAccountAssetDetails() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Managed Sub-account Asset Details（For Investor Master Account）- /sapi/v1/sub-account/managed-subaccount/asset
	queryManagedSubAccountAssetDetails, err := client.NewQueryManagedSubAccountAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountAssetDetails))
}
