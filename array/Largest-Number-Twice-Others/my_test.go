package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 表格驱动测试
func TestLargestNumberTwiceOthers(t *testing.T) {
	cases := []struct {
		Name  string
		Nums  []int
		Index int
	}{
		{"Input:[3, 6, 1, 0] Output:1", []int{3, 6, 1, 0}, 1},
		{"Input:[1, 2, 3, 4] Output:-1", []int{1, 2, 3, 4}, -1},
		{"Input:[] Output:-1", []int{}, -1},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			So(LargestNumberTwiceOthers(c.Nums), ShouldEqual, c.Index)
		})
	}

}

// 基准测试
func BenchmarkLargestNumberTwiceOthers(b *testing.B) {
	nums := []int{3, 6, 1, 0}
	for i := 0; i < b.N; i++ {
		LargestNumberTwiceOthers(nums)
	}
}
