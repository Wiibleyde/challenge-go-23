package student

func TrimAtoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}

	var negati bool
	for _, ch := range s {
		if containsIn0to9(ch) {
			break
		}

		if ch == '-' {
			negati = true
		}
	}

	nb := 0

	for _, ch := range s {
		if containsIn0to9(ch) {
			nb = nb*10 + charToInt(ch)
		}
	}

	if negati {
		nb *= -1
	}
	return nb
}
