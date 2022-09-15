package student

func IsUpper(s string) bool {
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			continue
		} else {
			return false
		}
	}
	return true
}
