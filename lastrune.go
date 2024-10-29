package poiscine

func LastRune(s string) rune {
	r := []rune(s)
	lastrune := r[len(r)-1]
	return lastrune
}
