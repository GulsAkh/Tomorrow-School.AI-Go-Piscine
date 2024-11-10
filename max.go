package piscine

func Max(a []int) int {
	var maX int
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			maX = a[i]
		}
	}
	return maX
}

// func main() {
// 	a := []int{23, 123, 1, 11, 55, 93}
// 	max := Max(a)
// 	fmt.Println(max)
// }
