package misc

import "fmt"

func findKthPositive(arr []int, k int) int {
	missing := 0
	prev := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == prev+1 {
			prev = arr[i]
			continue
		}

		missing += arr[i] - prev - 1
		if missing >= k {
			missingSoFar := missing - (arr[i] - prev - 1)
			for j := prev + 1; j < arr[i]; j++ {
				missingSoFar += 1
				if missingSoFar == k {
					return j
				}
			}
		}

		prev = arr[i]
	}

	if missing < k {
		return arr[n-1] + k - missing
	}

	return -1
}

func findKthPositiveBinSearch(arr []int, k int) int {
	// use bin search of the smallest idx in arr which has >= k missing elemets
	n := len(arr)
	low := 0
	high := n - 1
	for low < high {
		mid := low + (high-low)/2

		actualCount := mid + 1
		missing := arr[mid] - actualCount
		if missing >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	fmt.Printf("low is %d\n", low)
	if low == n-1 {
		missingSoFar := arr[n-1] - len(arr)
		if missingSoFar < k {
			return arr[n-1] + k - missingSoFar
		}
	}

	if low == 0 {
		return k
	}

	prev := low - 1
	missingTillPrev := arr[prev] - (prev + 1)
	fmt.Printf("prev %d, missingTillPrev %d ans %d\n", prev, missingTillPrev, arr[prev]+k-missingTillPrev)
	return arr[prev] + k - missingTillPrev
}
