package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 测试用例
type Case struct {
	Name     string
	Input    []int
	Excepted []int
}

// 功能测试case
var cases = []Case{
	{
		Name:     `empty`,
		Input:    []int{},
		Excepted: []int{},
	},
	{
		Name:     `[1, 4, 3, 2]`,
		Input:    []int{1, 4, 3, 2},
		Excepted: []int{1, 2, 3, 4},
	},
	{
		Name:     `[1, 4, 3, 2, 4, 3, 6, 8, 4, 2, 0, 1, 3, 5, 7, 3, 8, 17, 28, 90]`,
		Input:    []int{1, 4, 3, 2, 4, 3, 6, 8, 4, 2, 0, 1, 3, 5, 7, 3, 8, 17, 28, 90},
		Excepted: []int{0, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 17, 28, 90},
	},
}

// 基准测试case
var nums = []int{1, 4, 3, 2, 4, 3, 6, 8, 4, 2, 0, 1, 3, 5, 7, 3, 8, 17, 28, 90}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	for _, c := range cases {
		Convey("BubbleSort", t, func() {
			Convey(c.Name, func() {
				BubbleSort(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})
		})
	}
	for _, c := range cases {
		Convey("BubbleSort2", t, func() {
			Convey(c.Name, func() {
				BubbleSort2(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}
func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		BubbleSort2(numbers)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		BubbleSort(numbers)
	}
}

// 快速排序
func TestQuickSort(t *testing.T) {
	for _, c := range cases {
		Convey("BubbleSort2", t, func() {
			Convey(c.Name, func() {
				BubbleSort2(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		QuickSort(numbers)
	}
}
