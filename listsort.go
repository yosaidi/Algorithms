package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	sorted := true

	for sorted {
		sorted = false
		current := l

		for current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = true
			}
			current = current.Next
		}
	}
	return l
}
