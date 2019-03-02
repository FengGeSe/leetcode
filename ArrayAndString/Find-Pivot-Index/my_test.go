package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 表格驱动测试
func TestFindPivotIndex(t *testing.T) {
	cases := []struct {
		Name  string
		Nums  []int
		Pivot int
	}{
		{"Input:[] Output:-1", []int{}, -1},
		{"Input:[0, 0] Output:0", []int{0, 0}, 0},
		{"Input:[1, 7, 3, 6, 5, 6] Output:3", []int{1, 7, 3, 6, 5, 6}, 3},
		{"Input:[1, 2, 2, 1] Output:-1", []int{1, 2, 2, 1}, -1},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			So(FindPivotIndex(c.Nums), ShouldEqual, c.Pivot)
		})
	}

}

// 基准测试
func BenchmarkFindPivotIndex(b *testing.B) {
	nums := []int{1, 7, 3, 6, 5, 6}
	for i := 0; i < b.N; i++ {
		FindPivotIndex(nums)
	}
}
