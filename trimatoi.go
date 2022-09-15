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

	nm := 0

	for _, ch := range s {
		if containsIn0to9(ch) {
			nm = nm*10 + charToInt(ch)
		}
	}

	if negati {
		nm *= -1
	}
	return nm
}
