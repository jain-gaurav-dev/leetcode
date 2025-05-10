package tree

// * Definition for Node.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		ptr := node.Right

		for ptr.Left != nil {
			ptr = ptr.Left
		}

		return ptr
	}

	curr, parent := node, node.Parent
	for parent != nil && curr == parent.Right {
		curr, parent = parent, parent.Parent
	}

	return parent
}
