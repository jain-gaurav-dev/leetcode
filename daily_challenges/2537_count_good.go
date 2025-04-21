package daily_challenge

func countGood(nums []int, k int) int64 {
	n := len(nums)
	if n == 1 {
		return 0
	}

	freq := make(map[int]int)
	left, right := 0, 0

	pairs := 0
	freq[nums[0]] = 1
	goodSubArrs := int64(0)

	for right <= n-1 {
		right += 1
		for ; pairs < k && right < n; right++ {
			pairs += freq[nums[right]]
			freq[nums[right]] += 1
			// fmt.Printf("Moving right; left %d, right %d, pairs %d\n", left, right, pairs)
			if pairs >= k {
				break
			}

		}
		// pairs >= k at this point OR right == n
		if right == n {
			// fmt.Printf("returning %d, right = %d, n = %d\n", goodSubArrs, right, n)
			return goodSubArrs
		}

		goodSubArrs += int64(n - right)
		// fmt.Printf("Updated goodSubArrs to %d, n %d, right %d\n", goodSubArrs, n, right)
		// Attempt to move left rightwards; see how far you can go
		left += 1
		for ; pairs >= k && left < right; left++ {
			freq[nums[left-1]] -= 1
			pairs -= freq[nums[left-1]]

			// fmt.Printf("Moving left; left %d, right %d, pairs %d\n", left, right, pairs)
			if pairs >= k {
				goodSubArrs += int64(n - right)
			} else {
				break
			}
		}
	}

	return goodSubArrs
}
