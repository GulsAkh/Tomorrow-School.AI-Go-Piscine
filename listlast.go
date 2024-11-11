package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
// 	newNode := &NodeL{}
// 	newNode.Data = data
// 	newNode.Next = nil

// 	if l.Head == nil {
// 		l.Head = newNode
// 		l.Tail = newNode
// 	} else {
// 		l.Tail.Next = newNode
// 		l.Tail = newNode
// 	}
// }

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
