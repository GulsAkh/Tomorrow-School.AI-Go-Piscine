package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1

	for _, el := range s {
		if el >= '1' && el <= '9' {
			num = num*10 + int(el-'0')
		} else if el == '-' {
			sign = -1
		}
	}
	return num * sign
}
