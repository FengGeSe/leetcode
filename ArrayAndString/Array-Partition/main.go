package main

import ()

func ArrayPartition(nums []int) int {
	var (
		sum  = 0
		flag = false
	)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
		if flag {
			sum += nums[len(nums)-i-1]
		}
		flag = !flag
	}
	return sum
}
