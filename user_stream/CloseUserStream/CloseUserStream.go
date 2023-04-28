package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	CloseUserStream()
}

func CloseUserStream() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	close := client.NewCloseUserStream().ListenKey("your_listen_key").
		Do(context.Background())
	fmt.Println(close)
}
