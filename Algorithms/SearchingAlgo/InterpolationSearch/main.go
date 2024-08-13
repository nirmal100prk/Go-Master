package main

import "fmt"

func interpolation_search(arr []int, target, n int) int {
	low := 0
	high := n - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		// Calculate the position of the target element based on its value
		pos := low + (((target - arr[low]) * (high - low)) / (arr[high] - arr[low]))

		// Check if the target element is at the calculated position
		if arr[pos] == target {
			return pos
		}

		// If the target element is less than the element at the calculated position, search the left half of the list
		if arr[pos] > target {
			high = pos - 1
		} else {
			// If the target element is greater than the element at the calculated position, search the right half of the list
			low = pos + 1
		}
	}
	return -1
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(arr)
	target := 5
	index := interpolation_search(arr, target, n)

	fmt.Println(index) // Output: 4
}
