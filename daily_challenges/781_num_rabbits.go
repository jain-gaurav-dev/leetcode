package daily_challenge

func numRabbits(answers []int) int {
	colored := make(map[int]int)
	minRabbits := 0
	for _, ans := range answers {
		if ans == 0 {
			minRabbits += 1
		} else if seen, ok := colored[ans+1]; ok {
			if seen == ans+1 {
				// this has to be a rabbit of a different color
				minRabbits += (ans + 1)
				colored[ans+1] = 1
			} else {
				colored[ans+1] += 1
			}
		} else {
			colored[ans+1] = 1
		}
	}

	for n := range colored {
		minRabbits += n
	}

	return minRabbits
}
