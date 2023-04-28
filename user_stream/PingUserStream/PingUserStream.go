package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	PingUserStream()
}

func PingUserStream() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	ping := client.NewPingUserStream().ListenKey("your_listen_key").Do(context.Background())
	fmt.Println(ping)
}
