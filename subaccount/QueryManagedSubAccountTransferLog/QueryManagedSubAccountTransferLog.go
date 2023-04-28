package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	QueryManagedSubAccountTransferLog()
}

func QueryManagedSubAccountTransferLog() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Managed Sub Account Transfer Log (Investor) (USER_DATA)
	queryManagedSubAccountTransferLog, err := client.NewQueryManagedSubAccountTransferLogService().Email("email@email.com").
		StartTime(123123).EndTime(123123).Page(1).Limit(10).Transfers("").TransferFunctionAccountType("").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountTransferLog))
}
