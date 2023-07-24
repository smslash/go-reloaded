package goreloaded

import (
	"strconv"
)

func Bin(binStr string) string {
	decimalInt, _ := strconv.ParseInt(binStr, 2, 64)
	decimalStr := strconv.FormatInt(decimalInt, 10)
	return decimalStr
}
