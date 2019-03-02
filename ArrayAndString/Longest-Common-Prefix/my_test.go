package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name   string
	Strs   []string
	Output string
}

// 单元测试
func TestAddBinary(t *testing.T) {

	cases := []Case{
		{
			Name:   `["flower", "flow", "flight"] => "fl"`,
			Strs:   []string{"flower", "flow", "flight"},
			Output: "fl",
		},
		{
			Name:   `["dog", "racecar", "car"] => ""`,
			Strs:   []string{"dog", "racecar", "car"},
			Output: "",
		},
		{
			Name:   `["aca", "cba"] => ""`,
			Strs:   []string{"aca", "cba"},
			Output: "",
		},
		{
			Name:   `[] => ""`,
			Strs:   []string{},
			Output: "",
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := LongestCommonPrefix(c.Strs)
			So(result, ShouldEqual, c.Output)
		})
	}

}

// 基准测试
func BenchmarkStrStr(b *testing.B) {

	for i := 0; i < b.N; i++ {
		LongestCommonPrefix([]string{"flower", "flow", "flight"})
	}

}
