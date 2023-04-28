package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AssetDetailV2()
}

func AssetDetailV2() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// AssetDetailV2Service - /sapi/v1/asset/assetDetail
	assetDetailV2, err := client.NewAssetDetailV2Service().Asset("BTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(assetDetailV2))
}
