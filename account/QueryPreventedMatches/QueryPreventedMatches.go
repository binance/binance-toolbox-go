package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryPreventedMatches()
}

func QueryPreventedMatches() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Prevented Matches (USER_DATA) - GET /api/v3/preventedMatches
	getQueryPreventedMatchesService, err := client.NewGetQueryPreventedMatchesService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getQueryPreventedMatchesService))
}
