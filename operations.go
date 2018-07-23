package singlylinkedlist

import "fmt"

// New returns a new singly-linked list
func New() *List {
	return &List{nil, 0}
}

// PushFront adds a node to the front of a singly-linked list
func (l *List) PushFront(data interface{}) {

}

// PushBack adds a node to the front of a singly-linked list
func (l *List) PushBack(data interface{}) {
	if l.length == 0 {
		fmt.Println("current list is length 0")
		l.firstNode = &Node{data, l.length, nil}
		l.length++
		return
	}
	fmt.Printf("l.firstNode: %v, l.firstNode.Next(): %v\n", l.firstNode, l.firstNode.Next())
	for n := l.firstNode; n != nil; n = n.Next() {
		fmt.Println("checking next node")
		if n.Next() == nil {
			fmt.Println("node is nil, adding new node")
			n.next = &Node{data, l.length, nil}
			l.length++
			return
		}
	}
}

// Length returns the length of a singly-linked list
func (l *List) Length() int64 {
	return l.length
}

// Next returns the next pointer of a particular node
func (n *Node) Next() *Node {
	return n.next
}
