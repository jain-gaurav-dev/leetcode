package binsearch

func missingElement(nums []int, k int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := high - (high-low)/2

		total := nums[mid] - nums[0] + 1
		actual := mid + 1
		missing := total - actual
		if missing < k {
			low = mid
		} else if missing >= k {
			high = mid - 1
		}
	}

	missingTillLow := nums[low] - nums[0] - low
	return nums[low] + (k - missingTillLow)
}
