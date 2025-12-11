package main

import (
	"fmt"
	"helpers"
)

// int32 to string conversion table
// 64 = @ = roll of paper
// 46 = . = empty

var nextMatrix [][]int32

func main() {
	var array []string = helpers.ReadInputFile("../inputs/real_input.txt")
	var matrix [][]int32 = buildMatrix(array)

	part1(matrix)
	part2(matrix)
}

func part1(matrix [][]int32) {
	sum := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if checkSurroundingNodes(matrix, i, j) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func part2(matrix [][]int32) {
	nextMatrix = matrix

	totalSum := 0

whileloop:
	for {
		singleRunSum := 0
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if checkSurroundingNodes(matrix, i, j) {
					// count roll and remove it from next run
					singleRunSum++
					nextMatrix[i][j] = 46

				}
			}
		}

		if singleRunSum == 0 {
			break whileloop
		}

		matrix = nextMatrix
		totalSum += singleRunSum
	}

	fmt.Println(totalSum)
}

// check all nodes surrounding a coordinate within the matrix, excluding x/y values that would cause out of bounds errors
func checkSurroundingNodes(matrix [][]int32, x int, y int) bool {
	// ignore empty nodes
	if matrix[x][y] == 46 {
		return false
	}

	sum := 0
	// x - 1, y - 1 || x - 1, y || x - 1, y + 1
	// x, y - 1 || no check || x, y + 1
	// x + 1, y - 1 || x + 1, y || x + 1, y + 1

	if x > 0 {
		if y > 0 {
			if matrix[x-1][y-1] == 64 {
				sum++
			}
		}
		if y < len(matrix[x])-1 {
			if matrix[x-1][y+1] == 64 {
				sum++
			}
		}
		if matrix[x-1][y] == 64 {
			sum++
		}
	}

	if x < len(matrix)-1 {
		if y > 0 {
			if matrix[x+1][y-1] == 64 {
				sum++
			}
		}
		if y < len(matrix[x])-1 {
			if matrix[x+1][y+1] == 64 {
				sum++
			}
		}
		if matrix[x+1][y] == 64 {
			sum++
		}
	}

	if y > 0 {
		if matrix[x][y-1] == 64 {
			sum++
		}
	}

	if y < len(matrix[x])-1 {
		if matrix[x][y+1] == 64 {
			sum++
		}
	}

	return sum < 4
}

// builds a 2d slice from the input file
func buildMatrix(array []string) [][]int32 {
	var result [][]int32

	for _, line := range array {
		var intLine []int32

		for _, char := range line {
			intLine = append(intLine, char)
		}
		result = append(result, intLine)
	}

	return result
}

// debugging helper function
func printMatrix(matrix [][]int32) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
}
