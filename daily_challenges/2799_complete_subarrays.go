package daily_challenge

import "fmt"

func countCompleteSubarraysSlidingWin(nums []int) int {
	dict := make(map[int]bool)

	for _, n := range nums {
		dict[n] = true
	}

	distinct := len(dict)

	left, right := 0, -1
	n := len(nums)
	dct := make(map[int]int)
	res := 0
	for ; left < n; left++ {
		if left-1 >= 0 {
			dct[nums[left-1]] -= 1
			if dct[nums[left-1]] == 0 {
				delete(dct, nums[left-1])
			}
		}

		if len(dct) == distinct {
			res += n - right
			continue
		}

		right++
		for ; len(dct) < distinct && right < n; right++ {
			dct[nums[right]] += 1

			if len(dct) == distinct {
				res += n - right
				fmt.Printf("updated res to %d, left %d, right %d\n", res, left, right)
				break
			}
		}
	}

	return res
}

func countCompleteSubarrays(nums []int) int {
	dict := make(map[int]bool)

	for _, n := range nums {
		dict[n] = true
	}

	distinct := len(dict)

	start := distinct - 1
	n := len(nums)
	res := 0
	for i := start; i < n; i++ {
		dct := make(map[int]bool)
		dct[nums[i]] = true
		if len(dct) == distinct {
			res += 1
		}

		for j := i - 1; j >= 0 && len(dct) <= distinct; j-- {
			dct[nums[j]] = true

			if len(dct) == distinct {
				res += 1
			}
		}
	}

	return res
}
