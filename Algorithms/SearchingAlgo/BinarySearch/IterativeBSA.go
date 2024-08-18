package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(RecursiveBinarySearch(arr, 0, 9, 7))
}

/*Time Complexity:
Best Case: O(1)
Average Case: O(log N)
Worst Case: O(log N)*/

// BSA only in sorted array
func ItrBinarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == val {
			return mid
		}

		if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1

}

func RecursiveBinarySearch(arr []int, low, high, x int) int {
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			return RecursiveBinarySearch(arr, mid+1, high, x)

		} else {
			return RecursiveBinarySearch(arr, low, mid-1, x)
		}

	}
	return -1
}
