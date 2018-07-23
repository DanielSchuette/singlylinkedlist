package main

import (
	"fmt"

	"github.com/DanielSchuette/singlylinkedlist"
)

func main() {
	l := singlylinkedlist.New()
	l.PushBack([]int{1, 2, 3})
	l.PushBack([]int{4, 5, 6})
	for n := l.Head(); n != nil; n = n.Next() {
		fmt.Println(n.Value())
	}
}
