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

// inserting a new node after the tail
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &Node{data: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}

// inserting a new node before the head
func (dll *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{data: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode

}

// insert at kth position
func (dll *DoublyLinkedList) InsertAtKthPosition(value, k int) {

	newNode := &Node{data: value}
	if k == 1 {
		if dll.head == nil {
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}
	}

	// traverse to the (k-1)th node
	current := dll.head
	for i := 0; i < k-1 && current != nil; i++ {
		current = current.next
	}

	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	} else {
		dll.tail = newNode
	}
	current.next = newNode

}

// deletion in dll
func (dll *DoublyLinkedList) DeleteNode(data int) {
	if dll.head == nil {
		return
	}

	current := dll.head
	for current != nil {
		if current.data == data {
			if current == dll.head {
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				} else {
					dll.tail = nil
				}
			} else if current == dll.tail {
				dll.tail = current.prev
				dll.tail.next = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			return
		}
		current = current.next
	}
}
