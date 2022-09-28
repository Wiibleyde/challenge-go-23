package student

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return &NodeI{Data: data_ref}
	}
	if data_ref <= l.Data {
		return &NodeI{Data: data_ref, Next: l}
	}
	l.Next = SortListInsert(l.Next, data_ref)
	return l
}
