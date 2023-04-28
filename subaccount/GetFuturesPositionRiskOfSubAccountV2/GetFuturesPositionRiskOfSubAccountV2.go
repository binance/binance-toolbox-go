package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	GetFuturesPositionRiskOfSubAccountV2()
}

func GetFuturesPositionRiskOfSubAccountV2() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Get Futures Position-Risk of Sub-account V2 (For Master Account) - /sapi/v1/sub-account/futures/positionRisk
	getFuturesPositionRiskOfSubAccountV2, err := client.NewGetFuturesPositionRiskOfSubAccountV2Service().Email("email@email.com").
		FuturesType(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFuturesPositionRiskOfSubAccountV2))
}
