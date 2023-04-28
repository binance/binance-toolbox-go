package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	MarginCrossCollateralRatio()
}

func MarginCrossCollateralRatio() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// MarginCrossCollateralRatioService - /sapi/v1/margin/crossMarginCollateralRatio
	marginCrossCollateralRatio, err := client.NewMarginCrossCollateralRatioService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginCrossCollateralRatio))
}
