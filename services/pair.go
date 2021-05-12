package services

import (
	"flag"
	"fmt"
	"indox/ticker"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

var (
	Saldo float64 = 0
)

// DetailsPairs() -- get information for ticker
func DetailsPairs(tickers, currency *string) {
	flag.Parse()
	pair := fmt.Sprintf("%s_%s", *tickers, *currency)
	ticker, err := ticker.IsDoGetPair(pair)
	if err != nil {
		log.Println(err)
	}
	// - set modal awal trading
	cap := os.Getenv("indox_init_cap")
	if cap == "" {
		logrus.Error("please add modal awal in os.env indox_init_cap")
		os.Exit(1)
	}

	CapsFloat, _ := strconv.ParseFloat(cap, 64)
	balance := ticker.Buy * Saldo
	WinLose := tablewriter.FgHiRedColor
	if balance > CapsFloat {
		WinLose = tablewriter.FgHiGreenColor
	}
	var be float64 = 0
	if Saldo != 0 {
		be = balance - CapsFloat
	}
	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}

	data := [][]string{
		[]string{
			time.Now().Local().Format("2006.01.02 15:04:05"),
			ticker.PairName,
			fmt.Sprintf("%f", ticker.Sell),
			fmt.Sprintf("%f", ticker.Buy),
			fmt.Sprintf("%f", ticker.AssetVolume),
			fmt.Sprintf("%f", ticker.BaseVolume),
			fmt.Sprintf("%f", ticker.Low),
			fmt.Sprintf("%f", ticker.High),
			ac.FormatMoney(be),
			ac.FormatMoney(balance),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Name", "Sell", "Buy", "Assets Volume", "Base Volume", "Low", "High", "Win/Lose", "Conv Assets"})
	table.SetBorder(true) // Set Border to false
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, WinLose},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
	)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
