package student

func IsNumeric(s string) bool {
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
