package intervals

import (
	"container/heap"
	"slices"
)

type tripHeap [][]int

func (h tripHeap) Len() int {
	return len(h)
}

func (h tripHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h tripHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *tripHeap) Push(v any) {
	(*h) = append(*h, v.([]int))
}

func (h *tripHeap) Pop() any {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[0 : len(*h)-1]
	return last
}

func carPooling(trips [][]int, capacity int) bool {
	slices.SortFunc(trips, func(a, b []int) int {
		if a[1] < b[1] {
			return -1
		} else if a[1] == b[1] {
			return 0
		} else {
			return 1
		}
	})

	n := len(trips)
	remCap := capacity
	trpHeap := make(tripHeap, 0)
	heap.Init(&trpHeap)

	for i := 0; i < n; i++ {
		pas, start := trips[i][0], trips[i][1]
		for trpHeap.Len() > 0 {
			firstEndingTrip := trpHeap[0]

			if firstEndingTrip[2] <= start {
				heap.Pop(&trpHeap)
				remCap += firstEndingTrip[0]
			} else {
				break
			}
		}

		remCap -= pas
		if remCap < 0 {
			return false
		}

		// insert trip into the heap
		heap.Push(&trpHeap, trips[i])
	}

	return true
}
