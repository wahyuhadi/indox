package services

import (
	"encoding/json"
	"errors"
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
		curr := "BTCIDR"

		// -- convert name to tradeview format
		is_split := strings.Split(items, "_")
		if is_split[1] == *fiat {
			tempCurr := fmt.Sprintf("%s%s", is_split[0], *fiat)
			curr = strings.ToUpper(tempCurr)
		}

		// -- generate the timestamp
		now := time.Now() // current local time
		monthago := now.AddDate(0, -1, 0)
		to := now.Unix()
		from := monthago.Unix()

		// -- formating url
		url := fmt.Sprintf(URI+"symbol=%s&resolution=%s&from=%d&to=%d", curr, RESOLUTION, from, to)
		data, err := DOReqTradeView(url)
		if err != nil {
			SetLogger("Warning", "Responses error in func DOReqTradeView() ")
		}

		// -- in response tradeview url , S always ok if the respose not error message
		if data.S != "ok" {
			SetLogger("Warning", "Response in not ok from tradeview Apis")
			continue
		}

		// --  calculate the price history with fibonanci
	}
}

// -- make http request to indodax tradeview
// -- https://indodax.com/tradingview/history
func DOReqTradeView(url string) (tradeHistory TradeView, err error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // - hope is enough
	}

	// -- dont follow the redirect page
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
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
