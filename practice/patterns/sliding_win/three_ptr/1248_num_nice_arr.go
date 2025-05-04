package three_ptr

func numberOfSubarrays(nums []int, k int) int {
	left, mid, right := 0, 0, 0

	n := len(nums)
	oddCount := 0
	res := 0
	for ; right < n; right++ {
		if nums[right]%2 == 0 {
			// even

		} else {
			// odd number
			oddCount += 1
		}

		if oddCount > k {
			// reset left and mid to the nearest next odd number
			j := mid + 1
			for ; j < n; j++ {
				if nums[j]%2 != 0 {
					break
				}
			}
			left, mid = mid+1, j
			oddCount--
		}

		if oddCount == k {
			// move mid to an odd number if it is not at one already.
			for nums[mid]%2 == 0 && mid < n {
				mid++
			}
			// oddCount = k at this point; left and mid have been reset
			res += mid - left + 1
			//fmt.Printf("updates res to %d, left: %d, mid: %d, right: %d\n", res, left, mid, right)
		}
	}

	return res
}
