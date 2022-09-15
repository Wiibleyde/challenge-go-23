package student

func IsAlpha(s string) bool {
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			continue
		} else if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			continue
		} else if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
