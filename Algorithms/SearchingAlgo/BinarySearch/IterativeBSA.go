package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(RecursiveBinarySearch(arr, 0, len(arr)-1, 2))
}

/*Time Complexity:
Best Case: O(1)
Average Case: O(log N)
Worst Case: O(log N)*/

// BSA only in sorted array
func IterativeBinarySearch(arr []int, low, high, x int) int {
	for low <= high {
		mid := (high - low) / 2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			low = mid + 1

		} else {
			high = mid - 1
		}

	}
	return -1
}

func RecursiveBinarySearch(arr []int, low, high, x int) int {
	for low <= high {
		mid := (high - low) / 2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			return RecursiveBinarySearch(arr, low, mid+1, x)

		} else {
			return RecursiveBinarySearch(arr, low, mid-1, x)
		}

	}
	return -1
}
