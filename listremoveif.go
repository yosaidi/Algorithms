package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove matching nodes at the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// Handle remaining nodes
	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	// Update the tail
	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Tail = current
	}
}
