package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name     string
	Input    []byte
	Excepted []byte
}

// 单元测试
func TestReverseString(t *testing.T) {

	cases := []Case{
		{
			Name:     `empty`,
			Input:    []byte{},
			Excepted: []byte{},
		},
		{
			Name:     `['a'] => ['a']`,
			Input:    []byte{'a'},
			Excepted: []byte{'a'},
		},
		{
			Name:     `['h', 'e', 'l', 'l', 'o'] => ['o', 'l', 'l', 'e', 'h']`,
			Input:    []byte{'h', 'e', 'l', 'l', 'o'},
			Excepted: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			Name:     `['H', 'a', 'n', 'n', 'a', 'h'] => ['h', 'a', 'n', 'n', 'a', 'H']`,
			Input:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			Excepted: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			ReverseString(c.Input)
			So(c.Input, ShouldResemble, c.Excepted)
		})
	}

}

// 基准测试
func BenchmarkStrStr(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ReverseString([]byte{'h', 'a', 'n', 'n', 'a', 'H'})
	}

}
