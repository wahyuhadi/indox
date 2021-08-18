package services

import (
	"indox/go-indodax"
	"testing"
)

func TestGetBuySellPosition(t *testing.T) {
	var dummy = []*indodax.Order{}
	var tdata1 = &indodax.Order{
		Amount: 100000,
		Price:  2121,
	}
	var tdata2 = &indodax.Order{
		Amount: 1000000,
		Price:  2122,
	}
	var tdata3 = &indodax.Order{
		Amount: 1100000,
		Price:  2222,
	}

	dummy = append(dummy, tdata1)
	dummy = append(dummy, tdata2)
	dummy = append(dummy, tdata3)

	t.Run("Test posisition for buy and sell position ", func(t *testing.T) {
		tprices := GetBuySellPosition(dummy)
		if tprices != 2222 {
			t.Errorf("expected result is 2222 in but result is %f", tprices)
			t.Fail()
		}
	})
}
