package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name   string
	Input  int
	Output [][]int
}

// 单元测试
func TestGenerate(t *testing.T) {

	cases := []Case{
		{
			Name:   "0",
			Input:  0,
			Output: [][]int{},
		},
		{
			Name:  "1",
			Input: 1,
			Output: [][]int{
				{1},
			},
		},
		{
			Name:  "2",
			Input: 2,
			Output: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			Name:  "5",
			Input: 5,
			Output: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := Generate(c.Input)
			So(result, ShouldResemble, c.Output)
		})
	}

}

// 基准测试
func BenchmarkGenerate(b *testing.B) {
	input := 5

	for i := 0; i < b.N; i++ {
		Generate(input)
	}

}
