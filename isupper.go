package piscine

func IsUpper(s string) bool {
	for _, el := range s {
		if el < 'A' || el > 'Z' {
			return false
		}
	}
	return true
}
