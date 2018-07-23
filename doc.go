/*
Package singlylinkedlist implements a singly-linked list in Go without third-party libraries. Supported methods are 'PushFront' and 'PushBack'. The list keeps track of its length. Code example:

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

`main/main.go` contains this working example. It produces the following output:

	list length: 9
	-3
	-2
	-1
	1
	2
	3
	4
	11
	7

*/
package singlylinkedlist
