package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}
	if node.Data < root.Data {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else {
		root.Right = BTreeTransplant(root.Right, node, rplc)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}
