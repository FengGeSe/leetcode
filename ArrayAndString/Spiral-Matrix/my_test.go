package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name   string
	Input  [][]int
	Output []int
}

// 单元测试
func TestSprialOrder(t *testing.T) {

	cases := []Case{
		{
			Name:   "A",
			Input:  [][]int{},
			Output: []int{},
		},
		{
			Name: "B",
			Input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			Name: "C",
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			Output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := SpiralOrder(c.Input)
			So(result, ShouldResemble, c.Output)
		})
	}

}

// 基准测试
func BenchmarkSprialOrder(b *testing.B) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	for i := 0; i < b.N; i++ {
		SpiralOrder(input)
	}

}
