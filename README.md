# binance-toolbox-go

A collection of Go examples that connect to the Binance Spot API endpoints based on `binance-connector-go`
- Github repository: https://github.com/binance/binance-connector-go

## Installation

Ensure that the `go.mod` file contains the latest version of the Binance Go Connector, then execute the following:
```bash
go mod tidy
go install github.com/binance/binance-connector-go
```

## Usage

First, open `config.go` and replace the placeholders with your API Key and Secret Key. Depending on which endpoint you wish to use, you may also need to replace the Base URL and Websocket URL.

In order to run an example, execute `go run market/ExchangeInfo/ExchangeInfo.go` from the root directory. Replace `market/ExchangeInfo/ExchangeInfo.go` with the specific example you wish to run. Ensure you add your API Key and Secret Key (and Base URL if necessary) to the relevant file before running the example.