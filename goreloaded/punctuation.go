package goreloaded

func Punctuation(result string) string {
	result += " "
	new_result := ""
	for i := 0; i < len(result); i++ {
		if result[i] == ' ' && i+1 < len(result) && IsSign(result[i+1]) {
			for j := i + 1; j < len(result); j++ {
				if !IsSign(result[j]) {
					i = j
					new_result += " " + string(result[j])
					break
				} else {
					new_result += string(result[j])
				}
			}
		} else {
			new_result += string(result[i])
		}
	}
	return new_result
}

func IsSign(s byte) bool {
	return s == '.' || s == ',' || s == '?' || s == '!' || s == ':' || s == ';'
}
