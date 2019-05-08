package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name     string
	Nums     []int
	Val      int
	Excepted int
}

// 单元测试
func TestRemoveElement(t *testing.T) {

	cases := []Case{
		{
			Name:     `empty`,
			Nums:     []int{},
			Val:      0,
			Excepted: 0,
		},
		{
			Name:     `[3, 3] - 3 => 0`,
			Nums:     []int{3, 3},
			Val:      3,
			Excepted: 0,
		},
		{
			Name:     `[3, 3] - 5 => 2`,
			Nums:     []int{3, 3},
			Val:      5,
			Excepted: 2,
		},
		{
			Name:     `[3, 2, 2, 3] - 3 => 2`,
			Nums:     []int{3, 2, 2, 3},
			Val:      3,
			Excepted: 2,
		},
		{
			Name:     `[0, 1, 2, 2, 3, 0, 4, 2] - 2 => 5`,
			Nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:      2,
			Excepted: 5,
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := RemoveElement(c.Nums, c.Val)
			So(result, ShouldEqual, c.Excepted)
		})
	}

}

// 基准测试
func BenchmarkRemoveElement(b *testing.B) {

	for i := 0; i < b.N; i++ {
		RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	}

}
