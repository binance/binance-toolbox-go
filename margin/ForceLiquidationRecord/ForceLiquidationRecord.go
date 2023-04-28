package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	ForceLiquidationRecord()
}

func ForceLiquidationRecord() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// ForceLiquidationRecordService - /sapi/v1/margin/forceLiquidationRec
	forceLiquidationRecord, err := client.NewForceLiquidationRecordService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(forceLiquidationRecord))
}
