package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
	binance_toolbox "github.com/binance/binance-toolbox-go"
)

func main() {
	WsCombinedDepthHandlerExample()
}

func WsCombinedDepthHandlerExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(true, binance_toolbox.WS_BASE_URL)
	wsCombinedDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsCombinedDepthServe([]string{"LTCBTC", "BTCUSDT", "MATICUSDT"}, wsCombinedDepthHandler, errHandler)
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
