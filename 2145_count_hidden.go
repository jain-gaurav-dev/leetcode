package daily_challenge

import "math"

func numberOfArrays(differences []int, lower int, upper int) int {
	maxDecr, maxIncr := math.MaxInt, math.MinInt
	prefixSum := 0

	for _, d := range differences {
		prefixSum += d
		if prefixSum < maxDecr {
			maxDecr = prefixSum
		}

		if prefixSum > maxIncr {
			maxIncr = prefixSum
		}
	}

	left := lower
	if maxDecr < 0 {
		left = left - maxDecr
	}
	right := upper
	if maxIncr > 0 {
		right = right - maxIncr
	}

	numArrays := right - left + 1
	if numArrays > 0 {
		return numArrays
	}

	return 0
}
