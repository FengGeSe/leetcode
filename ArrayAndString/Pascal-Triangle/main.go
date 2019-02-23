package main

import ()

func Generate(numRows int) [][]int {

	result := [][]int{}

	for i := 0; i < numRows; i++ {
		rowResult := []int{}
		for j := 0; j < i+1; j++ {
			// (i, j)
			if j == 0 || j == i {
				rowResult = append(rowResult, 1)
				continue
			}

			val := result[i-1][j-1] + result[i-1][j]
			rowResult = append(rowResult, val)
		}
		result = append(result, rowResult)
	}

	return result

}
