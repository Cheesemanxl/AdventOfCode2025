package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func StrToInt64(input string) int64 {
	result, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return -1
	}

	return result
}

func ReadInputFile(filepath string) []string {
	var array = []string{}

	// Open inputs file
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Instantiate scanner to parse txt file
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return array
}
