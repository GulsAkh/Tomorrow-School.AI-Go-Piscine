package piscine

func Unmatch(a []int) int {
	// Map to store the frequency of each element
	c := make(map[int]int)
	for _, num := range a {
		c[num]++
	}

	// Check for elements with odd occurrences
	for _, count := range c {
		if c[count]%2 != 0 {
			return count // Return the first element with an odd count (no pair)
		}
	}

	// If all elements have pairs, return -1
	return -1
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
