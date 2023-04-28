package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	AggTradesExample()
}

func AggTradesExample() {
	// Initialise Websocket Client with Testnet BaseURL and false for "Combined" parameter
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, binance_toolbox.WS_BASE_URL)

	wsAggTradeHandler := func(event *binance_connector.WsAggTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAggTradeServe("BTCUSDT", wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
