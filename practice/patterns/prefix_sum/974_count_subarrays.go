package prefix_sum

func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	sumModK := make(map[int]int)
	sumModK[0] = 1 % k

	prefixSum := 0
	res := 0
	for i := 0; i < n; i++ {
		prefixSum += nums[i] + k // + k is key to handling negative numbers here!

		prefixSumModK := prefixSum % k
		if cnt, ok := sumModK[prefixSumModK]; ok {
			res += cnt
		}

		sumModK[prefixSumModK] += 1
	}

	return res
}
