package student

func Sqrt(nb int) int {
	for i := 1; i <= 99999; i++ {
		if nb/i == i {
			if nb%i == 0 {
				return i
			}
		}
	}
	return 0
}
