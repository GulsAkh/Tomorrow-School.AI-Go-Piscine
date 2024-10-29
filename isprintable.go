package piscine

func IsPrintable(s string) bool {
	for _, el := range s {
		if el < 32 || el > 126 {
			return false
		}
	}
	return true
}
