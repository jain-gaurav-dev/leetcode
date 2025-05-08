package misc

func findCommonResponse(responses [][]string) string {
	freq := make(map[string]int)

	for _, resps := range responses {
		uniqResp := make(map[string]bool)

		for _, resp := range resps {
			uniqResp[resp] = true
		}

		for r := range uniqResp {
			freq[r] += 1
		}
	}

	mostCommon := ""
	mostFreq := -1
	for r, f := range freq {
		if f > mostFreq {
			mostCommon = r
			mostFreq = f
		} else if f == mostFreq {
			if r < mostCommon {
				mostCommon = r
			}
		}
	}

	return mostCommon
}
