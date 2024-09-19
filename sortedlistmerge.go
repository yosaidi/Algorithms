package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	// Start with the list that has the smallest element
	var mergedHead *NodeI
	if n1.Data <= n2.Data {
		mergedHead = n1
		n1 = n1.Next
	} else {
		mergedHead = n2
		n2 = n2.Next
	}

	// Keep track of the current node in the merged list
	current := mergedHead

	// Traverse both lists and merge them
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// If one list is exhausted, append the other list
	if n1 != nil {
		current.Next = n1
	} else if n2 != nil {
		current.Next = n2
	}

	return mergedHead
}
