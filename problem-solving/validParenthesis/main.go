package main

import "fmt"

func main() {
	fmt.Println(ValidParenthesis("()[]{}"))
}

// Input: s = "()[]{}"
func ValidParenthesis(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, char := range s {
		if _, ok := pairs[char]; ok {
			stack = append(stack, char)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != char {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
