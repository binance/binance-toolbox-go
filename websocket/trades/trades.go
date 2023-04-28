package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	WsTradeExample()
}

func WsTradeExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, binance_toolbox.WS_BASE_URL)
	wsTradeHandler := func(event *binance_connector.WsTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsTradeServe("LTCBTC", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
