package daily_challenge

import "strconv"

func maxDiff(num int) int {
	ns := strconv.Itoa(num)
	numLen := len(ns)

	low := make([]byte, numLen)
	i := 0
	for ; i < numLen && (ns[i] == '1' || ns[i] == '0'); i++ {
		low[i] = ns[i]
	}

	if i < numLen {
		non1OrZero := ns[i]
		replaceWith := byte('1')
		if i > 0 {
			replaceWith = '0'
		}

		for ; i < numLen; i++ {
			if ns[i] == non1OrZero {
				low[i] = replaceWith
			} else {
				low[i] = ns[i]
			}
		}
	}

	lowNum, _ := strconv.Atoi(string(low))

	high := make([]byte, numLen)
	i = 0
	for ; i < numLen && ns[i] == '9'; i++ {
		high[i] = '9'
	}

	if i < numLen {
		non9 := ns[i]
		for ; i < numLen; i++ {
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
