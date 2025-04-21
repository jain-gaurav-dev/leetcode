package daily_challenge

import "testing"

func TestCountFairPairs(t *testing.T) {
	res := countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11)
	t.Logf("res is %d\n", res)

}
