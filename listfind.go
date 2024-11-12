package piscine

// import (
// 	"fmt"
// )

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
// 	node := &NodeL{}
// 	node.Data = data
// 	node.Next = nil

// 	if l.Head == nil {
// 		l.Head = node
// 		l.Tail = node
// 	} else {
// 		l.Tail.Next = node
// 		l.Tail = node
// 	}
// }

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "hello1")
// 	ListPushBack(link, "hello2")
// 	ListPushBack(link, "hello3")

// 	found := ListFind(link, interface{}("hello2"), CompStr)

// 	fmt.Println(found)
// 	fmt.Println(*found)
// }
