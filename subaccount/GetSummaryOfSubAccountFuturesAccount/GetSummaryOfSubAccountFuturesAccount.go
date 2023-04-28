package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetSummaryOfSubAccountFuturesAccount()
}

func GetSummaryOfSubAccountFuturesAccount() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get Summary of Sub-account's Futures Account (For Master Account) - /sapi/v1/sub-account/futures/accountSummary
	getSummaryOfSubAccountFuturesAccount, err := client.NewGetSummaryOfSubAccountFuturesAccountService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSummaryOfSubAccountFuturesAccount))
}
