package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{3, 6, 8, 4, 3, 2, 8, 9, 3, 2, 9, 6}))
}

//recursive
// splits and sorts the subarrays
func mergeSort(arr []int) []int {
	//base case
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

// merges two sorted subarrays
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
