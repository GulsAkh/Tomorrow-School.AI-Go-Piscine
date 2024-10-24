package piscine

func Swap(a *int, b *int) {
	word1 := *a
	word2 := *b
	*a = word2
	*b = word1
}
