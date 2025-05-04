package daily_challenge

type pair struct {
	a, b int
}

func numEquivDominoPairs(dominoes [][]int) int {
	dict := make(map[pair]int)

	res := 0
	for _, p := range dominoes {
		var currPair pair

		if p[0] < p[1] {
			currPair.a = p[0]
			currPair.b = p[1]
		} else {
			currPair.a = p[1]
			currPair.b = p[0]
		}

		if cnt, ok := dict[currPair]; ok {
			res += cnt
			dict[currPair] += 1
		} else {
			dict[currPair] = 1
		}
	}

	return res
}
