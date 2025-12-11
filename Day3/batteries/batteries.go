package main

import (
	"fmt"
	"helpers"
	"strconv"
)

type BatteryDigit struct {
	index   int
	voltage int64
}

func main() {
	var array = helpers.ReadInputFile("../inputs/input.txt")

	var sum int64 = 0

	for _, line := range array {
		sum += getHighestNumber(line, 12)
	}

	fmt.Println(sum)
}

func getHighestNumber(batteries string, numDigits int) int64 {
	var result string
	startIndex := 0

	// iterate an amount of times equal to the number of digits in the intended result
	for remainingDigits := numDigits; remainingDigits > 0; remainingDigits-- {
		var maxDigit int64 = 0
		maxIndex := startIndex

		// by removing the end portion of the batteries string we only look at potential digits that would be in the highest "tens place"
		// ex. input = 234234234234278, for 12 digits first iteration will look at only 2342 which would be the only potential options for the "hundred billons place"
		endIndex := len(batteries) - remainingDigits

		// starting at the last index looked at, iterate until the length of the
		// input minus the amount of digits still needed in the result
		for i := startIndex; i <= endIndex; i++ {
			var numeric int64 = helpers.StrToInt64(string(batteries[i])) // convert byte to string to int64 for comparison
			if numeric > maxDigit {
				maxDigit = numeric
				maxIndex = i
			}
		}

		// append result and iterate index
		result = result + strconv.FormatInt(maxDigit, 10)
		startIndex = maxIndex + 1
	}

	return helpers.StrToInt64(result)
}

// First attempt at the challenge, worked for 2 digits, was not feasible for 12 digits
func bruteForceTwoDigitJoltage(array []string) {
	sum := 0

	// Iterate over lines in input file
	for _, line := range array {
		highestDigit := BatteryDigit{index: 0, voltage: 0}
		joltage := 0

		// Iterate over characters in the string
		for i, char := range line {

			// Convert character back to int
			digit := helpers.StrToInt64(string(char))
			check := BatteryDigit{index: i, voltage: digit}
			if check.voltage > highestDigit.voltage {
				highestDigit = check
			}
		}

		// check every possible combination of 2 digit numbers

		// forward check
		if highestDigit.index < len(line)-1 {
			for i := highestDigit.index + 1; i < len(line); i++ {
				str := strconv.FormatInt(highestDigit.voltage, 10) + string(line[i])
				num, _ := strconv.Atoi(str)
				if num > joltage {
					joltage = num
				}
			}
		}

		// backward check
		if highestDigit.index > 0 {
			for i := highestDigit.index - 1; i >= 0; i-- {
				str := string(line[i]) + strconv.FormatInt(highestDigit.voltage, 10)
				num, _ := strconv.Atoi(str)
				if num > joltage {
					joltage = num
				}
			}
		}

		sum += joltage
	}
	fmt.Println("Sum:", sum)
}
