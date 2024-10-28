package piscine

func IterativePower(nb int, power int) int {
	var num int = 1
	if power < 0 {
		return 0
	}
	for i := 0; i < power; i++ {
		num *= nb
	}
	return num
}
