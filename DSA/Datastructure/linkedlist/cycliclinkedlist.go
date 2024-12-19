package main

//cyclic singly linked list
func (l *list) InsertAtStart(val int) {
	newNode := &Nod{data: val}
	if l.head == nil {
		newNode.next = newNode
		l.head = newNode
	} else {

		current := l.head
		for current.next != l.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = l.head
		l.head = newNode // remove this line to insert at end
	}
}

func (l *list) DeleteNode(val int) {
	if l.head == nil {
		return
	}

	curr := l.head
	prev := l.head
	// Deleting the head node
	if curr.data == val {
		if curr.next == l.head {
			l.head = nil
		} else {
			for curr.next != l.head {
				curr = curr.next
			}
			curr.next = l.head.next
			l.head = l.head.next
		}
		return
	}

	curr = curr.next
	for curr != l.head {
		if curr.data == val {
			prev.next = curr.next
			return
		}
		prev = curr
		curr = curr.next
	}
}
