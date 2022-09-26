package student

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, ele := range a {
		result = append(result, f(ele))
	}
	return result
}
