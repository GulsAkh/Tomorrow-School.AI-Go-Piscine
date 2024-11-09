package main

func DescendAppendRange(max, min int) []int {
	var num []int
	if max < min {
		return num
	}
	for i := max; i > min; i-- {
		num = append(num, i)
	}
	return num
}

// func main() {
// 	fmt.Println(DescendAppendRange(10, 5))
// 	fmt.Println(DescendAppendRange(5, 10))
// }
