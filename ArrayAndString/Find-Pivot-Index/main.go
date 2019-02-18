package main

import ()

func FindPivotIndex(nums []int) int {
	var (
		leftSum  = 0
		rightSum = 0
	)
	for _, v := range nums {
		rightSum += v
	}

	for i, v := range nums {
		rightSum -= v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}

	return -1

}
