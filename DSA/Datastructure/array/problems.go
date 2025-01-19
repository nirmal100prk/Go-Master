package main

// insert element in a sorted array
func InsertElement(data []int, val int) []int {
	pos := FindInsertPosition(data, val)

	data = append(data, 0)
	// shift elements to the right
	for i := len(data) - 1; i > pos; i-- {
		data[i] = data[i-1]
	}
	data[pos] = val
	return data
}

func FindInsertPosition(data []int, val int) int {
	left := 0
	right := len(data) - 1
	mid := left + (right-left)/2
	for left < right {
		if data[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
