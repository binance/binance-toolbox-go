package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AccountSnapshot()
}

func AccountSnapshot() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// GetAccountSnapshotService get all orders from account - /sapi/v1/accountSnapshot
	accountSnapshot, err := client.NewGetAccountSnapshotService().MarketType("SPOT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountSnapshot))
}
