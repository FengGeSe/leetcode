package main

import ()

func FindDiagonalOrder(matrix [][]int) []int {

	flag := true // 对角线方向  true:右上  false:左下
	x, y := 0, 0
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}
	m, n := len(matrix), len(matrix[0])

	for x < m && y < n {
		result = append(result, matrix[x][y])
		if flag {
			// flag = true时 (—1, +1)
			if x-1 < 0 || y+1 >= n {
				flag = !flag
			}

			// 右上遍历的时候，当前点的位置会影响怎么移动到下一个点
			if y == n-1 {
				x++
			} else if x == 0 {
				y++
			} else {
				x--
				y++
			}

		} else {
			// flag = false时 (+1, -1)
			if x+1 >= m || y-1 < 0 {
				flag = !flag
			}

			if x == m-1 {
				y++
			} else if y == 0 {
				x++
			} else {
				x++
				y--
			}

		}

	}

	return result

}
