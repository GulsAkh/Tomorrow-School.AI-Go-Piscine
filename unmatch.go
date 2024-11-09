package piscine

func Unmatch(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}

	return -1
}

// func main() {
// 	a := []int{1, 1, 2, 4, 3, 4, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
