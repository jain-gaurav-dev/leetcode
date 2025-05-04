package three_ptr

import "testing"

func TestNiceArr(t *testing.T) {
	arr := []int{1, 2, 0, 1, 2, 1, 2, 2, 2, 2, 2, 1, 1, 1}
	res := numberOfSubarrays(arr, 4)
	t.Logf("res is %d\n", res)
}
