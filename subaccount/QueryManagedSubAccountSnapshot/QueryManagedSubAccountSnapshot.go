package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryManagedSubAccountSnapshot()
}

func QueryManagedSubAccountSnapshot() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewQueryManagedSubAccountSnapshotService().Email("email@email.com").
		SubType("BTC").StartTime(123123123).EndTime(123132123).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}
