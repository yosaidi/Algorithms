package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
