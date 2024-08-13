package main

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop(val int) {
	*s = (*s)[:len(*s)-1]
}
