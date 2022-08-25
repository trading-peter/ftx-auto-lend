module github.com/trading-peter/ftx-auto-lend

go 1.16

require (
	github.com/akamensky/argparse v1.4.0
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grishinsana/goftx v1.2.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/shopspring/decimal v1.3.1
	go.uber.org/ratelimit v0.2.0
)

replace github.com/grishinsana/goftx => ../goftx
