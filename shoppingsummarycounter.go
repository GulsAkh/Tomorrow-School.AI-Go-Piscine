package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	m := make(map[string]int)
	if str == "" {
		m[string("")]++
		return m
	}
	str1 := ""
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == ' ' {
			if str != "" {
				m[str1]++
				str1 = ""
			}
		} else {
			str1 += string(char)
		}
	}

	if str != "" {
		m[str1]++
	}

	return m
}

// func main() {
// 	summary := ""
// 	for index, element := range ShoppingSummaryCounter(summary) {
// 		fmt.Println(index, "=>", element)
// 	}
// }
