package student

func Atoi(s string) int {
	if StrLen(s) == 0 {
		return 0
	}

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if StrLen(s) < 1 {
			return 0
		}
	}

	nb := 0

	for _, ch := range s {
		if !containsIn0to9(ch) {
			return 0
		}
		nb = nb*10 + charToInt(ch)
	}

	if s0[0] == '-' {
		nb *= -1
	}
	return nb
}

func StrLen(s string) int {
	counter := 0
	for _, ele := range s {
		_ = ele
		counter = counter + 1
	}
	return counter
}

func containsIn0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}

	return false
}


func charToInt(ele rune) int {
	count := 0
	if ele < 48 && ele > 58 {
		return 0
	}

	for i := '1'; i <= ele; i++ {
		count++
	}

	return count
}
