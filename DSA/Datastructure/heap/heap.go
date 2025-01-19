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

	if index < 0 || index >= len(h.data) {
		return
	}
	for index > 0 {
		parent := (index - 1) / 2
		// if the element is greater than or equal to parent, then heap property is maintained
		if h.data[index] >= h.data[parent] {
			break
		}
		// swap with parent
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

// find median from a data stream / Median in an infinite series of integers
// Median is the middle value in a sorted integer list
// solved using heaps, running integers are added to max and min heaps, then
// median is found from the heap
type MaxHeap struct {
	data []int
}

func (m *MaxHeap) Push(data int) {
	m.data = append(m.data, data)
	m.HeapifyUp()
}

func (m *MaxHeap) HeapifyUp() {
	if len(m.data) == 0 {
		return
	}
	index := len(m.data) - 1
	for index > 0 {
		parent := (index - 1) / 2
		if m.data[index] > m.data[parent] {
			break
		}
		m.data[index], m.data[parent] = m.data[parent], m.data[index]
		index = parent
	}
}

func (m *MaxHeap) ExtractMax() int {
	if len(m.data) == 0 {
		return 0
	}
	max := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.HeapifyDown(0)
	return max
}

func (m *MaxHeap) HeapifyDown(index int) {
	if len(m.data) == 0 {
		return
	}

	lastIndex := len(m.data) - 1

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		largest := index

		if leftChild <= lastIndex && m.data[leftChild] > m.data[largest] {
			largest = leftChild
		}

		if rightChild <= lastIndex && m.data[rightChild] > m.data[largest] {
			largest = rightChild
		}
		if largest == index {
			break
		}
		m.data[index], m.data[largest] = m.data[largest], m.data[index]
		index = largest
	}

}

