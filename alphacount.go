package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for _, el := range runes {
		if el >= 'A' && el <= 'Z' || el >= 'a' && el <= 'z' {
			count++
		}
	}
	return count
}
