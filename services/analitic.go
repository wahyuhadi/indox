package services

import (
	"flag"
	"fmt"
	"indox/ticker"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

var (
	fiat  = flag.String("f", "idr", "fiat idr/usdt")
	gap   = flag.Float64("g", 2.5, "gaps for order ")
	debug = flag.Bool("db", false, "debug log ")
	day   = flag.Int("d", 30, "analitic time (days)")
)

// type Summary struct {
//     Tickers   map[string]*Ticker
//     Prices24h map[string]float64
//     Prices7d  map[string]float64
// }
func Analyze() {
	log := logrus.New()
	flag.Parse()
	if !*debug {
		log.SetOutput(ioutil.Discard)
	}
	logrus.Info("Run analitic")
	summary, err := ticker.GetSummaries()
	if err != nil {
		logrus.Warn("Error when get data summary ")
		os.Exit(1)
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
			msg := fmt.Sprintf("Get Order book %s", i)
			log.Info(msg)
			orderBook, err := ticker.GetOrderBook(i)
			if err != nil {
				continue
			}

			var TOTALBUY float64 = 0
			var TOTALSELL float64 = 0

			// - calculate total buy
			msg = fmt.Sprintf("Calculate order buy %s", i)
			log.Info(msg)
			for _, buy := range orderBook.Buys {
				TOTALBUY = TOTALBUY + (buy.Amount * buy.Price)
			}
			// - calculate total sell
			msg = fmt.Sprintf("Calculate order sell %s", i)
			log.Info(msg)
			for _, sell := range orderBook.Sells {
				TOTALSELL = TOTALSELL + (sell.Amount * sell.Price)
			}

			msg = fmt.Sprintf("Counting gaps %s", i)
			log.Info(msg)
			gaps := TOTALBUY / TOTALSELL
			if gaps >= *gap {
				msg = fmt.Sprintf("Strong market %s", i)
				log.Info(msg)
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
						ac.FormatMoney(summary.Tickers[i].AssetVolume),
						ac.FormatMoney(summary.Tickers[i].BaseVolume),
					},
				)

			}

		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date analitic", "Name", "Order Buy", "Order sell", "Details", "Gaps", "Assets Volume", "Base Volume"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, cgaps},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

}
