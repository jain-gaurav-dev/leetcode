package shortestpath

import (
	"container/heap"
	"math"
)

type EdgeHeap [][]int

func (h EdgeHeap) Len() int {
	return len(h)
}

func (h EdgeHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h EdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeHeap) Push(v any) {
	(*h) = append(*h, v.([]int))
}

func (h *EdgeHeap) Pop() any {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return last
}

func networkDelayTime(times [][]int, n int, k int) int {
	edgeHeap := make(EdgeHeap, 0)
	heap.Init(&edgeHeap)

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	edges := make(map[int][][]int)
	for _, edge := range times {
		edges[edge[0]] = append(edges[edge[0]], []int{edge[1], edge[2]})
	}

	dist[k-1] = 0
	// init heap with edges from source
	srcEdges := edges[k]
	for _, edge := range srcEdges {
		heap.Push(&edgeHeap, []int{k, edge[0], edge[1]})
	}

	visited := 1
	for visited < n {
		// pick the edge which results in the smallest distance to so far unvisited node
		if edgeHeap.Len() == 0 {
			break
		}

		poppedEdge := heap.Pop(&edgeHeap).([]int)
		// if the target of edge has already been visited; continue the loop
		if dist[poppedEdge[1]-1] != math.MaxInt {
			continue
		}

		dist[poppedEdge[1]-1] = poppedEdge[2]
		// insert any edges emanating from the target to the heap
		currentDist := poppedEdge[2]
		for _, newEdge := range edges[poppedEdge[1]] {
			heap.Push(&edgeHeap, []int{poppedEdge[1], newEdge[0], currentDist + newEdge[1]})
		}
		visited++
	}

	maxDist := -1
	for i := 0; i < n; i++ {
		if dist[i] == math.MaxInt {
			return -1
		} else if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}

	return maxDist
}
