package main

import (
	"fmt"
	"helpers"
	"strings"
)

func main() {
	var array []string = helpers.ReadInputFile("../inputs/input.txt")

	part2(array)
}

func part2(array []string) {
	// loop through all lines in input
	for _, line := range array {
		// loop through runes in string
		for i, r := range line {
			// if the rune is whitespace check all other lines for whitespace at this index
			if string(r) == " " {
				emptySection := true

				for _, otherLine := range array {
					runes := []rune(otherLine)

					if i < len(otherLine) {
						if string(runes[i]) != " " {
							emptySection = false
						}
					}
				}

				// if all lines have a whitespace rune at this index,
				// replace them with hyphens | rune "-" = int32 45
				if emptySection {
					for anotherIndex, anotherLine := range array {
						anotherRunes := []rune(anotherLine)

						anotherRunes[i] = 45
						anotherLine = string(anotherRunes)
						array[anotherIndex] = anotherLine
					}
				}
			}
		}
	}

	count := 0
	for _, r := range array[len(array)-1] {
		if string(r) != "-" && string(r) != " " {
			count++
		}
	}

	longestCount := 0
	onceCount := 0

	for _, r := range array[0] {
		if string(r) != "-" && string(r) != " " {
			onceCount++
		} else {
			if onceCount > longestCount {
				longestCount = onceCount
			}
			onceCount = 0
		}
	}

	// define a matrix for the current format and the intended format
	matrix := make([][]string, len(array))
	problemsMatrix := make([][]string, count)

	for i := range len(array) {
		matrix[i] = make([]string, len(array))

	}

	for i := range count {
		problemsMatrix[i] = make([]string, longestCount+1)
	}

	// populate values in the current format matrix
	for i, line := range array {
		matrixLine := strings.Split(line, "-")
		matrix[i] = matrixLine
	}

	// populate values from the current format matrix to the intended format matrix
	problemsX := 0
	problemsY := 0

	// loop through each item in the first line of the matrix
	for itemIndex, item := range matrix[0] {
		// loop through each run in the string item
		for runeIndex := range item {
			// loop through all lines in matrix looking at the same rune location,
			// and build each string item
			stringNum := ""
			for lineIndex := range len(matrix) - 1 {
				if runeIndex < len(matrix[lineIndex][itemIndex]) && string(matrix[lineIndex][itemIndex][runeIndex]) != " " {
					stringNum += string(matrix[lineIndex][itemIndex][runeIndex])
				}
			}

			// add string to problems matrix
			problemsMatrix[problemsX][problemsY] = stringNum
			problemsY++
		}
		problemsY = 0
		problemsX++
	}

	// loop through each item in the last line of the matrix
	for i, item := range matrix[len(matrix)-1] {
		for _, r := range item {
			if string(r) != " " {
				// populate + or * values
				problemsMatrix[i][len(problemsMatrix[i])-1] = string(r)
			}
		}
	}

	fmt.Println("part 2 total:", solveProblems(problemsMatrix))
}

func part1(array []string) {
	problems := buildPart1ProblemsMatrix(array)

	fmt.Println("part 1 total:", solveProblems(problems))
}

func buildPart1ProblemsMatrix(array []string) [][]string {
	// count intended elements
	count := 0
	for _, r := range array[len(array)-1] {
		if string(r) != " " {
			count++
		}
	}

	problems := make([][]string, count)

	for i := range count {
		problems[i] = make([]string, len(array))
	}

	xIndex := 0
	var item string

	for yIndex, line := range array {
		for charIndex, r := range line {
			if string(r) == " " {
				if item != "" {
					problems[xIndex][yIndex] = item
					item = ""
					xIndex++
				}
			} else {
				item = item + string(r)
			}

			if charIndex == len(line)-1 && string(r) != " " {
				problems[xIndex][yIndex] = item
				item = ""
			}
		}
		xIndex = 0
	}

	return problems
}

func solveProblems(problems [][]string) int64 {
	var total int64 = 0

	for _, problem := range problems {
		var answer int64 = 0
		if problem[len(problem)-1] == "+" {
			for i := 0; i < len(problem)-1; i++ {
				if problem[i] != "" {
					answer += helpers.StrToInt64(problem[i])
				}
			}
		} else {
			answer = helpers.StrToInt64(problem[0])
			for i := 1; i < len(problem)-1; i++ {
				if problem[i] != "" {
					answer *= helpers.StrToInt64(problem[i])
				}
			}
		}

		total += answer
	}

	return total
}
