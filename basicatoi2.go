package student

func BasicAtoi2(s string) int {
	nb := 0
	for _, ch := range s {
		nb = nb*10 + charToInt2(ch)
		if charToInt2(ch) == 0 {
			return 0
		}
	}
	return nb
}


func charToInt2(char rune) int {
	count := 0
	if char < 48 || char > 58 {
		return 0
	}
	for i := '1'; i <= char; i++ {
		count++
	}
	return count
}
