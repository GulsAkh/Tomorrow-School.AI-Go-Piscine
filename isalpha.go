package piscine

func IsAlpha(s string) bool {
	flag := true
	for _, el := range s {
		if (el < 'a' || el > 'z') && (el < 'A' || el > 'Z') && (el < '0' || el > '9') {
			flag = false
			// break
		}
	}
	return flag
}
