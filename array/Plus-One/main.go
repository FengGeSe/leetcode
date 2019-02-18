package main

import ()

func PlusOne(digits []int) []int {

	flag := false
	for i := len(digits) - 1; i >= 0; i-- {
		if flag && digits[i] != 9 {
			digits[i] += 1
			flag = false
			break
		}
		if digits[i] == 9 {
			digits[i] = 0
			flag = true
			continue
		} else {
			digits[i] += 1
			break
		}

	}
	if flag {
		digits = append([]int{1}, digits...)
	}

	return digits

}
