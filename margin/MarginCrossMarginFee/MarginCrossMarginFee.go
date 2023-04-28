package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginCrossMarginFee()
}

func MarginCrossMarginFee() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginCrossMarginFeeService - /sapi/v1/margin/crossMarginData
	marginCrossMarginFee, err := client.NewMarginCrossMarginFeeService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginCrossMarginFee))
}
