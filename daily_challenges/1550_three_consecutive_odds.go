package daily_challenge

func threeConsecutiveOdds(arr []int) bool {
	odds := 0
	for _, k := range arr {
		if k&0x1 == 1 {
			odds++
		} else {
			odds = 0
		}

		if odds == 3 {
			return true
		}
	}

	return false
}
