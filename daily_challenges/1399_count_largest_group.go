package daily_challenge

func countLargestGroup(n int) int {
	var digitSum func(k int, sum int) int
	digitSum = func(k int, sum int) int {
		if k == 0 {
			return sum
		}

		dig := k % 10
		return digitSum(k/10, sum+dig)
	}

	groupsBySum := make(map[int]int)

	for i := 1; i <= n; i++ {
		digSum := digitSum(i, 0)
		groupsBySum[digSum] += 1
	}

	res := 0
	maxGroupSize := 0
	for _, groupSize := range groupsBySum {
		if groupSize > maxGroupSize {
			maxGroupSize = groupSize
			res = 1
		} else if groupSize == maxGroupSize {
			res += 1
		}
	}

	return res
}
