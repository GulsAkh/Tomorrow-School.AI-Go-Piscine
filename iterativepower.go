package piscine

func IterativePower(nb int, power int) int {
	var num int = 1
	if nb < 0 || power < 0 || nb < 0 && power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		num *= nb
	}
	return num
}
