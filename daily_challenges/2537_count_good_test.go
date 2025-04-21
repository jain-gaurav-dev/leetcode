package daily_challenge

import "testing"

func TestGoodSub(t *testing.T) {
	res := countGood([]int{1, 1, 1, 1, 1}, 10)
	t.Logf("res is %d\n", res)
}
