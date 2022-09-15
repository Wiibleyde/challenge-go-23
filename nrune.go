package student

func NRune(s string, n int) rune {
	taille := StrLen(s)
	if n <= 0 || n > taille {
		return 0
	}
	chars := []rune(s)
	return chars[n-1]
}
