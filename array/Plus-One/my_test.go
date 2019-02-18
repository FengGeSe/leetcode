package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 表格驱动测试
func TestPlusOne(t *testing.T) {
	cases := []struct {
		Name   string
		Digits []int
		Result []int
	}{
		{"Input:[1,2,3] Output:[1,2,4]", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Input:[4,3,2,1] Output:[4,3,2,2]", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Input:[0] Output:[1]", []int{0}, []int{1}},
		{"Input:[9] Output:[1, 0]", []int{9}, []int{1, 0}},
		{"Input:[8,9,9,9] Output:[9,0,0,0]", []int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			So(PlusOne(c.Digits), IntSliceEqual, c.Result)
		})
	}

}

func IntSliceEqual(actual interface{}, expected ...interface{}) string {
	s := actual.([]int)
	t := expected[0].([]int)

	str := `Expected: '%v'
Actual:   '%v'
(Should be equal)`
	if len(s) != len(t) {
		return fmt.Sprintf(str, t, s)
	}
	for i, v := range s {
		if v != t[i] {
			return fmt.Sprintf(str, t, s)
		}
	}
	return ""
}

// 基准测试
func BenchmarkPlusOne(b *testing.B) {
	nums := []int{3, 6, 1, 0}
	for i := 0; i < b.N; i++ {
		PlusOne(nums)
	}
}
