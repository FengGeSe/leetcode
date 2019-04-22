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
		Name:     `[1]`,
		Input:    []int{1},
		Excepted: []int{1},
	},
	{
		Name:     `[1, 4, 3, 2]`,
		Input:    []int{1, 4, 3, 2},
		Excepted: []int{1, 2, 3, 4},
	},
	{
		Name:     `[1, 4, 3, 2, 0]`,
		Input:    []int{1, 4, 3, 2, 0},
		Excepted: []int{0, 1, 2, 3, 4},
	},
	{
		Name:     `[3, 4, 1, 2, 0]`,
		Input:    []int{3, 4, 1, 2, 0},
		Excepted: []int{0, 1, 2, 3, 4},
	},
	{
		Name:     `[1, 4, 3, 2, 4, 3, 6, 8, 4, 2, 0, 1, 3, 5, 7, 3, 8, 17, 28, 90]`,
		Input:    []int{1, 4, 3, 2, 4, 3, 6, 8, 4, 2, 0, 1, 3, 5, 7, 3, 8, 17, 28, 90},
		Excepted: []int{0, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 6, 7, 8, 8, 17, 28, 90},
	},
}

// 基准测试case
// 排序最差情况
var nums = []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

// 排序最好情况
// var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// #########################################################################
// ####                            BubbleSort                           ####
// ####                             冒泡排序                            ####
// #########################################################################
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

// #########################################################################
// ####                            SelectionSort                        ####
// ####                               选择排序                          ####
// #########################################################################
func TestSelectionSort(t *testing.T) {
	for _, c := range cases {
		Convey("SelectionSort", t, func() {
			Convey(c.Name, func() {
				SelectionSort(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}
func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		SelectionSort(numbers)
	}
}

// #########################################################################
// ####                              InsertionSort                      ####
// ####                                 插入排序                        ####
// #########################################################################
func TestInsertionSort(t *testing.T) {
	for _, c := range cases {
		Convey("InsertionSort", t, func() {
			Convey(c.Name, func() {
				InsertionSort(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		InsertionSort(numbers)
	}
}

// #########################################################################
// ####                              ShellSort                          ####
// ####                               希尔排序                          ####
// #########################################################################
func TestShellSort(t *testing.T) {
	for _, c := range cases {
		Convey("ShellSort", t, func() {
			Convey(c.Name, func() {
				ShellSort(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		ShellSort(numbers)
	}
}

// #########################################################################
// ####                              MergeSort                          ####
// ####                               归并排序                          ####
// #########################################################################
func TestMergeSort(t *testing.T) {
	for _, c := range cases {
		Convey("MergeSort", t, func() {
			Convey(c.Name, func() {
				MergeSort(c.Input)
				So(c.Input, ShouldResemble, c.Excepted)
			})

		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := make([]int, len(nums))
		copy(numbers, nums)
		MergeSort(numbers)
	}
}

// #########################################################################
// ####                              QuickSort                          ####
// ####                               快速排序                          ####
// #########################################################################
func TestQuickSort(t *testing.T) {
	for _, c := range cases {
		Convey("QuickSort", t, func() {
			Convey(c.Name, func() {
				QuickSort(c.Input)
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
