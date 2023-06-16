package types

import (
	"strconv"
)

func IntToStr(x int) string {
	return strconv.Itoa(x)
}

func StrToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func BoolToInt(x bool) int {
	if x {
		return 1
	}
	return 0
}

func IntToBool(x int) bool {
	if x != 0 {
		return true
	}
	return false
}
