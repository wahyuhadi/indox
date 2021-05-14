package services

import "testing"

func TestSubstractDay(t *testing.T) {
	t.Run("Run with positive number ", func(t *testing.T) {
		num := 1
		subs := SubstractDay(num)
		if subs != -1 {
			t.Errorf("expected result is -1 in but return %d ", subs)
			t.Fail()
		}
	})

	t.Run("Run with negatif number ", func(t *testing.T) {
		num := -30
		subs := SubstractDay(num)
		if subs != -30 {
			t.Errorf("expected result is -30 in but return %d ", subs)
			t.Fail()
		}
	})
}
