package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var dia func(root *TreeNode) (int, int)
	dia = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, -1
		}

		dia1, longestTillRoot1 := dia(root.Left)
		dia2, longestTillRoot2 := dia(root.Right)

		return max(dia1, max(dia2, longestTillRoot1+longestTillRoot2+2)),
			max(longestTillRoot1+1, longestTillRoot2+1)
	}

	d, _ := dia(root)
	return d
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
