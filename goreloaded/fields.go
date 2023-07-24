package goreloaded

func Fields(s string) []string {
	var fields []string
	start := -1
	for i, r := range s {
		if IsSpace(r) {
			if start >= 0 {
				fields = append(fields, s[start:i])
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
		if r == 10 {
			fields = append(fields, "\n")
		}
	}
	if start >= 0 {
		fields = append(fields, s[start:])
	}
	return fields
}

func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
