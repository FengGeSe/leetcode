package main

import ()

func FindDiagonalOrder(matrix [][]int) []int {

	flag := true
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
