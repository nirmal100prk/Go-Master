package main

import "fmt"

type Nod struct {
	data int
	next *Nod
}

type list struct {
	head *Nod
}

func main() {
	l := &list{}
	l.AddRecursively(1)
	l.AddRecursively(2)
	l.Print()
}

func (l *list) Add(value int) {
	newNod := &Nod{data: value}
	if l.head == nil {
		l.head = newNod
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNod

}

func (l *list) Remove(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}
	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

//Recursive
func (l *list) AddRecursively(value int) {
	l.head = AddRec(l.head, value)
}

func AddRec(n *Nod, value int) *Nod {
	if n == nil {
		n = &Nod{data: value}
		return n
	}
	n.next = AddRec(n.next, value)
	return n
}

func (l *list) DeleteRecursively(value int) {
	l.head = DeleteRec(l.head, value)
}

func DeleteRec(n *Nod, value int) *Nod {
	if n == nil {
		return nil
	}

	if n.data == value {
		return n.next
	}
	n.next = DeleteRec(n.next, value)
	return n
}

func (l *list) Print() {
	printRec(l.head)
}

func printRec(n *Nod) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	printRec(n.next)
}
