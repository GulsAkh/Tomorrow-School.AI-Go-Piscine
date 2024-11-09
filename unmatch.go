package piscine

func Unmatch(a []int) int {
	c := make(map[int]int)
	for _, num := range a {
		c[num]++
	}
	var num int
	for _, count := range c {
		if c[count]%2 != 0 {
			num = count // Return the first element with an odd count (no pair)
		} else {
			num = -1
		}
	}

	return num
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
