package services

import (
	"encoding/json"
	"flag"
	"fmt"
	"indox/ticker"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	RESOLUTION = "D" // daily resolution in candlestick
	URI        = "https://indodax.com/tradingview/history?"
)

func Autotrade() {
	flag.Parse()
	// Get List Ticker
	getTicker, err := ticker.GetSummaries()
	if err != nil {
		log.Printf("Error when get ticker %s", err)
		os.Exit(1)
	}

	for items := range getTicker.Tickers {
		// -- convert name to tradeview format
		is_split := strings.Split(items, "_")

		if is_split[1] == *fiat {
			is_curr := fmt.Sprintf("%s_%s", is_split[0], *fiat)

			tempCurr := fmt.Sprintf("%s%s", is_split[0], *fiat)
			curr := strings.ToUpper(tempCurr)

			data, err := DOReqTradeView(curr)
			if err != nil {
				SetLogger("Warning", "Responses error in func DOReqTradeView() ")
			}

			// -- in response tradeview url , S always ok if the respose not error message
			if data.S != "ok" {
				SetLogger("Warning", "Response in not ok from tradeview Apis")
				continue
			}

			last := (getTicker.Tickers[is_curr].Buy)

			// --  calculate the price history with fibonanci
			big, _ := GetBigest(data.H)
			log.Println(big, "for ", curr, "last ", last)
		}
	}
}

// -- make http request to indodax tradeview
// -- https://indodax.com/tradingview/history
func DOReqTradeView(curr string) (tradeHistory TradeView, err error) {

	// -- generate the timestamp
	now := time.Now() // current local time

	substractday := SubstractDay(*day)
	monthago := now.AddDate(0, 0, substractday)
	to := now.Unix()
	from := monthago.Unix()

	// -- formating url
	url := fmt.Sprintf(URI+"symbol=%s&resolution=%s&from=%d&to=%d", curr, RESOLUTION, from, to)

	client := &http.Client{
		Timeout: 10 * time.Second, // - hope is enough
	}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		msg := fmt.Sprintf("error when get the data %s  | %s", url, err.Error())
		SetLogger("warning", msg)
	}

	if resp.StatusCode != http.StatusOK {
		// -- validate if response contain html
		msg := fmt.Sprintf("Http  status code  is  %d", resp.StatusCode)
		SetLogger("warning", msg)
	}

	// -- marshaling the responses
	body, err := ioutil.ReadAll(resp.Body)
	var data TradeView
	json.Unmarshal(body, &data)
	defer resp.Body.Close()

	// -- return the data
	return data, err
}

func SubstractDay(day int) int {
	if day > 0 {
		return -1 * day
	}
	return day
}
