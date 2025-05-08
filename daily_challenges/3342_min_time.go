package daily_challenge

import (
	"container/heap"
	"math"
)

type pnt struct {
	r, c int
}
type myNodeDist struct {
	node   pnt
	dst    int
	weight int
}

type NodeHeap []myNodeDist

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].dst < h[j].dst
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(v any) {
	(*h) = append(*h, v.(myNodeDist))
}

func (h *NodeHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return last
}

func minTimeToReach3(moveTime [][]int) int {
	rows := len(moveTime)
	cols := len(moveTime[0])
	dist := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt
		}
	}

	hp := make(NodeHeap, 0)
	heap.Init(&hp)

	src := myNodeDist{pnt{0, 0}, 0, 2}
	heap.Push(&hp, src)
	dist[0][0] = 0

	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

	for {
		popped := heap.Pop(&hp).(myNodeDist)
		r, c := popped.node.r, popped.node.c

		dist[r][c] = popped.dst
		if r == rows-1 && c == cols-1 {
			return popped.dst
		}

		// add edges from the popped node to the heap
		newWeight := 1
		if popped.weight == 1 {
			newWeight = 2
		}

		for _, dir := range dirs {
			newR := r + dir[0]
			newC := c + dir[1]

			if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
				dst := moveTime[newR][newC] + newWeight
				if popped.dst+newWeight > dst {
					dst = popped.dst + newWeight
				}

				// update the distance estimate as soon as we find the corresponding edge
				// and only if the new distance estimate is smaller, add that edge.
				if dst < dist[newR][newC] {
					dist[newR][newC] = dst
					heap.Push(&hp, myNodeDist{pnt{newR, newC}, dst, newWeight})
				}
			}
		}
	}

	return -1
}
