package piscine

func Max(a []int) int {
	maX := a[0]
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a); i++ {
		if maX < a[i] {
			maX = a[i]
		}
	}
	return maX
}
