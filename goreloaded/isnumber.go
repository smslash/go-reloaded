package goreloaded

func IsNumber(s byte) bool {
	if s < '0' || s > '9' {
		return false
	}
	return true
}
