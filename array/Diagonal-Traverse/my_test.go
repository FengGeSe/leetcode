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
func TestFindDiagonalOrder(t *testing.T) {

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
			Output: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			Name: "C",
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			Output: []int{1, 2, 5, 9, 6, 3, 4, 7, 10, 13, 14, 11, 8, 12, 15, 16},
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := FindDiagonalOrder(c.Input)
			So(result, ShouldResemble, c.Output)
		})
	}

}

// 基准测试
func BenchmarkFindDiagonalOrder(b *testing.B) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	for i := 0; i < b.N; i++ {
		FindDiagonalOrder(input)
	}

}
