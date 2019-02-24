package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name   string
	InputA string
	InputB string
	Output string
}

// 单元测试
func TestAddBinary(t *testing.T) {

	cases := []Case{
		{
			Name:   "A",
			InputA: "11",
			InputB: "1",
			Output: "100",
		},
		{
			Name:   "B",
			InputA: "1010",
			InputB: "1011",
			Output: "10101",
		},
		{
			Name:   "C",
			InputA: "10",
			InputB: "1100",
			Output: "1110",
		},
		{
			Name:   "D",
			InputA: "11",
			InputB: "111",
			Output: "1010",
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := AddBinary(c.InputA, c.InputB)
			So(result, ShouldEqual, c.Output)
		})
	}

}

// 基准测试
func BenchmarkAddBinary(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AddBinary("1010", "1011")
	}

}
