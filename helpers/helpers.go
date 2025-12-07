package helpers

import (
	"fmt"
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
