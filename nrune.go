package piscine

func NRune(s string, n int) rune {
	var runes rune
	if n >= 0 && n <= len(s) {
		runes = rune(s[n-1])
	}
	return runes
}
