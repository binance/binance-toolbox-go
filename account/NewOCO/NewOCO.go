package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	NewOCO()
}

func NewOCO() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	// Binance New OCO (TRADE) - POST /api/v3/order/oco
	newOCO, err := client.NewNewOCOService().Symbol("LTCBNB").
		Side("BUY").Quantity(0.1).Price(0.28).StopPrice(0.22).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(newOCO))
}
