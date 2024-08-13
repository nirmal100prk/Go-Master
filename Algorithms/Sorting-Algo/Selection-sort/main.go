package main

import "fmt"

func main() {
	fmt.Println(SelectionSort([]int{5, 4, 3, 2, 1}))
}

// sorting in ascending order
// finding the index of minimum value in the unsorted list and swapping it with the current element
// swapping the ith element with the smallest element
// sorting happens from left to right

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
