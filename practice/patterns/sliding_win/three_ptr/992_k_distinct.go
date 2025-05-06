package three_ptr

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	n := len(nums)

	freq := make(map[int]int)
	res := 0
	for left, mid, right := 0, 0, 0; right < n; right++ {
		freq[nums[right]] += 1

		if len(freq) < k {
			continue
		}

		// len(freq) > k // more than k distinct integers in [mid, right]
		// k distinct integers in [mid, right] at this point; This does not guarantee that
		// [mid, right] is the smallest window with k distinct integers so far.
		// Consider [1,1,1,2,2,3] and k = 2
		// When mid moves to first 2, and right is at 3, it is not the smallest window with k distinct integers.
		// The loop for shortening the window must come after the loop below to keep number of distinct integers
		// <= k
		winWasGreaterThanK := len(freq) > k
		for ; len(freq) > k && mid < n; mid++ {
			freq[nums[mid]] -= 1
			if freq[nums[mid]] == 0 {
				delete(freq, nums[mid])
			}
			fmt.Printf("1: len(freq): %d left %d mid %d right %d\n", len(freq), left, mid, right)

		}
		if winWasGreaterThanK {
			left = mid
		}

		for ; freq[nums[mid]] > 1 && mid < n; mid++ {
			// advance mid until [mid, right] is the smallest window with k distinct integers
			freq[nums[mid]] -= 1
			// fmt.Printf("2: len(freq): %d left %d mid %d right %d\n", len(freq), left, mid, right)

		}

		res += mid - left + 1
		// fmt.Printf("3: len(freq): %d left %d mid %d right %d res %d\n", len(freq), left, mid, right, res)

	}

	return res
}
