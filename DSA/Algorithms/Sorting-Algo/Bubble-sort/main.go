package main

import "fmt"

func main() {
	fmt.Println(BubbleSort([]int{6, 8, 2, 7, 2, 6, 4, 1}))
}

// swapping with next element in the j for loop
// largest element get 
func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
