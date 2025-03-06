package main

import "math"

func maxPath(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			triangle[row][col] += int(math.Max(float64(triangle[row+1][col]), float64(triangle[row+1][col+1])))
		}
	}
	return triangle[0][0]
}
