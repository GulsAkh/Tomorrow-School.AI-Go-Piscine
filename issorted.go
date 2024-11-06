package piscine

func f(n1, n2 int) int {
	if n1 > n2 {
		return 1
	} else if n1 == n2 {
		return 0
	}
	return -1
}

func IsSorted(f func(a, b int) int, a []int) bool {
	flag := true
	for i := 0; i < len(a)-1; i++ {
		if (f(a[i], a[i+1])) > 0 {
			flag = false
		}
	}
	return flag
}
