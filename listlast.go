package piscine

func ListLast(l *List) interface{} {
	if ListSize(l) == 0 {
		return nil
	}
	return l.Tail.Data
}
