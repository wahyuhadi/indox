package services

import "testing"

func TestGetBiggest(t *testing.T) {
	test1 := []int{1, 2, 3, 4}
	test2 := []int{-20, 2, 4, 6, 1, 3}
	test3 := []int{-20, -2, -4, -6, -1, -3, 0, 0, 0}

	t.Run("Get Biggest skenario 1 rtest1", func(t *testing.T) {
		rtest1, _ := GetBigest(test1)
		if rtest1 != 4 {
			t.Errorf("expected result is 4 in test1 but return %d", rtest1)
			t.Fail()
		}
	})

	t.Run("Get Biggest skenario 2 rtest2", func(t *testing.T) {
		rtest2, _ := GetBigest(test2)
		if rtest2 != 6 {
			t.Errorf("expected result is 6 in test2 but return %d", rtest2)
			t.Fail()
		}
	})

	t.Run("Get Biggest skenario 3 rtest3", func(t *testing.T) {
		rtest3, _ := GetBigest(test3)
		if rtest3 != 0 {
			t.Errorf("expected result is 0 in test3 but return %d", rtest3)
			t.Fail()
		}
	})

}

func TestGetSmallest(t *testing.T) {
	test1 := []int{1, 2, 3, 4}
	test2 := []int{-20, 2, 4, 6, 1, 3}
	test3 := []int{-20, -2, -4, -6, -1, -3, 0, 0, 0, -21}

	rtest1, _ := GetSmallest(test1)
	if rtest1 != 1 {
		t.Errorf("expected result is 1 in test1 but return %d ", rtest1)
		t.Fail()
	}

	t.Run("Get Smallest skenario 1 rtest1 ", func(t *testing.T) {
		rtest1, _ := GetSmallest(test1)
		if rtest1 != 1 {
			t.Errorf("expected result is 1 in test1 but return %d ", rtest1)
			t.Fail()
		}
	})

	t.Run("Get Smallest skenario 2 rtest2 ", func(t *testing.T) {
		rtest2, _ := GetSmallest(test2)
		if rtest2 != -20 {
			t.Errorf("expected result is -20 in test2 but return %d ", rtest2)
			t.Fail()
		}
	})

	t.Run("Get Smallest skenario 3 rtest3 ", func(t *testing.T) {
		rtest3, _ := GetSmallest(test3)
		if rtest3 != -21 {
			t.Errorf("expected result is -20 in test3 but return %d ", rtest3)
			t.Fail()
		}
	})

}

func TestFiboArea(t *testing.T) {
	// -- test with max in min value
	var min = []int{1, 2, 3, 4, 5, 6, 7}
	var max = []int{100, 101, 102, 103, 104, 105, 106}
	for i := 0; i < len(min); i++ {
		t.Run("Fibonanci 0 is minimal value", func(t *testing.T) {
			fib0 := FiboArea(min[i], max[i], 0)
			if fib0 != float64(min[i]) {
				t.Errorf("expected result is %d but result %f ", min[i], fib0)
				t.Fail()
			}
		})

		t.Run("Fibonanci 100 is maximal value", func(t *testing.T) {
			fib100 := FiboArea(min[i], max[i], 1)
			if fib100 != float64(max[i]) {
				t.Errorf("expected result is %d but result %f ", min[i], fib100)
				t.Fail()
			}
		})
	}
}

func TestGetPositionWithFibo(t *testing.T) {
	t.Run("Testing with buy position", func(t *testing.T) {
		buy, wait, sell := GetPositionWithFibo(1, 100, 2)
		if buy != true {
			t.Errorf("expected result in buy is true but result %t", buy)
			t.Fail()
		}

		if wait != false {
			t.Errorf("expected result wait is false but result %t", wait)
			t.Fail()
		}

		if sell != false {
			t.Errorf("expected result sell is false but result %t", sell)
			t.Fail()
		}
	})

	t.Run("Testing with wait position", func(t *testing.T) {
		buy, wait, sell := GetPositionWithFibo(1, 100, 49)
		if buy != false {
			t.Errorf("expected result in buy is false but result %t", buy)
			t.Fail()
		}

		if wait != true {
			t.Errorf("expected result wait is true but result %t", wait)
			t.Fail()
		}

		if sell != false {
			t.Errorf("expected result sell is false but result %t", sell)
			t.Fail()
		}
	})

	t.Run("Testing with sell position", func(t *testing.T) {
		buy, wait, sell := GetPositionWithFibo(1, 100, 80)
		if buy != false {
			t.Errorf("expected result in buy is false but result %t", buy)
			t.Fail()
		}

		if wait != false {
			t.Errorf("expected result wait is true but result %t", wait)
			t.Fail()
		}

		if sell != true {
			t.Errorf("expected result sell is false but result %t", sell)
			t.Fail()
		}
	})
}
