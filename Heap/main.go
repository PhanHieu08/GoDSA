// max heap - parrent > child
package main

import "fmt"

type MaxHeap struct {
	array []int // slices
}

//method insert new value to the end of heap
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.maxHeapifyUp(len(h.array) - 1)
}

// extract root of heap
func (h *MaxHeap) Extract() int {

	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	extracted := h.array[0]

	// last index
	last := len(h.array) - 1

	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

// index is where the heapify start - current index
func (h *MaxHeap) maxHeapifyUp(current int) {
	for h.array[getParentIndex(current)] < h.array[current] {
		h.swap(getParentIndex(current), current)
		// update current index
		current = getParentIndex(current)
	}
}

func (h *MaxHeap) maxHeapifyDown(current int) {
	last := len(h.array) - 1
	l, r := getLeftIndex(current), getRightIndex(current)

	childToCompare := 0

	for l <= last {
		if l == last || h.array[l] > h.array[r] { // 1 child - left child || left > right
			childToCompare = l
		} else {
			childToCompare = r
		}

		// if current smaller than child to compare then swap and
		if h.array[current] < h.array[childToCompare] {
			h.swap(current, childToCompare)

			current = childToCompare
			l, r = getLeftIndex(current), getRightIndex(current)
		} else {
			return
		}
	}
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getLeftIndex(i int) int {
	return 2*i + 1
}

func getRightIndex(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 6, 44, 21, 34, 3}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	fmt.Println("Extract: ")
	for i := 0; i < 5; i++ {
		fmt.Println(m.Extract())
		fmt.Println(m)
	}
}
