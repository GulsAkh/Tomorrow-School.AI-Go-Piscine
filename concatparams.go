package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, el := range args {
		if i != len(args)-1 {
			str += el + "\n"
		} else {
			str += el
		}
	}
	return str
}
