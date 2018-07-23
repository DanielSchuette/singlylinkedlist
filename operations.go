package singlylinkedlist

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
		l.firstNode = &Node{data, l.length, nil}
		l.length++
		return
	}
	for n := l.firstNode; ; n = n.Next() {
		if n.Next() == nil {
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

// Head returns a pointer to the head of the singly-linked list
func (l *List) Head() *Node {
	return l.firstNode
}

// Value returns the data element of a particular node
func (n *Node) Value() interface{} {
	return n.data
}
