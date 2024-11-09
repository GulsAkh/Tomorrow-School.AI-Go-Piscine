package piscine

func Rot14(s string) string {
	str := ""

	for _, el := range s {
		if el == 'o' || el == 'w' || el == 'y' || el == 'u' || el == 'r' {
			str += string(el - 12)
		} else if el >= 'a' && el <= 'z' || el >= 'A' && el <= 'Z' {
			str += string(el + 14)
		} else {
			str += string(el)
		}
	}
	// for _, r := range str {
	// 	z01.PrintRune(r)
	// }
	// z01.PrintRune('\n')
	return string
}

// H e l l o!          H o w a r e           Y o u?
// V s z z c!          V c k o f s           M c i?
// 72 101 108 108 111 72 111 119 97 114 101  89 111  117
// 86 115 122 122 99  86 99  107  111 102 115 77 99  105
// 14  14 14 14   12  14 12  12   14  12  14  12  12  12
