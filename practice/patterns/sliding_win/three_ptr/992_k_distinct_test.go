package three_ptr

import "testing"

func TestKDistinct(t *testing.T) {
	res := subarraysWithKDistinct([]int{1, 1, 1, 2, 2, 3}, 2)
	t.Logf("res is %d\n", res)
}
