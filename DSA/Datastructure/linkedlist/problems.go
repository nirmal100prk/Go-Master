package main

import "fmt"

func main() {
	l1 := &list{}
	// l2 := &list{}
	// l3 := &list{}
	l1.Push(7)
	l1.Push(3)
	l1.Push(1)
	// l2.Push(5)
	// l2.Push(4)
	// l2.Push(2)
	// l1.PrintAll()
	// l2.PrintAll()
	// nod := MergeSortedLinkedList(l1.head, l2.head)
	// l3.head = nod
	// l3.PrintAll()
	l1.Push(3)
	l1.Push(7)
	fmt.Println(l1.IsPalindrome())

}

type Nod struct {
	data int
	next *Nod
}

type list struct {
	head *Nod
}

// implement stack using linked list
func (l *list) PrintAll() {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func (l *list) Push(val int) {
	newNode := &Nod{data: val}
	newNode.next = l.head
	l.head = newNode
}

func (l *list) Pop() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

// Find the nth node from the end of a linked list
// two-pointer technique
// can be solved using hashtables
func (l *list) FindNthNodeFromEnd(n int) *Nod {
	first := l.head
	second := l.head

	for i := 0; i < n; i++ {
		if first == nil {
			return nil
		}
		first = first.next
	}

	for first != nil {
		first = first.next
		second = second.next
	}
	return second
}

// check whether the given linked list is cyclic
// using hashtable
// using Floyd cycle finding algorithm - tortoise and hare agorithm
func (l *list) HasCycle() bool {
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// if the list is cyclic, find the starting node of the cycle
func (l *list) FindCycleStart() *Nod {
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			// cycle detected
			// reset one pointer to the head, to find the starting node of the cycle
			slow = l.head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			return slow // starting node of the cycle
		}
	}
	return nil
}

// Reverse a singly linked list
func (l *list) ReverseSinglyLinkedList() {
	curr := l.head
	var prev *Nod
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

// Suppose there are two singly linked lists both of which  intersect at some point
// and become a singly linked list.
// Give an algorithm for finding the merging point
// can be solved using hashtables
// can be solved using stacks(stacks for each list)

// find the middle of the linked list
// 1. traverse and find the length of the list, find n/2 node from the beginning
// 2. use two pointers, one moves 2x speed and other moves x speed. When first pointer reaches
//    end of list, second pointer will be at the middle
func (l *list) FindMiddle() *Nod {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

// is linked list even or odd
func (l *list) IsEven() bool {
	fast := l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
	}
	return fast == nil
}

// merge two sorted linked list
func MergeSortedLinkedList(l1, l2 *Nod) *Nod {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	dummy := &Nod{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			tail.next = l1
			l1 = l1.next
		} else {
			tail.next = l2
			l2 = l2.next
		}
		tail = tail.next
	}

	if l1 != nil {
		tail.next = l1
	}
	if l2 != nil {
		tail.next = l2
	}
	return dummy.next
}

// convert binary tree to a linked list

// check if a linked list is palindrome or not
func (l *list) IsPalindrome() bool {
	if l.head == nil || l.head.next == nil {
		return false
	}

	// find the middle node
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	secondHalf := reverseList(slow)

	firstHalf := l.head
	isPalindrome := true

	for secondHalf != nil {
		if firstHalf.data != secondHalf.data {
			return false
		}
		firstHalf = firstHalf.next
		secondHalf = secondHalf.next
	}
	return isPalindrome

}

func reverseList(curr *Nod) *Nod {
	var prev *Nod

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
