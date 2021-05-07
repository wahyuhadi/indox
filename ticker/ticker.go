package ticker

import (
	"errors"
	"log"
	"os"
	"time"

	"indox/go-indodax"
)

var (
	Url = "https://indodax.com/api/ticker/"
)

func IsDoGetPair(pair string) (*indodax.Ticker, error) {
	cl, err := indodax.NewClient(
		"",
		"",
	)
	ticker, err := cl.GetTicker(pair)
	if err != nil {
		return nil, errors.New("Failed get ticker")
	}

	return ticker, nil
}

func IsDoGetInfo() (*indodax.UserInfo, error) {
	cl, err := indodax.NewClient(
		os.Getenv("indox_key"),
		os.Getenv("indox_sec"),
	)
	userInfo, err := cl.GetInfo()
	if err != nil {
		return nil, errors.New("error get user info")
	}

	return userInfo, nil
}

func History(pair string) ([]indodax.TradeHistory, error) {
	cl, err := indodax.NewClient(
		os.Getenv("indox_key"),
		os.Getenv("indox_sec"),
	)

	if err != nil {
		log.Println(err)
	}

	// pair name is required
	var pairName string = pair

	// set time is optional
	sinceTime := time.Now().AddDate(0, -1, 0)
	endTime := time.Now()
	log.Println(sinceTime, endTime, pair)
	// count, start trade id, and end trade id is optional
	var count, startTradeId, endTradeId int64 = 10, 1, 1000000

	// order is optional
	var sortOrder string = "desc"
	tradeHitory, err := cl.TradeHistory(pairName, count, startTradeId, endTradeId, sortOrder, &sinceTime, &endTime)
	if err != nil {
		log.Println(err)
		return tradeHitory, err
	}
	return tradeHitory, nil
}

func TransHistory() (*indodax.TransHistory, error) {
	cl, err := indodax.NewClient(
		os.Getenv("indox_key"),
		os.Getenv("indox_sec"),
	)

	if err != nil {
		log.Println(err)
	}

	transHistory, err := cl.TransHistory()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return transHistory, nil

}

func GetOrderBook(pair string) (*indodax.OrderBook, error) {
	cl, err := indodax.NewClient(
		os.Getenv("indox_key"),
		os.Getenv("indox_sec"),
	)

	if err != nil {
		log.Println(err)
	}

	orderBook, err := cl.GetOrderBook(pair)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orderBook, nil
}

// -- get summary
func GetSummaries() (*indodax.Summary, error) {
	cl, err := indodax.NewClient(
		os.Getenv("indox_key"),
		os.Getenv("indox_sec"),
	)

	if err != nil {
		log.Println(err)
	}

	summaries, err := cl.GetSummaries()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return summaries, nil
}
