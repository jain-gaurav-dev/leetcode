package prefix_sum

func numberOfSubarrays(nums []int, k int) int {
	oddCounts := make(map[int]int)
	oddCounts[0] = 1

	oddCount := 0
	res := 0
	for _, num := range nums {
		if num%2 == 0 {
			// even number
			oddCounts[oddCount] += 1
		} else {
			// odd number
			oddCount += 1
			oddCounts[oddCount] += 1
		}

		if oddCount >= k {
			cnts, _ := oddCounts[oddCount-k]
			res += cnts
		}
	}

	return res
}
