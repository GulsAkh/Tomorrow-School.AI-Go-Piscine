package piscine

func SortIntegerTable(table []int) {
	num := len(table)

	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
