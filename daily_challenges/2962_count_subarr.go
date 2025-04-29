package daily_challenge

import (
	"math"
)

func countSubarraysSlidingWin(nums []int, k int) int64 {
	n := len(nums)
	maxElement := math.MinInt

	for i := 0; i < n; i++ {
		if nums[i] > maxElement {
			maxElement = nums[i]
		}
	}

	start, end := 0, 0
	maxCount := 0
	res := int64(0)
	for ; end < n; end++ {
		if nums[end] == maxElement {
			maxCount++
		}

		if maxCount >= k {
			// Shift start to the right
			s := start
			for ; s <= end; s++ {
				if nums[s] == maxElement && maxCount == k {
					// we will end up removing the maxElement from the window if we move further right
					// and the count will drop below k
					break
				} else if nums[s] == maxElement {
					maxCount--
				}
			}

			start = s
			res += int64(start + 1)
		}
	}

	return res
}

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)
	maxIdxForCnt := make(map[int]int)

	maxElem := math.MinInt
	for i := 0; i < n; i++ {
		if nums[i] > maxElem {
			maxElem = nums[i]
		}
	}

	res := int64(0)
	maxCount := 0
	maxIdxForCnt[0] = -1
	for i := 0; i < n; i++ {
		if nums[i] == maxElem {
			maxCount += 1
		}

		lookFor := maxCount - k
		if cnt, ok := maxIdxForCnt[lookFor]; ok {
			res += int64(cnt) + 1 + 1
		}

		maxIdxForCnt[maxCount] = i
	}

	return res
}
