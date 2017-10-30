package modules
/* Priority Queue Implementation Using Golang container/heap */

import (
    "container/heap"
)

/* A Min Priority Queue using golang container/heap */
/* =========================================================================== */
type MinHeap []int

/* Min Heap Functions */
func (h MinHeap) Len() int              { return len(h) }
func (h MinHeap) Less(i, j int) bool    { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    // method push data to min heap
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    // method to pop data from min heap
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
/* =========================================================================== */


/* A Max Priority Queue using golang container/heap */
/* =========================================================================== */
type MaxHeap []int

/* Min Heap Functions */
func (h MaxHeap) Len() int              { return len(h) }
func (h MaxHeap) Less(i, j int) bool    { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    // method push data to min heap
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    // method to pop data from min heap
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
/* =========================================================================== */

/* Priority Queue Functions */
/* =========================================================================== */

/* function to init min heap */
func MakeMinHeap() *MinHeap {
    minheap := new(MinHeap)
    heap.Init(minheap)
    return minheap
}


/* function to init max heap */
func MakeMaxHeap() *MaxHeap {
    maxheap := new(MaxHeap)
    heap.Init(maxheap)
    return maxheap
}


/* function to push to a maxheap */
func PushMaxHeap(h *MaxHeap, x int) {
    heap.Push(h, x)
}


/* function to push to a minheap */
func PushMinHeap(h *MinHeap, x int) {
    heap.Push(h, x)
}


/* function to push to a maxheap */
func PopMaxHeap(h *MaxHeap) interface{}{
    return heap.Pop(h)
}


/* function to push to a minheap */
func PopMinHeap(h *MinHeap) interface{}{
    return heap.Pop(h)
}
/* =========================================================================== */
