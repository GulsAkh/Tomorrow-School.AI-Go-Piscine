package piscine

func UltimateDivMod(a *int, b *int) {
	first := *a
	*a = first / *b
	*b = first % *b
}
