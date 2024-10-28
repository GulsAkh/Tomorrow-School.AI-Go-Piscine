package piscine

func IterativeFactorial(nb int) int {
	var num int = 1
	if nb < 0 || nb >= 21 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else if nb > 1 {
		num = nb * IterativeFactorial(nb-1)
	}
	return num
}
