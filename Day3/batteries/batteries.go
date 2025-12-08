package main

import (
	"fmt"
	"helpers"
)

func main() {
	var array = helpers.ReadInputFile("../inputs/test.txt")

	for _, line := range array {
		fmt.Println(line)
	}
}
