package main

import (
	"fmt"
	"helpers"
)

func main() {
	var array []string = helpers.ReadInputFile("../inputs/test.txt")

	part1(array)
}

func part1(array []string) {
	indeciesToCheck := make(map[int]bool)

	for i, r := range array[0] {
		if string(r) == "S" {
			indeciesToCheck[i] = true
		}
	}

	total := 0

	for _, line := range array {
		for i, r := range line {
			if string(r) == "^" && indeciesToCheck[i] {
				indeciesToCheck[i] = false

				if i-1 >= 0 {
					indeciesToCheck[i-1] = true
				}

				if i+1 < len(line) {
					indeciesToCheck[i+1] = true
				}

				total++
			}
		}
	}

	fmt.Println(total)
}
