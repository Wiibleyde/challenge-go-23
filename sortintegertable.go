package student

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && table[j-1] > table[j]; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
	return table
}
