package daily_challenge

import "fmt"

func pushDominoes(dominoes string) string {
	domino := make([]byte, len(dominoes)+2)
	domino[0] = 'L'
	domino[len(domino)-1] = 'R'
	copy(domino[1:], dominoes)

	lastChar := byte('L')
	lastIdx := 0
	for i := 1; i < len(domino); i++ {
		if domino[i] == '.' {
			continue
		}

		if domino[i] == lastChar {
			// set all indices from lastIdx + 1 to i to lastChar
			for j := lastIdx + 1; j < i; j++ {
				domino[j] = lastChar
				fmt.Printf("1: set byte at pos %d to %s; string is %s\n", j, string(lastChar), domino)
			}

			lastIdx = i
		} else if lastChar == byte('L') && domino[i] == byte('R') {
			// only update the lastchar and lastidx
			lastChar = byte('R')
			lastIdx = i
		} else {
			// lastChar = R and domino[i] = L
			for p, q := lastIdx, i; p < q; {
				domino[p] = 'R'
				fmt.Printf("set byte at pos %d to %s; string is %s\n", p, "R", domino)

				domino[q] = 'L'
				fmt.Printf("set byte at pos %d to %s; string is %s\n", q, "L", domino)

				p++
				q--
			}

			lastChar = 'L'
			lastIdx = i
		}
	}

	return string(domino[1 : len(domino)-1])
}
