package main

import (
	"fmt"
)

func main() {
	maze := InitMaze()

	steps := Walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println(steps)
}

func InitMaze() [][]int {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}

	return maze
}

type Point struct {
	I, J int
}

func (p Point) Add(r Point) Point {
	return Point{p.I + r.I, p.J + r.J}
}

func (p Point) At(grid [][]int) (int, bool) {
	if p.I < 0 || p.I >= len(grid) {
		return 0, false
	}
	if p.J < 0 || p.J >= len(grid[p.I]) {
		return 0, false
	}
	return grid[p.I][p.J], true
}

var dirs = [4]Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func Walk(maze [][]int, start, end Point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []Point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.Add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			if val, ok := next.At(maze); !ok || val == 1 {
				continue
			}
			if val, ok := next.At(steps); !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}

			//
			curSteps, _ := cur.At(steps)
			steps[next.I][next.J] = curSteps + 1
			queue = append(queue, next)

		}
	}

	return steps
}
