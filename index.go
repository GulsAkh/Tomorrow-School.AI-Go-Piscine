package piscine

func Index(s string, toFind string) int {
	var index int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			index = i
			break
		} else {
			index = -1
		}
	}
	return index
}
