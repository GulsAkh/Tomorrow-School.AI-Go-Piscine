package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	flag := false

	for _, el := range s {
		if el >= '1' && el <= '9' {
			num = num*10 + int(el-'0')
			flag = true
		} else if el == '-' && !flag {
			sign = -1
		}
	}
	return num * sign
}
