package piscine

func IsNumeric(s string) bool {
	for _, el := range s {
		if el < '0' || el > '9' {
			return false
		}
	}
	return true
}
