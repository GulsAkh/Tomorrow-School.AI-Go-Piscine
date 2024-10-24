package main

func StrRev(s string) string {
	var r []rune

	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, rune(s[i]))
	}
	return string(r)
}
