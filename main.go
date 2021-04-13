package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	"github.com/avast/retry-go"
	"github.com/grishinsana/goftx"
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
	"go.uber.org/ratelimit"
)

var (
	limiter ratelimit.Limiter = ratelimit.New(30)
	client  *goftx.Client
)

func main() {
	job := cron.New()
	parser := argparse.NewParser("ftx-auto-lend", "Automatically compounds lending payouts.")
	apiKey := parser.String("k", "key", &argparse.Options{Required: true, Help: "API key"})
	apiSecret := parser.String("s", "secret", &argparse.Options{Required: true, Help: "API secret"})
	subAcc := parser.String("a", "subaccount", &argparse.Options{Required: false, Help: "Subaccount"})
	coinList := parser.List("c", "coin", &argparse.Options{Required: false, Help: "Coin to lend"})
	rate := parser.String("r", "min-rate", &argparse.Options{Required: false, Help: "Coin to lend"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	strCoins := []string{}
	strRate := *rate

	if len(*coinList) == 0 {
		strCoins = append(strCoins, "USD")
	} else {
		for i := range *coinList {
			strCoins = append(strCoins, strings.ToUpper((*coinList)[i]))
		}
	}

	if strRate == "" {
		strRate = "0.000001"
	}

	minRate, err := decimal.NewFromString(strRate)

	if err != nil {
		Error.Fatal("Min Rate: Invalid number")
	}

	client = goftx.New(
		goftx.WithAuth(*apiKey, *apiSecret),
		goftx.WithSubaccount(*subAcc),
	)

	_, err = client.GetAccountInformation()

	if err != nil {
		Error.Fatalln("It seems like the supplied API key is wrong. Please check and try again")
	}

	job.Start()

	job.AddFunc("5 * * * *", func() {
		for i := range strCoins {
			coin := strCoins[i]
			Info.Printf("Running lending offer update for %s.\n", coin)
			lendable, delta, err := getMaxLendingAmount(coin)

			if err != nil {
				Error.Println(err)
				continue
			}

			if delta.Equal(decimal.Zero) || delta.LessThan(decimal.Zero) {
				Info.Println("No increase in funds to update lending offer.")
				continue
			}

			Info.Printf("New lendable amount of %s is %s (+%s).", coin, lendable, delta)
			Info.Printf("Updating lending offer with a minimum rate of %s.", minRate)

			err = updateLendingOffer(coin, lendable, minRate)

			if err != nil {
				Error.Println(err)
				continue
			}
		}
	})

	fmt.Printf("I will attempt to update your lending offers once per hour.\nPress any key if you want to stop and exit the program.")
	fmt.Scanln()
	fmt.Println("Bye!")
}

func updateLendingOffer(coin string, amount decimal.Decimal, minRate decimal.Decimal) (err error) {
	err = retry.Do(
		func() error {
			limiter.Take()
			err := client.SubmitLendingOffer(coin, amount, minRate)

			if err != nil {
				fmt.Printf("%+v\n", err)
				return err
			}

			return nil
		},
		retry.Delay(time.Minute),
		retry.Attempts(10),
		retry.DelayType(retry.FixedDelay),
	)

	return
}

func getMaxLendingAmount(coin string) (lendable decimal.Decimal, delta decimal.Decimal, err error) {
	err = retry.Do(
		func() error {
			limiter.Take()
			resp, err := client.GetLendingInfo()

			if err != nil {
				return err
			}

			for i := range resp {
				if resp[i].Coin == coin {
					lendable = resp[i].Lendable
					delta = resp[i].Lendable.Sub(resp[i].Offered)
				}
			}

			return nil
		},
		retry.Delay(time.Minute),
		retry.Attempts(10),
		retry.DelayType(retry.FixedDelay),
	)

	return
}
