package piscine

func BasicJoin(elems []string) string {
	str := ""

	for _, el := range elems {
		str += string(el)
	}
	return str
}
