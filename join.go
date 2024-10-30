package piscine

func Join(strs []string, sep string) string {
	str := ""

	for _, el := range strs {
		str += string(el) + sep
	}
	return str
}
