package main

import ()

func LargestNumberTwiceOthers(nums []int) int {
	var (
		first  = -1
		second = -1
		index  = -1
	)

	for i, v := range nums {
		if v >= first {
			second = first
			first = v
			index = i
		} else if v < first && v > second {
			second = v
		}
	}
	if first >= 2*second {
		return index
	}
	return -1
}
