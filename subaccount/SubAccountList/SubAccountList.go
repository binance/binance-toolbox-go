package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	SubAccountList()
}

func SubAccountList() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Query Sub-account List (For Master Account) - /sapi/v1/sub-account/list
	subaccountList, err := client.NewQuerySubAccountListService().Email("test@email.com").
		IsFreeze("").Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountList))
}
