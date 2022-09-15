package student

func LastRune(s string) rune {
	chars := []rune(s)
	taille := StrLen(s)
	return chars[taille-1]
}
