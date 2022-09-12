package student

func FindNextPrime(nb int) int {
	for i := nb; i < 99999999999; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return nb
}
