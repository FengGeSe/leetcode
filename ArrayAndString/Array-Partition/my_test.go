package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name     string
	Input    []int
	Excepted int
}

// 单元测试
func TestArrayPartition(t *testing.T) {

	cases := []Case{
		{
			Name:     `empty`,
			Input:    []int{},
			Excepted: 0,
		},
		{
			Name:     `[1, 4, 3, 2] => 4`,
			Input:    []int{1, 4, 3, 2},
			Excepted: 4,
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := ArrayPartition(c.Input)
			So(result, ShouldEqual, c.Excepted)
		})
	}

}

// 基准测试
func BenchmarkStrStr(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ArrayPartition([]int{1, 4, 3, 2})
	}

}
