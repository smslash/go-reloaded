package goreloaded

func Article(result string) string {
	template := Fields(result)
	for i := 0; i < len(template); i++ {
		if template[i] == "an" || template[i] == "An" || template[i] == "AN" {
			if i+1 < len(template) && !IsVowel(template[i+1][0]) {
				if template[i] == "an" {
					template[i] = "a"
				} else {
					template[i] = "A"
				}
			}
		} else if template[i] == "a" || template[i] == "A" {
			if i+1 < len(template) && IsVowel(template[i+1][0]) {
				if template[i] == "a" {
					template[i] = "an"
				} else {
					template[i] = "An"
				}
			}
		}
	}

	var final_res string
	for _, value := range template {
		final_res += value + " "
	}
	final_res = final_res[:len(final_res)-1]

	for i := 0; i < len(final_res); i++ {
		if i+1 < len(final_res) && final_res[i+1] == '\n' {
			tmp1 := final_res[:i]
			tmp2 := final_res[i+1:]
			final_res = tmp1 + tmp2
		}
		if i+2 < len(final_res) && final_res[i] == '\n' && final_res[i+1] == ' ' {
			tmp1 := final_res[:i+1]
			tmp2 := final_res[i+2:]
			final_res = tmp1 + tmp2
		}
	}
	return final_res
}

func IsVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'h'
}
