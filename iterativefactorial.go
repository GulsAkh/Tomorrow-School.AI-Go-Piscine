package main

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	}
	if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	}
	return 1
}
