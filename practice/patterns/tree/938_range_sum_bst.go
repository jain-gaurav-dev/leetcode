package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var rngSum func(root *TreeNode, low, high int) int
	rngSum = func(root *TreeNode, low, high int) int {
		if root == nil {
			return 0
		}

		if root.Val < low {
			return rngSum(root.Right, low, high)
		}

		if root.Val > high {
			return rngSum(root.Left, low, high)
		}

		sum := root.Val
		sum += rngSum(root.Left, low, root.Val)
		sum += rngSum(root.Right, root.Val, high)
		return sum
	}

	return rngSum(root, low, high)
}
