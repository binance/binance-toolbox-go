package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsDepthHandlerExample()
}

func WsDepthHandlerExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneCh
}
