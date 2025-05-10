package daily_challenge

func minSum(nums1 []int, nums2 []int) int64 {
	m := len(nums1)
	n := len(nums2)

	zeroes1, zeroes2 := 0, 0
	sum1, sum2 := 0, 0
	for i := 0; i < m; i++ {
		if nums1[i] == 0 {
			zeroes1++
		} else {
			sum1 += nums1[i]
		}
	}

	for j := 0; j < n; j++ {
		if nums2[j] == 0 {
			zeroes2++
		} else {
			sum2 += nums2[j]
		}
	}

	if zeroes1 == 0 && zeroes2 == 0 && sum1 != sum2 {
		return -1
	}

	if sum1 > sum2 {
		sum1, sum2 = sum2, sum1
		zeroes1, zeroes2 = zeroes2, zeroes1
	}

	if zeroes2 == 0 {
		targetSum := sum2
		if sum1+zeroes1 > targetSum {
			return -1
		}
	}

	if zeroes1 == 0 {
		targetSum := sum1
		if sum2+zeroes2 > targetSum {
			return -1
		}
	}

	if sum2+zeroes2 > sum1+zeroes1 {
		return int64(sum2 + zeroes2)
	} else {
		return int64(sum1 + zeroes1)
	}

}
