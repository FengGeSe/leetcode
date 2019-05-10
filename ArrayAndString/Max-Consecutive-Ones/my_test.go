package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name     string
	Nums     []int
	Excepted int
}

// 单元测试
func TestFindMaxConsecutiveOnes(t *testing.T) {

	cases := []Case{
		{
			Name:     `{1, 1, 0, 1, 1, 1} -> 3`,
			Nums:     []int{1, 1, 0, 1, 1, 1},
			Excepted: 3,
		},
		{
			Name:     `{0, 1, 1, 0, 1, 1, 1} -> 3`,
			Nums:     []int{0, 1, 1, 0, 1, 1, 1},
			Excepted: 3,
		},
		{
			Name:     `{0, 0, 0} -> 0`,
			Nums:     []int{0, 0, 0},
			Excepted: 0,
		},
		{
			Name:     `{1, 1, 1} -> 3`,
			Nums:     []int{1, 1, 1},
			Excepted: 3,
		},
		{
			Name:     `{1, 1, 1, 0} -> 3`,
			Nums:     []int{1, 1, 1, 0},
			Excepted: 3,
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			actual := FindMaxConsecutiveOnes(c.Nums)
			So(actual, ShouldEqual, c.Excepted)
		})
	}

}

// 基准测试
func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	}

}
