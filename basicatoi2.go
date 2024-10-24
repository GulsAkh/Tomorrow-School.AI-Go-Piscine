package piscine

func BasicAtoi2(s string) int {
	num := 0
	var char int
	for _, i := range s {
		if i >= '0' && i <= '9' {
			char = int(i - '0')
		} else {
			return 0
		}
		num = num*10 + char
	}
	return num
}
