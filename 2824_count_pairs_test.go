package daily_challenge

import "testing"

func TestCountEasyPairs(t *testing.T) {
	res := countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2)
	t.Logf("res is %d\n", res)
}
