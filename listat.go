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

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	for l != nil {
		if i == pos {
			return l
		}
		l = l.Next
		i++
	}
	return nil
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
// }
