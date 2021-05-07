package services

import (
	"flag"
	"fmt"
	"indox/ticker"
	"os"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

var (
	fiat = flag.String("f", "idr", "fiat idr/usdt")
)

// type Summary struct {
//     Tickers   map[string]*Ticker
//     Prices24h map[string]float64
//     Prices7d  map[string]float64
// }
func Analyze() {
	flag.Parse()
	logrus.Info("Run analitic")
	summary, err := ticker.GetSummaries()
	if err != nil {
		logrus.Warn("Error when get data ")
	}

	data := [][]string{}
	cgaps := tablewriter.FgHiCyanColor

	for i := range summary.Tickers {
		idrsummary := strings.Split(i, "_")
		// -- makesure use the ticker
		if idrsummary[1] == *fiat {
			// data will rendered

			ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}

			//currency := fmt.Sprintf("%s%s", idrsummary[0], *fiat)
			isStrongmarket := false
			// -- get market price by pairs
			orderBook, err := ticker.GetOrderBook(i)
			if err != nil {
				continue
			}

			var TOTALBUY float64 = 0
			var TOTALSELL float64 = 0

			// - calculate total buy
			for _, buy := range orderBook.Buys {
				TOTALBUY = TOTALBUY + (buy.Amount * buy.Price)
			}
			// - calculate total sell
			for _, sell := range orderBook.Sells {
				TOTALSELL = TOTALSELL + (sell.Amount * sell.Price)
			}

			gaps := TOTALBUY / TOTALSELL
			if gaps >= 2.5 {
				isStrongmarket = true
			}

			if isStrongmarket {
				// -- if strong  market
				data = append(data,
					[]string{
						time.Now().Local().Format("2006.01.02 15:04:05"),
						i,
						ac.FormatMoney(TOTALBUY),
						ac.FormatMoney(TOTALSELL),
						"Strong market",
						fmt.Sprintf("%f", gaps),
					},
				)

			}

		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date analitic", "Name", "Order Buy", "Order sell", "Details", "Gaps"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, cgaps},
	)
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

}