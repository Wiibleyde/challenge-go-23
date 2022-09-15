package student

func IsLower(s string) bool {
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
