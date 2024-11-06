package piscine

func Map(f func(int) bool, a []int) []bool {
	var bLst []bool
	for _, el := range a {
		bLst = append(bLst, f(el))
	}
	return bLst
}
