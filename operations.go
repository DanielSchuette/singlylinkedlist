package singlylinkedlist

// New returns a new singly-linked list
func New() *List {
	return &List{nil, 0}
}

// PushBack adds a node to the front of a singly-linked list
func (l *List) PushBack(data interface{}) {
	if l.length == 0 {
		l.firstNode = &Node{data, nil}
		l.length++
		return
	}
	for n := l.Head(); ; n = n.Next() {
		if n.Next() == nil {
			n.next = &Node{data, nil}
			l.length++
			return
		}
	}
}

// PushFront adds a node to the front of a singly-linked list
func (l *List) PushFront(data interface{}) {
	temp := l.firstNode
	l.firstNode = &Node{data, temp}
	l.length++
}

// Delete deletes the first node that contains a particular item
// Delete returns true if a node was deleted, otherwise it returns false
func (l *List) Delete(item interface{}) bool {
	if l.Head().Value() == item {
		l.firstNode = l.Head().Next()
		l.length--
		return true
	}
	for n := l.Head(); ; n = n.Next() {
		if n.Next() != nil {
			if n.Next().Value() == item {
				n.next = n.Next().Next()
				l.length--
				return true
			}
		} else {
			return false
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

// Head returns a pointer to the first of the singly-linked list
func (l *List) Head() *Node {
	return l.firstNode
}

// Tail returns a pointer to the last element of the singly-linked list
func (l *List) Tail() *Node {
	for n := l.Head(); ; n = n.Next() {
		if n.Next() == nil {
			return n
		}
	}
}

// Value returns the data element of a particular node
func (n *Node) Value() interface{} {
	return n.data
}
