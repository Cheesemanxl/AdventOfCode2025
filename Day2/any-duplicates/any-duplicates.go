package main

import (
	"bufio"
	"fmt"
	"helpers"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open inputs file
	file, err := os.Open("../inputs/real_input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Instantiate scanner to parse txt file
	scanner := bufio.NewScanner(file)

	var sum int64 = 0

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		array := strings.Split(line, ",")

		for _, str := range array {
			ids := strings.Split(str, "-")

			var start int64 = helpers.StrToInt64(ids[0])
			var end int64 = helpers.StrToInt64(ids[1])

			// iterate over each range of ids
			for i := start; i <= end; i++ {
				numStr := strconv.FormatInt(i, 10)

				match_found := false

				for j := 1; j <= len(numStr)/2; j++ {
					// only substrings of length equal to
					// factors of the length of the string being observed matter
					if len(numStr)%j == 0 {
						check := numStr[:j] // substring to check from beginning of string to j
						last_index := 0

						// iterate through each divided segment
						for k := len(check); k <= len(numStr); k = k + len(check) {
							next_segment := numStr[last_index:k]
							last_index = k
							if check == next_segment {
								match_found = true
							} else {
								// if a match is ever not found in the string
								// break out to next substring check and reset
								match_found = false
								break
							}
						}

						// if the k loops is completed with match_found still equal to true
						// then add full id to sum
						if match_found {
							sum = sum + helpers.StrToInt64(numStr)
							match_found = false
							break
						}
					}
				}
			}
		}
	}

	fmt.Println("Sum:", sum) // print sum for input into browser
}
