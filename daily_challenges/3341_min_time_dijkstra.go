package daily_challenge

import (
	"container/heap"
	"math"
)

type mypoint struct {
	r, c int
}
type nodeDist struct {
	node mypoint
	dist int
}

type NodeDistHeap []nodeDist

func (h NodeDistHeap) Len() int {
	return len(h)
}

func (h NodeDistHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h NodeDistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeDistHeap) Push(v any) {
	(*h) = append(*h, v.(nodeDist))
}

func (h *NodeDistHeap) Pop() any {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return last
}

func minTimeToReach2(moveTime [][]int) int {
	nodedistHeap := make(NodeDistHeap, 0)
	heap.Init(&nodedistHeap)

	heap.Push(&nodedistHeap, nodeDist{mypoint{0, 0}, 0})
	rows := len(moveTime)
	cols := len(moveTime[0])
	visited := 0
	totalNodes := rows * cols

	dist := make([][]int, rows)
	for r := 0; r < rows; r++ {
		dist[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			dist[r][c] = math.MaxInt
		}
	}

	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{-1, 0}, []int{0, -1}}

	for visited < totalNodes {
		popped := heap.Pop(&nodedistHeap).(nodeDist)

		if popped.dist >= dist[popped.node.r][popped.node.c] {
			continue
		}

		// get edges emanating from the popped point and add them to the heap
		r := popped.node.r
		c := popped.node.c

		dist[r][c] = popped.dist
		for _, dir := range dirs {
			nr := r + dir[0]
			nc := c + dir[1]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				candDist := moveTime[nr][nc] + 1
				if popped.dist+1 > candDist {
					candDist = popped.dist + 1
				}

				heap.Push(&nodedistHeap, nodeDist{mypoint{nr, nc}, candDist})
			}
		}
		visited++
	}

	return dist[rows-1][cols-1]
}
