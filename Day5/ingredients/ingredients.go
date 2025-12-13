package main

import (
	"fmt"
	"helpers"
	"slices"
	"strings"
)

func main() {
	var array []string = helpers.ReadInputFile("../inputs/input.txt")

	part2(array)
}

func part2(array []string) {
	var lows []int64
	var highs []int64

	// translate data from one string slice into two int64 slices
	for _, line := range array {
		if strings.Contains(line, "-") {
			splitString := strings.Split(line, "-")

			lows = append(lows, helpers.StrToInt64(splitString[0]))
			highs = append(highs, helpers.StrToInt64(splitString[1]))
		} else {
			break
		}
	}

	// consolidate all ranges that overlap
whileloop:
	for {
		isRemainingOverlaps := false
	overlapCheck:
		for i := 0; i < len(lows); i++ {
			for j := i + 1; j < len(lows); j++ {
				if isInRange(lows[i], highs[i], lows[j], highs[j]) {
					isRemainingOverlaps = true
					break overlapCheck
				}
			}
		}

		if !isRemainingOverlaps {
			break whileloop
		}

		for i := 0; i < len(lows); i++ {
			for j := i + 1; j < len(lows); j++ {
				// if either end of one range is in the other range,
				// use the lowest or highest value from the combination of ranges,
				// assign it to the first range values and delete the other range values
				if isInRange(lows[i], highs[i], lows[j], highs[j]) {
					if lows[i] > lows[j] {
						lows[i] = lows[j]
					}

					if highs[i] < highs[j] {
						highs[i] = highs[j]
					}

					lows = slices.Delete(lows, j, j+1)
					highs = slices.Delete(highs, j, j+1)
				}
			}
		}
	}

	// remove duplicate ranges
	for i := 0; i < len(lows); i++ {
		for j := i + 1; j < len(lows); j++ {
			if lows[i] == lows[j] && highs[i] == highs[j] {
				lows = slices.Delete(lows, j, j+1)
				highs = slices.Delete(highs, j, j+1)
			}
		}
	}

	var sum int64 = 0

	for index, item := range lows {
		sum += highs[index] - item + 1
	}

	fmt.Println(sum)
}

func isInRange(low1 int64, high1 int64, low2 int64, high2 int64) bool {
	if low1 >= low2 && low1 <= high2 {
		return true
	}

	if high1 <= high2 && high1 >= low2 {
		return true
	}

	if low2 >= low1 && low2 <= high1 {
		return true
	}

	if high2 <= high1 && high2 >= low1 {
		return true
	}

	return false
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
