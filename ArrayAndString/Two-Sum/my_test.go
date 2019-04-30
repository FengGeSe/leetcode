package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Case struct {
	Name    string
	Numbers []int
	Target  int
	Output  []int
}

// 单元测试
func TestTwoSum(t *testing.T) {

	cases := []Case{
		{
			Name:    "[2, 7, 11, 15] => 9",
			Numbers: []int{2, 7, 11, 15},
			Target:  9,
			Output:  []int{1, 2},
		},
		{
			Name:    "[...] => 542",
			Numbers: []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997},
			Target:  542,
			Output:  []int{24, 32},
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			result := TwoSum(c.Numbers, c.Target)
			So(result, ShouldResemble, c.Output)
		})
	}

}

// 基准测试
func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{2, 7, 11, 15}, 9)
	}

}
