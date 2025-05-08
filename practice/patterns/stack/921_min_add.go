package stack

func minAddToMakeValid(s string) int {
	n := len(s)
	openingPar := 0
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			openingPar++
		} else {
			if openingPar > 0 {
				openingPar--
			} else {
				res++
			}
		}
	}

	res += openingPar
	return res
}
