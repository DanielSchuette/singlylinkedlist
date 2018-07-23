/*
A singly-linked list implementation in Go. Supported methods are 'PushFront' and 'PushBack'. The list keeps
track of its length. Code example:

	func main() {
		list := singlylinkedlist.New()
		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)
		list.PushBack(4)
		list.PushBack(11)
		list.PushBack(7)
		list.PushFront(-1)
		list.PushFront(-2)
		list.PushFront(-3)
		fmt.Printf("list length: %v\n", list.Length())
		for n := list.Head(); n != nil; n = n.Next() {
			fmt.Println(n.Value())
		}
	}

*/

package main

import (
	"fmt"

	"github.com/DanielSchuette/singlylinkedlist"
)

func main() {
	list := singlylinkedlist.New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(11)
	list.PushBack(7)
	list.PushFront(-1)
	list.PushFront(-2)
	list.PushFront(-3)
	fmt.Printf("list length: %v\n", list.Length())
	for n := list.Head(); n != nil; n = n.Next() {
		fmt.Println(n.Value())
	}
}
