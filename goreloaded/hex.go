package goreloaded

import "strconv"

func Hex(hexStr string) string {
	decimalInt, _ := strconv.ParseInt(hexStr, 16, 64)
	decimalStr := strconv.FormatInt(decimalInt, 10)
	return decimalStr
}
