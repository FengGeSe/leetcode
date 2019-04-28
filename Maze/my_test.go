package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	maze = [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	steps = [][]int{
		{0, 0, 4, 5, 6},
		{1, 2, 3, 0, 7},
		{2, 0, 4, 0, 8},
		{0, 0, 0, 10, 9},
		{0, 0, 12, 11, 0},
		{0, 0, 13, 12, 13},
	}
)

// 表格驱动测试
func TestWalk(t *testing.T) {
	cases := []struct {
		Name   string
		Maze   [][]int
		Result [][]int
	}{
		{
			Name:   "5*6迷宫",
			Maze:   maze,
			Result: steps,
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			rst := Walk(c.Maze, Point{0, 0}, Point{len(c.Maze) - 1, len(c.Maze[0]) - 1})
			So(rst, ShouldResemble, c.Result)
		})
	}

}

// 基准测试
func BenchmarkWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})
	}
}
