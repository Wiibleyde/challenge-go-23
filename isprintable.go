package student

func IsPrintable(s string) bool {
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) >= 0 && rune(s[i]) <= 32 {
			return false
		}
	}
	return true
}
