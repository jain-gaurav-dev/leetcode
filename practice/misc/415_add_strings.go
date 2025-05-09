package misc

func addStrings(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)

	res := make([]byte, 0, max(m, n))
	carry := uint8(0)
	for i, j := m-1, n-1; i >= 0 || j >= 0; {
		dig1 := uint8(0)
		dig2 := uint8(0)

		if i >= 0 {
			dig1 = num1[i] - '0'
		}
		if j >= 0 {
			dig2 = num2[j] - '0'
		}

		sum := dig1 + dig2 + carry
		sumDig := sum % 10
		carry = sum / 10
		res = append(res, sumDig+'0')

		i--
		j--
	}

	if carry == uint8(1) {
		res = append(res, carry+'0')
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]

		i++
		j--
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
