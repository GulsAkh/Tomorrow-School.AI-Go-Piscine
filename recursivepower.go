package piscine

func RecursivePower(nb int, power int) int {
	num := 1
	if power < 0 || nb == 0 {
		return 0
	} else if power == 0 || nb == 0 && power == 0 {
		return 1
	} else {
		num = nb * RecursivePower(nb, power-1)
	}
	return num
}
