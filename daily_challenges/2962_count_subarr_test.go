package daily_challenge

import "testing"

func TestCountSubArr(t *testing.T) {
	res := countSubarrays([]int{1, 3, 2, 3, 3}, 2)
	t.Logf("res is %d\n", res)
}
