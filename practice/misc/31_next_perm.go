package misc

import "math"

func nextPermutation(nums []int) {
	n := len(nums)

	prev := -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] >= prev {
			prev = nums[i]
		} else {
			// bring the num >= nums[i] at ith position by swap
			next := math.MaxInt
			nextIdx := -1
			pivot := nums[i]
			for j := i + 1; j < n; j++ {
				if nums[j] > pivot && nums[j] < next {
					nextIdx = j
					next = nums[j]
				}
			}

			// swap
			nums[i], nums[nextIdx] = nums[nextIdx], nums[i]

			// ~ to insertion sort for the rest of the array
			for j := i + 1; j < n; j++ {
				for k := j + 1; k < n; k++ {
					if nums[k] < nums[j] {
						nums[k], nums[j] = nums[j], nums[k]
					}
				}
			}

			return
		}
	}

	// reverse the array, if we are here
	for left, right := 0, n-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
