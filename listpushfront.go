package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{}
	newNode.Data = data
	newNode.Next = l.Head

	l.Head = newNode
	l.Tail = newNode
}
