package tree

func delNodesOptimized(root *TreeNode, to_delete []int) []*TreeNode {
	ds := make(map[int]bool)
	for _, d := range to_delete {
		ds[d] = true
	}

	roots := make(map[*TreeNode]bool)
	roots[root] = true

	var insertChildrenAsRoots func(root *TreeNode)
	insertChildrenAsRoots = func(root *TreeNode) {
		if root.Left != nil {
			roots[root.Left] = true
		}
		if root.Right != nil {
			roots[root.Right] = true
		}
	}

	var searchAndDel func(par, root *TreeNode)
	searchAndDel = func(par, root *TreeNode) {
		// use post-order traversal, search and delete in left and right subtrees first
		if root == nil {
			return
		}

		searchAndDel(root, root.Left)
		searchAndDel(root, root.Right)

		if ds[root.Val] {
			// node pointed by root needs to be deleted
			if par == nil {
				delete(roots, root)
			} else {
				// parent exists
				if par.Left == root {
					par.Left = nil
				} else {
					par.Right = nil
				}
			}

			insertChildrenAsRoots(root)
		}
	}

	searchAndDel(nil, root)
	res := make([]*TreeNode, 0, len(roots))
	for k := range roots {
		res = append(res, k)
	}

	return res
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	roots := make(map[*TreeNode]bool)
	roots[root] = true

	var searchAndDel func(par, root *TreeNode, v int) bool
	searchAndDel = func(par, root *TreeNode, v int) bool {
		if root == nil {
			return false
		}

		if root.Val == v {
			if par == nil {
				// root needs to be deleted
				delete(roots, root)
				// insert new roots
				if root.Left != nil {
					roots[root.Left] = true
				}
				if root.Right != nil {
					roots[root.Right] = true
				}

				return true
			} else {
				if par.Left == root {
					par.Left = nil
				} else if par.Right == root {
					par.Right = nil
				}

				if root.Left != nil {
					roots[root.Left] = true
				}
				if root.Right != nil {
					roots[root.Right] = true
				}

				return true
			}
		}
		res := searchAndDel(root, root.Left, v)
		if res {
			return true
		}

		return searchAndDel(root, root.Right, v)
	}

	var del func(v int)
	del = func(v int) {
		for r := range roots {
			res := searchAndDel(nil, r, v)
			if res {
				return
			}
		}
	}

	for _, d := range to_delete {
		del(d)
	}

	res := make([]*TreeNode, 0, len(roots))
	for r := range roots {
		res = append(res, r)
	}

	return res
}
