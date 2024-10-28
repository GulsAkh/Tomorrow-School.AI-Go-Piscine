package piscine

func RecursiveFactorial(nb int) int {
	var num int = 1
	if nb == 0 || nb == 1 {
		return 1
	} else if nb < 0 || nb >= 21 {
		return 0
	} else {
		num = nb * RecursiveFactorial(nb-1)
	}
	return num
}
