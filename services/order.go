package services

import (
	"indox/ticker"
	"log"

	"github.com/leekchan/accounting"
	"github.com/sirupsen/logrus"
)

func GetOrdersDetails(pair string) {
	Open, err := ticker.GetOrderBook(pair)
	if err != nil {
		log.Println(err)
	}

	var TOTALBUY float64 = 0
	var TOTALSELL float64 = 0

	// - calculate total buy
	for _, buy := range Open.Buys {
		TOTALBUY = TOTALBUY + (buy.Amount * buy.Price)
	}
	// - calculate total sell
	for _, sell := range Open.Sells {
		TOTALSELL = TOTALSELL + (sell.Amount * sell.Price)
	}

	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}
	logrus.Info("Total order buy  : ", ac.FormatMoney(TOTALBUY))
	logrus.Info("Total order sell : ", ac.FormatMoney(TOTALSELL))
	if TOTALBUY > TOTALSELL {
		logrus.Info("Order Gaps       : ", ac.FormatMoney(TOTALBUY-TOTALSELL))
		logrus.Info("Gain  Gaps       : ", TOTALBUY/TOTALSELL)
	}
}
