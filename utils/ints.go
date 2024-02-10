package utils

import "strconv"

func ConvertStrToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return val
}
