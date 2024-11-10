package piscine

func RockAndRoll(n int) string {
	str := ""
	if n < 0 {
		str += "number is negative\n"
	}
	if n%2 == 0 && n%3 == 0 {
		str += "rock and roll\n"
	} else if n%2 == 0 {
		str += "rock\n"
	} else if n%3 == 0 {
		str += "roll\n"
	} else {
		str += "non divisible\n"
	}
	return str
}

// func main() {
// 	fmt.Println(RockAndRoll(4))
// 	fmt.Println(RockAndRoll(9))
// 	fmt.Println(RockAndRoll(6))
// }
