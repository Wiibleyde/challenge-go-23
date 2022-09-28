package student

type NodeLB struct {
	Data interface{}
	Next *NodeLB
}

type ListB struct {
	Head *NodeLB
	Tail *NodeLB
}

func ListPushBack(l *ListB, data interface{}) {
	node := &NodeLB{Data: data}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}
