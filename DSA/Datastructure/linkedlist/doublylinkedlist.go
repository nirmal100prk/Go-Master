package main

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

/*
To append a new node with data 40, we set the next pointer of the current tail (Node 3)
to the new node and set the prev pointer of the new node to Node 3.
The tail is then updated to the new node.
*/
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &Node{data: value, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}

func (dll *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{data: value, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}

}
