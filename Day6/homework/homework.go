package main

import (
	"fmt"
	"helpers"
)

// rune(int32) to string conversion table
// 32 =   = whitespace

func main() {
	var array []string = helpers.ReadInputFile("../inputs/test.txt")

	part1(array)
}

// intended problems array
// [[123 45 6 *] [328 64 98 +] [51 387 215 *] [64 23 314 +]]

func part1(array []string) {
	//var problems = [][]string

	for _, line := range array {
		for _, r := range line {

			// r is a rune
			fmt.Println(string(r))
		}
	}
}
