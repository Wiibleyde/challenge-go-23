package student

func Any(f func(string) bool, a []string) bool {
	for _, ele := range a {
		if f(ele) {
			return true
		}
	}
	return false
}
