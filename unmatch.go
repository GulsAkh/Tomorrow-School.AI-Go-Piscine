package piscine

func Unmatch(a []int) int {
	c := make(map[int]int)
	for _, num := range a {
		c[num]++
	}

	var num int
	for i, el := range c {
		if el%2 != 0 {
			num = i
		}
	}
	return num
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
