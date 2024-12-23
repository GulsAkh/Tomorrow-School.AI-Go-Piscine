package piscine

// import "fmt"

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

// Iterative reverse
func ListReverse(l *List) {
	var prev *NodeL
	l.Tail = l.Head
	for l.Head != nil {
		temp := l.Head.Next
		l.Head.Next = prev
		// l.Head.Next (last element) == nil
		prev = l.Head
		l.Head = temp
	}
	l.Head = prev
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }
