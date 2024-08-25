package main

import "fmt"

func main() {
	fib := nextFibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(fib())
	}
}

func nextFibonacci() func() int {

	var fib0, fib1 int = 0, 1

	return func() int {
		next := fib0
		fib0, fib1 = fib1, fib0+fib1
		return next

	}
}
