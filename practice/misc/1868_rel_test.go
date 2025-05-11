package misc

import "testing"

func TestRel(t *testing.T) {
	res := findRLEArray([][]int{[]int{1, 3}, []int{2, 1}, []int{3, 2}}, [][]int{[]int{2, 3}, []int{3, 3}})
	t.Logf("res is %v\n", res)
}
