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
	fmt.Printf("list length: %v\n", list.Length())
	for n := list.Head(); n != nil; n = n.Next() {
		fmt.Println(n.Value())
	}
}
