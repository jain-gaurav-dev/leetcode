package daily_challenge

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)
	var binGreaterEqual func(ge int) int
	binGreaterEqual = func(ge int) int {
		i, j := 0, n-1

		for i < j {
			mid := i + (j-i)/2
			if nums[mid] >= ge {
				j = mid
			} else if nums[mid] < ge {
				i = mid + 1
			}
		}

		if nums[i] >= ge {
			return i
		} else {
			return n
		}
	}

	var binLessEqual func(le int) int
	binLessEqual = func(le int) int {
		i, j := 0, n-1

		for i < j {
			mid := j - (j-i)/2
			if nums[mid] <= le {
				i = mid
			} else if nums[mid] > le {
				j = mid - 1
			}
		}

		if nums[i] <= le {
			return i
		} else {
			return -1
		}
	}

	fairPairs := int64(0)
	for i, num := range nums {
		ge := lower - num
		le := upper - num

		left := binGreaterEqual(ge)
		right := binLessEqual(le)
		if left == n || right == -1 {
			continue
		}
		pairs := right - left + 1
		if i >= left && i <= right {
			pairs -= 1
		}

		fairPairs += int64(pairs)
	}

	return fairPairs / 2
}
