package piscine

func Compact(ptr *[]string) int {
	count := 0

	for _, el := range *ptr {
		if el != " " {
			count++
		}
	}
	compacted := make([]string, 0, count)

	for _, el := range *ptr {
		if el != "" {
			compacted = append(compacted, el)
		}
	}
	*ptr = compacted
	return (len(*ptr))
}

// func main() {
// 	N := 6
// 	a := make([]string, N)
// 	a[0] = "a"
// 	a[2] = "b"
// 	a[4] = "c"

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("Size after compacting:", Compact(&a))

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
