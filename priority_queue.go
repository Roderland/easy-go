package easy_go

import (
	"fmt"
)

/*
	Always reserve the first slot
	-----------------------------
    |   | A | B | C | D |   |   |
	-----------------------------
*/

type PriorityQueue struct {
	array []Interface
	size  int
}

func NewPriorityQueueWithSlice(slice []Interface) *PriorityQueue {
	pq := NewPriorityQueue()
	for _, e := range slice {
		pq.Push(e)
	}
	return pq
}

func NewPriorityQueue() *PriorityQueue {
	return NewPriorityQueueWithCapacity(0)
}

func NewPriorityQueueWithCapacity(cap int) *PriorityQueue {
	array := make([]Interface, cap+1)
	return &PriorityQueue{
		array: array,
		size:  0,
	}
}

func (pq *PriorityQueue) Size() int {
	return pq.size
}

func (pq *PriorityQueue) Push(x Interface) {
	pq.size++
	pq.array = append(pq.array, x)
	pq.up()
}

func (pq *PriorityQueue) Peek() Interface {
	array := pq.array
	return array[1]
}

func (pq *PriorityQueue) Pop() Interface {
	array := pq.array
	array[0] = array[1]
	array[1] = array[pq.size]
	pq.size--
	pq.down()
	return array[0]
}

func (pq *PriorityQueue) ElementZero() Interface {
	array := pq.array
	return array[0]
}

func (pq *PriorityQueue) PrintInfo() {
	array := pq.array
	fmt.Printf("size of heap: %d, len of array: %d, cap of array: %d, element in first slot: %d\n",
		pq.size, len(array), cap(array), array[0])
	fmt.Println(array[1:])
}

func (pq *PriorityQueue) up() {
	array := pq.array
	tmp := array[pq.size]
	child := pq.size
	parent := child / 2
	array[0] = tmp
	for tmp.Less(array[parent]) {
		array[child] = array[parent]
		child = parent
		parent = child / 2
	}
	array[child] = tmp
}

func (pq *PriorityQueue) down() {
	array := pq.array
	tmp := array[1]
	parent := 1
	left := parent * 2
	for left <= pq.size {
		child := left
		right := left + 1
		if right <= pq.size && array[right].Less(array[left]) {
			child = right
		}
		if array[child].Less(tmp) {
			array[parent] = array[child]
		} else {
			break
		}
		parent = child
		left = parent * 2
	}
	array[parent] = tmp
}
