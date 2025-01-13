//https://leetcode.com/problems/spiral-matrix/

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int {}
	}

	var op []int 

	row := len(matrix)
	col := len(matrix[0])

	matrixSize := row * col

	const (
		UP = 0
		RIGHT = 1
		LEFT = 2
		DOWN = 3
	)

	var i, j int
	
	direction := RIGHT 

	UPWALL := 0
	RIGHTWALL := col
	BOTTOMWALL := row
	LEFTWALL := -1

	for matrixSize != len(op) {
		if direction == RIGHT {
			for j < RIGHTWALL {
				op = append(op, matrix[i][j])
				j++
			}
			i++
			j--
			RIGHTWALL--
			direction = DOWN
		} else if direction == DOWN {
			for i < BOTTOMWALL {
				op = append(op, matrix[i][j])
				i++
			}
			i--
			j--
			BOTTOMWALL--
			direction = LEFT
		} else if direction == LEFT {
			for j > LEFTWALL {
				op = append(op, matrix[i][j])
				j--
			}
			i--
			j++
			LEFTWALL++
			direction = UP 
		} else {
			for i > UPWALL {
				op = append(op, matrix[i][j])
				i--
			}
			i++
			j++
			UPWALL++
			direction = RIGHT
		}
	} 

	return op
}

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	fmt.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
	fmt.Println(spiralOrder([][]int{{1,2,3}}))
	fmt.Println(spiralOrder([][]int{{1},{2},{3}}))
}