package daily_challenge

func countPairs(nums []int, target int) int {

	counts := make([]int, 101)
	for _, num := range nums {
		counts[num+50]++
	}

	for i := 1; i < 101; i++ {
		counts[i] += counts[i-1]
	}

	pairs := 0
	for _, num := range nums {
		otherMax := target - num
		if otherMax+50-1 >= 0 && otherMax+50-1 < 101 {
			pairs += counts[otherMax+50-1]

		} else if otherMax+50-1 > 100 {
			pairs += counts[100]
		}

		if num < otherMax {
			pairs -= 1
		}

	}

	return pairs / 2
}
