package services

type TradeView struct {
	S string    `json:"s"`
	T []int     `json:"t"`
	C []int     `json:"c"`
	O []int     `json:"o"`
	H []int     `json:"h"`
	L []int     `json:"l"`
	V []float64 `json:"v"`
}
