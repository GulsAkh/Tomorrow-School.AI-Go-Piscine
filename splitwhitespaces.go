package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	var str []string
	fstring := ""

	for _, el := range runes {
		if el == ' ' || el == '\t' || el == '\n' {
			if fstring != "" {
				str = append(str, fstring)
				fstring = ""
			}
		} else {
			fstring += string(el)
		}
	}
	if fstring != "" {
		str = append(str, fstring)
	}
	return str
}
