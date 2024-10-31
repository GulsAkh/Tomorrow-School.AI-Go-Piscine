package piscine

func Split(s, sep string) []string {
	var str []string
	mainstr := len(s)
	substr := len(sep)
	just_str := ""
	for i := 0; i < mainstr; i++ {
		if i <= mainstr-substr && s[i:i+substr] == sep {
			if just_str != "" {
				str = append(str, just_str)
			}
			just_str = ""
			i += substr - 1
		} else {
			just_str += string(s[i])
		}
	}
	if just_str != "" {
		str = append(str, just_str)
	}
	return str
}
