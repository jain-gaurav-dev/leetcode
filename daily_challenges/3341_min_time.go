package daily_challenge

import "math"

type point struct {
	r, c int
}

func minTimeToReach(moveTime [][]int) int {
	rows := len(moveTime)
	cols := len(moveTime[0])

	visited := make(map[point]int)
	visited[point{0, 0}] = 0
	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{-1, 0}, []int{0, -1}}

	minTs := math.MaxInt
	var bt func(r, c, ts int)
	bt = func(r, c, ts int) {
		if r == rows-1 && c == cols-1 {
			if ts < minTs {
				minTs = ts
			}
			return
		}

		for _, dir := range dirs {
			neigh := point{r + dir[0], c + dir[1]}
			if neigh.r < rows && neigh.r >= 0 && neigh.c < cols && neigh.c >= 0 {
				candTs := moveTime[neigh.r][neigh.c] + 1
				if ts+1 > candTs {
					candTs = ts + 1
				}

				currentTs, ok := visited[neigh]
				if !ok {
					// visit the neighbor point
					visited[neigh] = candTs
					bt(neigh.r, neigh.c, candTs)
				} else {
					// neigh point has been visited previously, see if this attempt is with a smaller timestamp
					if candTs < currentTs {
						visited[neigh] = candTs
						bt(neigh.r, neigh.c, candTs)
					}
				}
			}
		}
	}

	bt(0, 0, 0)
	return minTs
}
