package misc

func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	res := make([][]int, 0)

	idx1, idx2 := 0, 0
	m, n := len(encoded1), len(encoded2)

	count1 := 0
	count2 := 0

	for i := 0; idx1 < m && idx2 < n; {
		r1 := i - count1
		r2 := i - count2

		if r1 == encoded1[idx1][1] {
			count1 += encoded1[idx1][1]
			idx1++
			continue
		}
		if r2 == encoded2[idx2][1] {
			count2 += encoded2[idx2][1]
			idx2++
			continue
		}

		prod := encoded1[idx1][0] * encoded2[idx2][0]
		if len(res) == 0 {

			res = append(res, []int{prod, 1})
			// fmt.Printf("Append first: res is %v, i %d, count1 %d, count2 %d, idx1 %d, idx2 %d\n", res, i, count1, count2, idx1, idx2)
		} else {
			last := res[len(res)-1]
			if prod == last[0] {
				res[len(res)-1][1]++
				//fmt.Printf("Incr: res is %v, i %d, count1 %d, count2 %d, idx1 %d, idx2 %d\n", res, i, count1, count2, idx1, idx2)
			} else {
				res = append(res, []int{prod, 1})
				//fmt.Printf("Append: res is %v, i %d, count1 %d, count2 %d, idx1 %d, idx2 %d\n", res, i, count1, count2, idx1, idx2)

			}
		}
		i++
	}

	return res
}
