package singlylinkedlist

// List holds a pointer to the first node in a singly-linked list with a length of zero or greater
// and List keeps track of the current number of nodes that is also always zero or greater
type List struct {
	firstNode *Node
	length    int64
}

// Node holds some data and a pointer to the next node or nil if Node is the last element in
// this particular singly-linked list
type Node struct {
	data interface{}
	next *Node
}
