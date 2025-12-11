package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

func main() {
	var array = helpers.ReadInputFile("../inputs/input.txt")

	var sum int64 = 0

	// Iterate over each line
	for _, line := range array {
		array := strings.Split(line, ",")

		for _, str := range array {
			ids := strings.Split(str, "-")

			var start int64 = helpers.StrToInt64(ids[0])
			var end int64 = helpers.StrToInt64(ids[1])

			// iterate over each range of ids
			for i := start; i <= end; i++ {
				numStr := strconv.FormatInt(i, 10)

				// only strings where the first half and last half match matter
				// ignore all others
				if len(numStr)%2 == 0 {
					mid := len(numStr) / 2

					firstHalf := numStr[:mid]
					secondHalf := numStr[mid:]

					// if a match is found add full id to sum
					if firstHalf == secondHalf {
						sum = sum + i
					}
				}
			}
		}

		fmt.Println(sum) // print sum for input into browser
	}
}
