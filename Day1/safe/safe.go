package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open inputs file
	file, err := os.Open("real_input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Instantiate scanner to parse txt file
	scanner := bufio.NewScanner(file)

	count := 0
	current := 50

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Get int representation of direction and amount from input string
		direction := line[:1] //really cool substring syntax
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		// if amount is greater than 100 count each 100 as one hit of zero
		extra_zero_hits := amount / 100
		count = count + extra_zero_hits

		// normalize amount for values over 100
		amount = amount % 100

		// Add or subtract depending on direction
		if direction == "R" {
			current = current + amount
		} else {
			// if already starting on zero and rotating left we have already counted that "zero click"
			// so remove one to avoid duplication in the next normalization stage
			if current == 0 {
				count--
			}
			current = current - amount
		}

		// normalize current location between 0-99
		if current > 100 {
			current = current - 100
			count++ // count rightways zero clicks
		} else if current < 0 {
			current = 100 + current
			count++ // count leftways zero clicks
		} else if current == 100 {
			current = 0
		}

		if current == 0 {
			count++ // count exact zero clicks
		}
	}

	fmt.Println(count) // print result for input into browser

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
