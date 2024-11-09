package piscine

func JumpOver(str string) string {
	runes := []rune(str)
	if str == "" {
		return "\n"
	}
	if len(str) < 3 {
		return "\n"
	}
	word := ""
	for i, el := range runes {
		if (i+1)%3 == 0 {
			word += string(el)
		}
	}
	return word + "\n"
}

// func main() {
// 	fmt.Print(JumpOver("1010101010"))
// 	fmt.Print(JumpOver(""))
// 	fmt.Print(JumpOver("t w e l v e"))
// 	//                  0 1 2 3 4 5
// 	fmt.Print(JumpOver("12"))
// }
