package goreloaded

import "fmt"

func Apostrophe(result string) string {
	result = " " + result + " "
	new_result := ""
	col := 0
	for i := 1; i < len(result)-1; i++ {
		if result[i] == 39 && (Space(result[i-1]) || Space(result[i+1])) && col == 0 {
			if Space(result[i-1]) {
				new_result += "'"
				i++
			} else {
				new_result += "' "
				i++
			}
			col++
		} else if result[i] == 39 && (Space(result[i-1]) || Space(result[i+1])) && col == 1 {
			if Space(result[i-1]) {
				new_result = new_result[:len(new_result)-1]
				new_result += "' "
				i++
			} else {
				new_result += "'"
			}
			col--
		} else {
			new_result += string(result[i])
		}
	}
	if col != 0 {
		fmt.Println("Error: close apostrophe")
		return " "
	}
	return new_result
}

func Space(s byte) bool {
	return s == ' '
}
