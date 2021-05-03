package services

import (
	"fmt"
	"indox/ticker"
	"log"
	"os"
	"time"

	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

func GetEstimateBalances() {
	user, err := ticker.IsDoGetInfo()
	if err != nil {
		logrus.Error(err)
	}

	var balance float64 = 0
	for i, item := range user.Balance {
		if user.Balance[i] > 0 {
			pairname := fmt.Sprintf("%s_idr", i)

			pair, err := ticker.IsDoGetPair(pairname)
			if err != nil {
				log.Println(err)
			}
			balance = balance + (item * pair.Buy)
		}
	}
	logrus.Info("Calculate assets")
	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}

	data := [][]string{
		[]string{
			time.Now().Local().Format("2006.01.02 15:04:05"),
			ac.FormatMoney(balance),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Assets IDR"})
	table.SetBorder(true) // Set Border to false
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

}
