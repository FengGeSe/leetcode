package main

import (
	_ "fmt"
)

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

// 插入排序
func InsertionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		temp := nums[i+1]
		for j := i; j >= 0; j-- {
			if temp >= nums[j] {
				nums[j+1] = temp
				break
			}
			if temp < nums[j] {
				nums[j+1] = nums[j]
				if j == 0 {
					nums[0] = temp
				}
			}
		}
	}
}

// 希尔排序
func ShellSort(nums []int) {
	for step := len(nums) / 2; step > 0; step /= 2 {
		for i := step; i < len(nums); i++ {
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
}

// 归并排序
func MergeSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	middle := len(nums) / 2
	left := nums[:middle]
	right := nums[middle:]
	MergeSort(left)
	MergeSort(right)
	nums = merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和rignt的index位置
	for m < len(left) && n < len(right) {
		if left[m] <= right[n] {
			result = append(result, left[m])
			m++
		} else {
			result = append(result, right[n])
			n++
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)
	return result
}

// 快速排序

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	var (
		key  = nums[0] //以第一个数为key
		low  = 0
		high = len(nums) - 1
	)
	for low < high {
		for low < high && nums[high] > key { //开始从后往前找小与key的
			high--
		}
		nums[low] = nums[high]               //找到key,交换，交换
		for low < high && nums[low] <= key { //找到key,交换一次之后，换方向从前往后遍历
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = key

	QuickSort(nums[:low])
	QuickSort(nums[low+1:])
}
