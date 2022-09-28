package student

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	middle := ListMiddle(l)
	right := ListSort(middle.Next)
	middle.Next = nil
	left := ListSort(l)
	return ListMergeNew(left, right)
}

func ListMiddle(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	slow := l
	fast := l.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func ListMergeNew(left *NodeI, right *NodeI) *NodeI {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Data < right.Data {
		left.Next = ListMergeNew(left.Next, right)
		return left
	}
	right.Next = ListMergeNew(left, right.Next)
	return right
}
