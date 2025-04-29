package daily_challenge

import "testing"

func TestCompleteSubArr(t *testing.T) {
	res := countCompleteSubarraysSlidingWin([]int{5, 5, 5, 5})
	t.Logf("res is %d", res)
}
