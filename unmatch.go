package piscine

func Unmatch(a []int) int {
	c := make(map[int]int)
	for _, num := range a {
		c[num]++
	}

	for _, count := range c {
		if c[count]%2 != 0 {
			return count // Return the first element with an odd count (no pair)
		}
	}

	return -1
}

// func main() {
// 	a := []int{1, 1, 2, 4, 3, 4, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
