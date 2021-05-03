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

	logrus.Info("Running analitic ....")

	ac := accounting.Accounting{Symbol: "RP ", Precision: 2, Thousand: ".", Decimal: ","}
	logrus.Info("Total order buy  : ", ac.FormatMoney(TOTALBUY))
	logrus.Info("Total order sell : ", ac.FormatMoney(TOTALSELL))
	gaps := TOTALBUY / TOTALSELL
	if TOTALBUY > TOTALSELL {
		logrus.Info("Order Gaps       : ", ac.FormatMoney(TOTALBUY-TOTALSELL))
		logrus.Info("Gain  Gaps       : ", gaps)
	} else {
		logrus.Warn("Order Gaps       : ", ac.FormatMoney(TOTALBUY-TOTALSELL))
		logrus.Warn("Gain  Gaps       : ", TOTALBUY/TOTALSELL)
	}
	if gaps >= 5 {
		logrus.Info("Gain Severity    : Very Strong market")
	}

	if gaps >= 2 && gaps < 5 {
		logrus.Info("Gain Severity    : Strong market")
	}

	if gaps >= 1 && gaps < 2 {
		logrus.Warn("Gain Severity    : Unstable market")
	}

	if gaps < 1 {
		logrus.Warn("Gain Severity    : Not recomendation market")
	}
}
