package piscine

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}

	return num[2]
}

// func main() {
// 	middle := Abort(8206158666394987370, -7726695374529039112, 4647475675926194097, -5217061567791790716, -3643945018500837920)
// 	fmt.Println(middle)
// }
