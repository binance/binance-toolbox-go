package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	Ping()
}

func Ping() {
	client := binance_connector.NewClient("", "", binance_toolbox.BASE_URL)

	// NewPingService
	ping := client.NewPingService().Do(context.Background())
	if ping == nil {
		fmt.Println("Success")
		return
	}
	fmt.Println(binance_connector.PrettyPrint(ping))
}
