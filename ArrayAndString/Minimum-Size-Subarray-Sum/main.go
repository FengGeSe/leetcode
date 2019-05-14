package main

import (
//	"fmt"
)

func MinSubArrayLen(s int, nums []int) int {
	min := 0
	l := 0
	r := 0
	length := len(nums)

	// The sum of all numbers in current sliding window
	sum := 0
	for r < length {
		sum += nums[r]

		if sum >= s && min == 0 {
			min = r - l + 1
		}

		for sum > s && l < r {
			sum -= nums[l]
			l++
			if sum < s {
				break
			}

			if min > r-l+1 {
				min = r - l + 1
			}
		}

		r++
	}

	return min

}
