package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = current // The old head will become the new tail

	for current != nil {
		next := current.Next // Store the next node
		current.Next = prev  // Reverse the current node's pointer
		prev = current       // Move prev to the current node
		current = next       // Move to the next node
	}

	l.Head = prev // The last node processed is the new head
}
