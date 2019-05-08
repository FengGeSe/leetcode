package main

import ()

func RemoveElement(nums []int, val int) int {

	var (
		header = 0
		tail   = len(nums) - 1
	)

	for {
		for tail > 0 && nums[tail] == val {
			tail--
		}
		for header <= tail && nums[header] != val {
			header++
		}

		if header < tail {
			nums[header], nums[tail] = nums[tail], nums[header]
		} else {
			break
		}
	}
	return header
}
