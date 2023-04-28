package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	BUSDConvert()
}

func BUSDConvert() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// BUSDConvertService - /sapi/v1/asset/convert-transfer
	bUSDConvert, err := client.NewBUSDConvertService().ClientTranId("118263407119").
		Asset("BUSD").Amount(20.0).AccountType("MAIN").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(bUSDConvert))
}
