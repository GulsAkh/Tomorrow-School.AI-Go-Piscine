package piscine

func StrRev(s string) {
	var r []rune

	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, rune(s[i]))
	}
	_ = string(r)
}

/*
func main() {
	s := "Hello World!"
	StrRev(s)
	fmt.Println(s)
}*/
