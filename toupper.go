package piscine

func ToUpper(s string) string {
	result := make([]rune, len(s))
	for i, el := range s {
		if el >= 'a' && el <= 'z' {
			result[i] = el - 32
		} else {
			result[i] = el
		}
	}
	return string(result)
}
