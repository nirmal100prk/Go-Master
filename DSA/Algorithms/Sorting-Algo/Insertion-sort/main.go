package main

import "fmt"

func main() {
	arr := []int{6, 4, 2}
	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertionSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--

		}
	}
}

func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
