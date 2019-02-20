package main

import ()

func SpiralOrder(matrix [][]int) []int {
	direction := 0 // 0 右, 1 下, 2 左, 3 上

	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])

	trace := [][]int{}
	for i := 0; i < m; i++ {
		trace = append(trace, make([]int, n))
	}

	x, y := 0, 0 // 当前点的坐标
	result := []int{}
	for len(result) < m*n {
		switch direction % 4 {
		case 0: // 向右走
			if trace[x][y] != 1 {
				result = append(result, matrix[x][y])
				trace[x][y] = 1
			}
			if y+1 >= n || trace[x][y+1] == 1 {
				direction++
				continue
			}
			y++
		case 1: // 向下走
			if trace[x][y] != 1 {
				result = append(result, matrix[x][y])
				trace[x][y] = 1
			}
			if x+1 >= m || trace[x+1][y] == 1 {
				direction++
				continue
			}
			x++
		case 2: // 向左走
			if trace[x][y] != 1 {
				result = append(result, matrix[x][y])
				trace[x][y] = 1
			}
			if y-1 < 0 || trace[x][y-1] == 1 {
				direction++
				continue
			}
			y--
		case 3: // 向上走
			if trace[x][y] != 1 {
				result = append(result, matrix[x][y])
				trace[x][y] = 1
			}
			if x-1 < 0 || trace[x-1][y] == 1 {
				direction++
				continue
			}
			x--
		}

	}

	return result

}
