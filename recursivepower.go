package student

func RecursivePower(nb int, power int) int {
	a := nb
	i := power
	if power < 0 {
		return 0
	} else {
		if power == 0 {
			return 1
		}
		if i > 1 {
			a = a * RecursivePower(nb, i-1)
		}
		return a
	}
}
