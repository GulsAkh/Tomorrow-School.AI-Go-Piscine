package piscine

func Atoi(s string) int {
	num := 0
	var char int
	sign := 1

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		} else {
			char = int(i - '0')
		}
		num = num*10 + char
	}
	return num * sign
}
