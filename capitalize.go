package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	str := ""
	if s == "" {
		return ""
	}
	if runes[0] >= 'a' && runes[0] <= 'z' {
		str += string(runes[0] - 32)
	} else {
		str += string(runes[0])
	}
	for i := 1; i < len(runes); i++ {
		if (runes[i-1] < 'A' || runes[i-1] > 'Z') && (runes[i-1] < 'a' || runes[i-1] > 'z') && (runes[i-1] < '0' || runes[i-1] > '9') {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				str += string(runes[i] - 32)
			} else {
				str += string(runes[i])
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				str += string(runes[i] + 32)
			} else {
				str += string(runes[i])
			}
		}
	}
	return str
}
