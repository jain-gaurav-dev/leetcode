package misc

import "slices"

func customSortString(order string, s string) string {
	bs := []byte(s)
	idx := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		idx[order[i]] = i
	}

	slices.SortFunc(bs, func(a, b byte) int {
		i1, ok1 := idx[a]
		i2, ok2 := idx[b]
		if ok1 && ok2 {
			if i1 < i2 {
				return -1
			} else {
				return 1
			}
		}

		if ok1 {
			return -1
		}

		if ok2 {
			return 1
		}
		
		return 0
	})

	return string(bs)
}
