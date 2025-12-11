package main

import (
	"fmt"
	"helpers"
	"strings"
)

type Range struct {
	low  int64
	high int64
}

func main() {
	var array []string = helpers.ReadInputFile("../inputs/test.txt")

	part2(array)
}

func part2(array []string) {
	var ranges []Range

	for _, line := range array {
		if strings.Contains(line, "-") {
			splitString := strings.Split(line, "-")

			convertedLow := helpers.StrToInt64(splitString[0])
			convertedHigh := helpers.StrToInt64(splitString[1])

			ranges = append(ranges, Range{low: convertedLow, high: convertedHigh})
		} else {
			break
		}
	}

	for _, item := range ranges {
		fmt.Print(item.low)
		fmt.Print("-")
		fmt.Println(item.high)
	}
}

func part1(array []string) {
	var lowEnd []int64
	var highEnd []int64
	var idsToCheck []int64

	// build out data arrays from input file
	for _, line := range array {
		if strings.Contains(line, "-") {
			splitString := strings.Split(line, "-")

			lowEnd = append(lowEnd, helpers.StrToInt64(splitString[0]))
			highEnd = append(highEnd, helpers.StrToInt64(splitString[1]))
		} else if line != "" {
			idsToCheck = append(idsToCheck, helpers.StrToInt64(line))
		}

	}

	freshIngredients := 0

	// iterate through each id to check
	for i := 0; i < len(idsToCheck); i++ {
		// iterate through each range
		for j := 0; j < len(lowEnd); j++ {
			// if id to check is in the range mark it as a fresh
			// ingredient and move to the next id
			if idsToCheck[i] >= lowEnd[j] && idsToCheck[i] <= highEnd[j] {
				freshIngredients++
				break
			}
		}
	}

	fmt.Println("freshIngredients:", freshIngredients)
}
