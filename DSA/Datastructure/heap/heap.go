package main

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {

	h.data = append(h.data, val) // add element at the end
	h.HeapifyUp(len(h.data) - 1) // restore the heap property
}

// when new element is added to heap. it is added to the end of the array.
// then it is moved upwards
// if the element is greater than parent
func (h *MinHeap) HeapifyUp(index int) {
	// there is no relation between left child and right child so no
	// need to compare with right child
	for index > 0 {
		parent := (index - 1) / 2
		// if the element is greater than or equal to parent, then heap property is maintained
		if h.data[index] >= h.data[parent] {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent

	}
}

func (h *MinHeap) Remove() {
	if len(h.data) == 0 {
		return
	}

	// remove root element
	h.data[0] = h.data[len(h.data)-1] //  swap root with last element
	h.data = h.data[:len(h.data)-1]   //  remove last element

	// restore heap property- by moving down
	h.HeapifyDown(0)
}

func (h *MinHeap) HeapifyDown(index int) {

	lastIndex := len(h.data) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		// Check left child
		if leftChild <= lastIndex && h.data[leftChild] < h.data[smallest] {
			smallest = leftChild
		}

		// Check right child
		if rightChild <= lastIndex && h.data[rightChild] < h.data[smallest] {
			smallest = rightChild
		}

		// If the smallest is still the parent, heapify is complete
		if smallest == index {
			break
		}

		// Swap with the smallest child
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}
