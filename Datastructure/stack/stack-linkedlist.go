package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type list struct {
	head *Node
}

func main() {
	l := &list{}
	l.Add(1)
	l.Add(2)
	l.Print()
	l.Remove()
	l.Print()

}

func (l *list) Add(value int) {
	newNode := &Node{data: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *list) Print() {
	if l.head == nil {
		return
	}
	curr := l.head
	for curr.next != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println(curr.data)
}

func (l *list) Remove() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}
