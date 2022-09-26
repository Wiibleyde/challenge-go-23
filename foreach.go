package student

func ForEach(f func(int), a []int) {
	for _, ele := range a {
		f(ele)
	}
}
