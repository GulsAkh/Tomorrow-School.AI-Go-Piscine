package piscine

func Rot14(s string) string {
	runes := []rune(s)

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r-'A'+14)%26
		}
	}
	// for _, r := range str {
	// 	z01.PrintRune(r)
	// }
	// z01.PrintRune('\n')
	return string(runes)
}

// H e l l o!          H o w a r e           Y o u?
// V s z z c!          V c k o f s           M c i?
// 72 101 108 108 111 72 111 119 97 114 101  89 111  117
// 86 115 122 122 99  86 99  107  111 102 115 77 99  105
// 14  14 14 14   12  14 12  12   14  12  14  12  12  12
