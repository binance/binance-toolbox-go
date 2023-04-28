package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	DustTransfer()
}

func DustTransfer() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// DustTransferService - /sapi/v1/asset/dust
	dustTransfer, err := client.NewDustTransferService().Asset([]string{"ETH", "LTC", "TRX"}).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(dustTransfer))
}
