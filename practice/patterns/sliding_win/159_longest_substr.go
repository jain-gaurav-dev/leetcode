package sliding_win

func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)

	left := 0
	freq := make(map[byte]int)
	maxWin := 0
	for right := 0; right < n; right++ {
		// insert char at right
		freq[s[right]] += 1

		if len(freq) <= 2 {
			if right-left+1 > maxWin {
				maxWin = right - left + 1
			}
			continue
		}

		// shrink from left
		for len(freq) > 2 {
			freq[s[left]] -= 1
			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}

			left++
		}
	}
	return maxWin
}
