package services

// -- function get biggest number in array (price)
// -- [1,2,3,4,5]
func GetBigest(data []int) (int, error) {
	var n, biggest int
	for _, v := range data {
		if v > n {
			n = v
			biggest = n
		}
	}
	// --return value
	return biggest, nil
}

// -- function get smallest number in array (price)
// -- [1,2,3,4,5]
func GetSmallest(data []int) (int, error) {
	var n = data[0]
	for _, v := range data {
		if v < n {
			n = v
		}
	}
	// --return value
	return n, nil
}

// -- 100 is area 1 is area to sell
// -- small is lower price
// -- high is higs prices
// -- position is fibo tertrace
func FiboArea(small, high int, position float64) float64 {
	var gaps float64 = float64(high - small)
	var retrace = gaps * position
	var a = float64(small) + retrace
	return a
}

// -- GetPositionWithFibo is decicion
// -- small is lower price
// -- high is higs prices
// -- price is price for at this time
func GetPositionWithFibo(small, high int, price float64) (buy bool, wait bool, sell bool) {

	var isBuy bool = false
	var isWaitAndSee = false
	var isSell bool = false

	isBuyArea := FiboBuyArea(small, high)
	if price < isBuyArea {
		isBuy = true
	}

	isWaitAndSeeArea := WaitAndSeeArea(small, high)
	if price >= isBuyArea && price < isWaitAndSeeArea {
		isWaitAndSee = true
	}

	isSellArea := SellArea(small, high)
	if price >= isWaitAndSeeArea && price < isSellArea {
		isSell = true
	}

	return isBuy, isWaitAndSee, isSell
}

func FiboBuyArea(small, high int) float64 {
	var gaps float64 = float64(high - small)
	var retrace = gaps * 0.236
	var a = float64(small) + retrace
	return a
}

func WaitAndSeeArea(small, high int) float64 {
	var gaps float64 = float64(high - small)
	var retrace = gaps * 0.5
	var a = float64(small) + retrace
	return a
}

func SellArea(small, high int) float64 {
	var gaps float64 = float64(high - small)
	var retrace = gaps * 1
	var a = float64(small) + retrace
	return a
}
