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

// 	if l.Head == nil {
// 		l.Head = newNode
// 		l.Tail = newNode
// 	} else {
// 		l.Tail.Next = newNode
// 		l.Tail = newNode
// 	}
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

// func PrintList(l *List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	fmt.Println("----normal state----")
// 	ListPushBack(link2, 1)
// 	PrintList(link2)
// 	ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "There")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "How")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)
// 	PrintList(link)

// 	ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
// }
