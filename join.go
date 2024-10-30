package piscine

func Join(strs []string, sep string) string {
	str := ""

	for i, el := range strs {
		if i != (len(strs) - 1) {
			str += string(el) + sep
		} else {
			str += string(el)
		}
	}
	return string
}
