package student

func FindNextPrime(nb int) int {
	for i := nb + 1; i < 9999999; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return nb
}
