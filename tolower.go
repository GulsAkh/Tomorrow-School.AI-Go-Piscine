package piscine

func ToLower(s string) string {
	result := make([]rune, len(s))
	for i, el := range s {
		if el >= 'A' && el <= 'Z' {
			result[i] = el + 32
		} else {
			result[i] = el
		}
	}
	return string(result)
}
