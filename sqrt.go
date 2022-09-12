package student

func Sqrt(nb int) int {
	if nb <= 1 {
		return nb
	}
	for i := 1; i <= 99999; i++ {
		if nb/i == i && nb%i == 0 {
			return i
		}
	}
	return 0
}
