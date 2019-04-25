package main

import ()

func ArrayPartition(nums []int) int {
	QuickSort(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

// 快排
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	var (
		key   = nums[0]
		left  = 0
		right = len(nums) - 1
	)

	for left < right {
		for left < right && nums[right] > key {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= key {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = key
	QuickSort(nums[:left])
	QuickSort(nums[left+1:])
}
