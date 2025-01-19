package main

func main() {

}

type stack []rune

func (s *stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *stack) Pop() rune {
	lastIn := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return lastIn
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// how stacks can be used for balancing of symbols

var symbolsMap = map[rune]rune{'{': '}', '(': ')', '[': ']'}

func BalancingSymbols(symbols string) bool {
	stc := &stack{}
	for _, char := range symbols {
		if _, ok := symbolsMap[char]; ok {
			stc.Push(char)
		} else if stc.IsEmpty() || symbolsMap[stc.Pop()] != char {
			return false
		} else {
			stc.Pop()
		}
	}
	return len(*stc) == 0
}

// Infix to postfix conversion algorithm using stack
// infix = operator is placed between operands
// human readable mathematics and programming languages
// prefix = operator is placed before operands
// used in programming languages and compilers
// postfix = operator is placed after the operands
// used in calculators and stack based  computations

// Reverse a stack using only push and pop

// Visualization of the Recursive Flow
// ReverseStack([1, 2, 3, 4])
//     Pop 4 → ReverseStack([1, 2, 3])
//         Pop 3 → ReverseStack([1, 2])
//             Pop 2 → ReverseStack([1])
//                 Pop 1 → ReverseStack([])
//                     Base case reached
//                 insertAtBottom(1) → [1]
//             insertAtBottom(2) → [2, 1]
//         insertAtBottom(3) → [3, 2, 1]
//     insertAtBottom(4) → [4, 3, 2, 1]

type intStack struct {
	data []int
}

func (s *intStack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *intStack) Pop() int {
	n := len(s.data)
	top := s.data[n-1]
	s.data = s.data[:n-1]
	return top
}

func (s *intStack) IsEmpty() bool {

	return len(s.data) == 0
}

func (s *intStack) ReverseStack() {

	if s.IsEmpty() {
		return
	}

	// eg: [1,2,3,4] where 4 is the top element and 1 is the bottom element
	top := s.Pop()

	// ReverseStack will pop the element recursively until the base case is reached
	s.ReverseStack()

	// the first element to be passed to insertAtBottom will be 1, then 2,3,4
	s.insertAtBottom(top)
}

func (s *intStack) insertAtBottom(val int) {
	if s.IsEmpty() {
		s.Push(val)
		return
	}

	// insertAtBottom receives the values in the order 1,2,3,4
	// this has to be reversed
	// so before pushing, pop the existing elements then push the incoming element,
	// after that append the already existing elements. this is done recursively.
	top := s.Pop()
	s.insertAtBottom(val)
	s.Push(top)
}



// 