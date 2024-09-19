package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node has no children (Leaf node)
	if node.Left == nil && node.Right == nil {
		if node == root {
			return nil
		}
		replaceNodeInParent(node, nil)
		return root
	}

	// Case 2: Node has one child
	if node.Left == nil || node.Right == nil {
		var child *TreeNode
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}

		if node == root {
			child.Parent = nil
			return child
		}
		replaceNodeInParent(node, child)
		return root
	}

	// Case 3: Node has two children
	successor := findMin(node.Right)
	node.Data = successor.Data
	BTreeDeleteNode(root, successor)
	return root
}

// Helper function to replace a node in the parent with a new child
func replaceNodeInParent(node, newChild *TreeNode) {
	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = newChild
		} else if node.Parent.Right == node {
			node.Parent.Right = newChild
		}
	}

	if newChild != nil {
		newChild.Parent = node.Parent
	}
}

// Helper function to find the minimum node in a tree
func findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
        if root == nil {
                return nil
        }

        if elem > root.Data {
                return BTreeSearchItem(root.Right, elem)
        } else if elem < root.Data {
                return BTreeSearchItem(root.Left, elem)
        } else if elem == root.Data {
                return root
        } else {
                return root
        }
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
        if root == nil {
                return
        }
        BTreeApplyInorder(root.Left, f)
        f(root.Data)
        BTreeApplyInorder(root.Right, f)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

        if root == nil {
                return &TreeNode{Data: data}
        }

        if data < root.Data {
                root.Left = BTreeInsertData(root.Left, data)
                root.Left.Parent = root
        } else {
                root.Right = BTreeInsertData(root.Right, data)
                root.Right.Parent = root
        }

        return root
}
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
