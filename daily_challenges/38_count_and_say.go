package daily_challenge

import "strconv"

func countAndSayIter(n int) string {
	if n == 1 {
		return "1"
	}

	s := "1"
	for i := 2; i <= n; i++ {
		count := 1
		res := make([]byte, 0)
		for j := 1; j < len(s); j++ {
			if s[j] == s[j-1] {
				count++
				continue
			}

			res = append(res, strconv.Itoa(count)...)
			res = append(res, s[j-1])

			count = 1
		}

		res = append(res, strconv.Itoa(count)...)
		res = append(res, s[len(s)-1])
		s = string(res)
	}

	return s
}

func countAndSay(n int) string {
	var countSay func(n int) string
	countSay = func(n int) string {
		if n == 1 {
			return "1"
		} else {
			s := countSay(n - 1)
			// x is a string
			if len(s) == 1 {
				return "1" + string(s[0])
			}

			res := make([]byte, 0)
			count := 1
			for i := 1; i < len(s); i++ {
				if s[i] == s[i-1] {
					count++
					continue
				}

				// append the current count and digit
				res = append(res, strconv.Itoa(count)...)
				res = append(res, s[i-1])

				// reset the count
				count = 1
			}

			res = append(res, strconv.Itoa(count)...)
			res = append(res, s[len(s)-1])
			return string(res)
		}
	}

	return countSay(n)
}
