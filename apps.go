package main

import (
	"flag"
	"fmt"
	"indox/services"
	"indox/ticker"
	"log"
	"os"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

var (
	tickers          = flag.String("t", "btc", "get all tikcer or spesifications")
	currency         = flag.String("c", "idr", "currency mode / usd or idr")
	balance          = flag.Bool("b", false, "Get saldo info")
	analitic         = flag.Bool("a", false, "use analitic")
	saldo    float64 = 0
)

func appinit() {
	logrus.Info("User Info")
	flag.Parse()
	user, err := ticker.IsDoGetInfo()
	if err != nil {
		log.Println(err)
	}
	services.Saldo = user.Balance[*tickers]
	balanceHold := user.BalanceHold[*tickers]

	bCol := tablewriter.FgHiRedColor
	if balanceHold > 0 {
		bCol = tablewriter.FgHiGreenColor
	}

	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}

	data := [][]string{
		[]string{
			time.Now().Local().Format("2006.01.02 15:04:05"),
			user.UserName,
			fmt.Sprintf("%f", user.Balance[*tickers]),
			ac.FormatMoney(user.BalanceHold["idr"]),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Name", "Balance", "Balance Hold (IDR)"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, bCol},
	)
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	pair := fmt.Sprintf("%s_%s", *tickers, *currency)
	services.GetOrdersDetails(pair)

}

func main() {
	logrus.Info("Ticker Info")

	flag.Parse()
	if !*analitic {
		appinit()
		if *tickers == "all" {
			log.Println("get all ticker")
		} else {
			services.DetailsPairs(*&tickers, *&currency)
		}
	}

	if *balance {
		services.GetEstimateBalances()
	}

	if *analitic {
		services.Analyze()
	}

}
