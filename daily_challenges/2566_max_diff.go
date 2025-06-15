package daily_challenge

import "strconv"

func minMaxDifference(num int) int {
	ns := strconv.Itoa(num)
	nsLen := len(ns)

	low := make([]byte, nsLen)
	i := 0
	msd := ns[0]
	for ; i < nsLen; i++ {
		if ns[i] == msd {
			low[i] = '0'
		} else {
			low[i] = ns[i]
		}
	}

	lowNum, _ := strconv.Atoi(string(low))

	high := make([]byte, nsLen)
	i = 0
	for ; i < nsLen && ns[i] == '9'; i++ {
		high[i] = '9'
	}

	if i < nsLen {
		non9 := ns[i]

		for ; i < nsLen; i++ {
			if ns[i] == non9 {
				high[i] = '9'
			} else {
				high[i] = ns[i]
			}
		}
	}

	highNum, _ := strconv.Atoi(string(high))
	return highNum - lowNum
}
