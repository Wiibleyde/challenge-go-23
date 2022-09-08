package student

func BasicAtoi(s string) int {
	nb := 0
	for _, ch := range s {
		nb = nb*10 + charToInt(ch)
	}
	return nb
}

func charToInt(char rune) int {
	count := 0
	if char < 48 && char > 58 {
		return 0
	}

	for i := '1'; i <= char; i++ {
		count++
	}

	return count
}
