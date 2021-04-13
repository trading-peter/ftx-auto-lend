module github.com/trading-peter/ftx-auto-lend

go 1.16

require (
	github.com/akamensky/argparse v1.2.2
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/grishinsana/goftx v1.2.0
	github.com/robfig/cron/v3 v3.0.0
	github.com/shopspring/decimal v1.2.0
	go.uber.org/ratelimit v0.2.0
)

replace github.com/grishinsana/goftx => ../goftx
