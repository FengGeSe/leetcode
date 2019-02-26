package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name     string
	Haystack string
	Needle   string
	Output   int
}

// 单元测试
func TestAddBinary(t *testing.T) {

	cases := []Case{
		{
			Name:     "Haystack is an empty string",
			Haystack: "",
			Needle:   "a",
			Output:   -1,
		},
		{
			Name:     "Needle is an empty string",
			Haystack: "hello",
			Needle:   "",
			Output:   0,
		},
		{
			Name:     "Haystack and Needld are empty string",
			Haystack: "",
			Needle:   "",
			Output:   0,
		},
		{
			Name:     "Find it",
			Haystack: "hello",
			Needle:   "ll",
			Output:   2,
		},
		{
			Name:     "Not Find it",
			Haystack: "aaaaa",
			Needle:   "bba",
			Output:   -1,
		},
		{
			Name:     "Same string",
			Haystack: "a",
			Needle:   "a",
			Output:   0,
		},
		{
			Name:     "manay string",
			Haystack: "aaaa",
			Needle:   "a",
			Output:   0,
		},
		{
			Name:     "Find it",
			Haystack: "mississippi",
			Needle:   "issip",
			Output:   4,
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := StrStr(c.Haystack, c.Needle)
			So(result, ShouldEqual, c.Output)
		})
	}

}

// 基准测试
func BenchmarkStrStr(b *testing.B) {

	for i := 0; i < b.N; i++ {
		StrStr("Hello", "ll")
	}

}
