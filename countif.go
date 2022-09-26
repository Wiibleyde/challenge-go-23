package student

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, ele := range tab {
		if f(ele) {
			counter++
		}
	}
	return counter
}
