package piscine

func IsLower(s string) bool {
	lower := true
	for _, el := range s {
		if el < 'a' || el > 'z' {
			lower = false
		}
	}
	return lower
}
