package piscine

func f(n1, n2 int) int {
	if n1 < 0 && n2 < 0 {
		if n1 < n2 {
			return -1
		}
	} else {
		if n1 > n2 {
			return 1
		} else if n1 == n2 {
			return 0
		}
	}
	return -1
}

func IsSorted(f func(a, b int) int, a []int) bool {
	flag := true
	for i := 0; i < len(a)-1; i++ {
		b := a[i]
		c := a[i+1]
		if (f(b, c)) > 0 {
			flag = false
		}
	}
	return flag
}
