package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QuerySubAccountSpotAssetsSummary()
}

func QuerySubAccountSpotAssetsSummary() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Sub-account Spot Assets Summary (For Master Account) - /sapi/v1/sub-account/spotSummary
	querySubAccountSpotAssetsSummary, err := client.NewQuerySubAccountSpotAssetsSummaryService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(querySubAccountSpotAssetsSummary))
}
