package piscine

func ReverseMenuIndex(menu []string) []string {
	num := make([]string, len(menu))

	for i, el := range menu {
		num[len(menu)-1-i] = el
	}
	return num
}

// func main() {
// 	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
// }
