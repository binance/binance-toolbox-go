package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	BUSDConvertHistory()
}

func BUSDConvertHistory() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// BUSDConvertHistoryService - /sapi/v1/asset/convert-transfer/queryByPage
	bUSDConvertHistory, err := client.NewBUSDConvertHistoryService().
		StartTime(1664442061000).EndTime(1664442078000).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(bUSDConvertHistory))
}
