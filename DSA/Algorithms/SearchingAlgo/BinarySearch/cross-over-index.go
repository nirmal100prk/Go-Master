package main

/*
	x0<x1<...<xn
	y0<y1<....<yn
	y0<x0, yn>xn
	there is an index i at which cross over happens
	find that index
	hint: use binary search
*/
func CrossOverIndex(x, y []int) int {
	low, high := 0, len(x)-1

	for low <= high {
		mid := low + (high-low)/2

		if y[mid] < x[mid] && y[mid+1] > x[mid+1] {
			return mid
		}

		if y[mid] > x[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
