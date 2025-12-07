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

		// Process inputs
		direction := line[:1] //really cool substring syntax
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		extra_zero_hits := amount / 100

		amount = amount % 100

		count = count + extra_zero_hits

		// Add or subtract depending on direction
		if direction == "R" {
			current = current + amount
		} else {
			if current == 0 {
				count--
			}
			current = current - amount
		}

		// normalize current location between 0-99
		if current > 100 {
			current = current - 100
			count++
		} else if current < 0 {
			current = 100 + current
			count++
		} else if current == 100 {
			current = 0
		}

		if current == 0 {
			count++
		}
	}

	fmt.Println(count)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
