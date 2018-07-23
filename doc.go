/*
Package singlylinkedlist implements a singly-linked list in Go without third-party libraries. Supported methods are 'PushFront' and 'PushBack'. The list keeps track of its length. Deletion of elements is implemented as well.

Code example:

	package main

	import (
		"fmt"

		sll "github.com/DanielSchuette/singlylinkedlist"
	)

	func main() {
		// create a new list
		list := sll.New()

		// add some nodes/elements to the list
		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)
		list.PushBack(4)
		list.PushBack(11)
		list.PushBack(7)
		list.PushFront(-1)
		list.PushFront(-2)
		list.PushFront(-3)

		// delete an existing and non-existing item from the list
		del := list.Delete(11)
		fmt.Printf("node deleted: %v\n", del)
		del = list.Delete(-12)
		fmt.Printf("node deleted: %v\n", del)

		// print length of list
		fmt.Printf("list length: %v\n", list.Length())

		// print list elements
		for n := list.Head(); n != nil; n = n.Next() {
			fmt.Println(n.Value())
		}
	}

`main/main.go` contains this working example. It produces the following output:

	node deleted: true
	node deleted: false
	list length: 8
	-3
	-2
	-1
	1
	2
	3
	4
	7

*/
package singlylinkedlist
