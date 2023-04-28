package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	WsUserData()
}

func WsUserData() {
	client := binance_connector.NewClient(binance_toolbox.API_KEY, binance_toolbox.SECRET_KEY, binance_toolbox.BASE_URL)

	listenKey, err := client.NewCreateListenKeyService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, binance_toolbox.WS_BASE_URL)

	wsUserDataHandler := func(event *binance_connector.WsUserDataEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsUserDataServe(listenKey, wsUserDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
