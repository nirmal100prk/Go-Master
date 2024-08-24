package main

import (
	"fmt"
)

func main() {
	S := Stack[int]{}

	S.Push(1)
	S.Push(2)
	fmt.Println(S)
	S.Pop()
	fmt.Println(S)

}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() {
	s.items = s.items[:len(s.items)-1]
}

//////////////////////////////////////////

// func main() {
// 	fmt.Println(Add(1, 2))
// 	fmt.Println(Add(1.2, 2.8))
// }

// type Adder interface {
// 	int | float64
// }

// func Add[T Adder](a, b T) T {
// 	return a + b
// }
