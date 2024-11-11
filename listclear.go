package main

import (
	"fmt"

	"piscine"
)

type (
	List = piscine.List
	Node = piscine.NodeL
)

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

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

// func ListClear(l *List) {
// 	for l.Head != nil {
// 		l.Head = l.Head.Next
// 	}
// }

// func main() {
// 	link := &List{}
// 	ListPushBack(link, "I")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "something")
// 	ListPushBack(link, 2)

// 	fmt.Println("------list------")
// 	PrintList(link)
// 	ListClear(link)
// 	fmt.Println("------updated list------")
// 	PrintList(link)
// }
