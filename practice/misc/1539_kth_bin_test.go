package misc

import "testing"

func TestKthMissingBin(t *testing.T) {
	res := findKthPositiveBinSearch([]int{2, 3, 4, 7, 11}, 5)
	t.Logf("res is %d\n", res)
}
