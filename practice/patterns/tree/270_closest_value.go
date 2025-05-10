package tree

import (
	"math"
)

func closestValue(root *TreeNode, target float64) int {
	closestVal := root.Val
	closestDiff := math.Abs(target - float64(root.Val))

	var search func(root *TreeNode)
	search = func(root *TreeNode) {
		if root == nil {
			return
		}

		diff := math.Abs(float64(root.Val) - target)
		if diff < closestDiff || (diff == closestDiff && root.Val < closestVal) {
			closestDiff = diff
			closestVal = root.Val
		}

		if target < float64(root.Val) {
			search(root.Left)
		} else if target == float64(root.Val) {
			return
		} else {
			search(root.Right)
		}
	}

	search(root)
	return closestVal
}
