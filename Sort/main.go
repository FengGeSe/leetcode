package main

import ()

// 冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 冒泡排序-改
func BubbleSort2(nums []int) {
	var flag = false
	for !flag {
		flag = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				flag = false
			}
		}
	}
}

// 选择排序
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		p := i
		for j := i + 2; j < len(nums)-1; j++ {
			if nums[j] < nums[p] {
				p = j
			}
		}
		nums[i], nums[p] = nums[p], nums[i]
	}

}

// 快速排序
func QuickSort(nums []int) {

}
