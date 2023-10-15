package main

import (
	"container/heap"
	"fmt"
)

// how to implement a heap in go
// https://golang.org/pkg/container/heap/

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func myHeap() {
	// 17, 2, 8, 27, 12, 9, 13, 25, 29, 1, 19, 20, 33, 42, 6, 11, 49, 18
	h := &IntHeap{17, 2, 8, 27, 12, 9}
	heap.Init(h)
	heap.Push(h, 13)
	heap.Push(h, 25)
	heap.Push(h, 29)
	heap.Push(h, 1)
	heap.Push(h, 19)
	heap.Push(h, 20)
	heap.Push(h, 33)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
