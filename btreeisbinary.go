package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, nil, nil)
}

// isBSTHelper is a helper function to validate the BST property
func isBSTHelper(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	// Check if the current node's value is within the valid range
	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}

	// Recursively check the left and right subtrees
	return isBSTHelper(node.Left, min, &node.Data) && isBSTHelper(node.Right, &node.Data, max)
}
