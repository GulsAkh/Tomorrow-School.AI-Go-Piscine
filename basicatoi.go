package piscine

func BasicAtoi(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		num = num*10 + digit
	}
	return num
}
