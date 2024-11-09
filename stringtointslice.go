package piscine

func StringToIntSlice(str string) []int {
	var num []int

	for _, el := range str {
		num = append(num, int(el))
	}
	return num
}

// func main() {
// 	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
// 	fmt.Println(StringToIntSlice("Converted this string into an int"))
// 	fmt.Println(StringToIntSlice("hello THERE"))
// }
