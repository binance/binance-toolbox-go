package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryManagedSubAccountFuturesAssetDetails()
}

func QueryManagedSubAccountFuturesAssetDetails() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Managed Sub-account Futures Asset Details（For Investor Master Account）(USER_DATA)
	queryManagedSubAccountFuturesAssetDetails, err := client.NewQueryManagedSubAccountFuturesAssetDetailsService().Email("email@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountFuturesAssetDetails))
}
