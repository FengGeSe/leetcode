package main

import (
//	"fmt"
)

func FindMaxConsecutiveOnes(nums []int) int {

	var (
		n       = 0
		p       = -1
		q       = 0
		isCount = false
	)

	for ; q < len(nums); q++ {
		if nums[q] == 1 && !isCount {
			isCount = true
			p = q
		}

		if isCount && (nums[q] == 0 || q == len(nums)-1) {
			if q == len(nums)-1 && nums[q] == 1 {
				q++
			}
			if n < q-p {
				n = q - p
			}

			p = q
			isCount = false
		}

	}

	return n

}
