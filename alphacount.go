package student

func AlphaCount(s string) int {
	// taille := StrLen(s)
	nb := 0
	for _, ele := range s {
		if rune(ele) >= 65 && rune(ele) <= 90 {
			nb = nb + 1
		}
		if rune(ele) >= 97 && rune(ele) <= 122 {
			nb = nb + 1
		}
	}
	return nb
}
