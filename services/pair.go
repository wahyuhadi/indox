package services

import (
	"flag"
	"fmt"
	"indox/ticker"
	"log"
	"os"
	"strconv"
	"strings"
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
	// - set modal awal trading
	cap := os.Getenv("indox_init_cap")
	if cap == "" {
		logrus.Error("please add modal awal in os.env indox_init_cap")
		os.Exit(1)
	}

	pair := fmt.Sprintf("%s_%s", *tickers, *currency)
	tickerpairs, err := ticker.IsDoGetPair(pair)
	if err != nil {
		log.Println(err)
	}

	orderBook, err := ticker.GetOrderBook(pair)
	if err != nil {
		log.Println("continue")
	}

	SetLogger("info", "calculate buy and sell position")
	buyposition := GetBuySellPosition(orderBook.Buys)
	sellposition := GetBuySellPosition(orderBook.Sells)
	msg := fmt.Sprintf("Buy Position in %f", buyposition)
	logrus.Info(msg)
	msg = fmt.Sprintf("Sell Position in %f", sellposition)
	logrus.Info(msg)

	CapsFloat, _ := strconv.ParseFloat(cap, 64)
	balance := tickerpairs.Buy * Saldo
	WinLose := tablewriter.FgHiRedColor
	if balance > CapsFloat {
		WinLose = tablewriter.FgHiGreenColor
	}
	var be float64 = 0
	if Saldo != 0 {
		be = balance - CapsFloat
	}
	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}
	// - get data from tradeview
	curr := fmt.Sprintf("%s%s", *tickers, *currency)
	curr = strings.ToUpper(curr)
	dataTradeView, err := DOReqTradeView(curr)
	if err != nil {
		SetLogger("Warning", "Responses error in func DOReqTradeView() ")
	}

	isdecicion := "Trade"
	if len(dataTradeView.H) == 0 || len(dataTradeView.L) == 0 {
		log.Println("continue")
	}
	SetLogger("info", "Get high and low prices")
	high, _ := GetBigest(dataTradeView.H)
	low, _ := GetSmallest(dataTradeView.L)

	SetLogger("info", "Get fibo position to take action")
	buy, wait, sell := GetPositionWithFibo(low, high, tickerpairs.Last)

	if buy {
		isdecicion = "Buy"
	}

	if wait {
		isdecicion = "wait"
	}

	if sell {
		isdecicion = "sell"
	}
	data := [][]string{
		[]string{
			time.Now().Local().Format("2006.01.02 15:04:05"),
			tickerpairs.PairName,
			fmt.Sprintf("%f", tickerpairs.Sell),
			fmt.Sprintf("%f", tickerpairs.Buy),
			fmt.Sprintf("%f", tickerpairs.BaseVolume),
			fmt.Sprintf("%f", tickerpairs.Low),
			fmt.Sprintf("%f", tickerpairs.High),
			ac.FormatMoney(be),
			ac.FormatMoney(balance),
			isdecicion,
			fmt.Sprintf("%d", *day),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Name", "Sell", "Buy", "Base Volume", "Low", "High", "Win/Lose", "Conv Assets", "Action", "Day"})
	table.SetBorder(true) // Set Border to false
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, WinLose},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
	)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
