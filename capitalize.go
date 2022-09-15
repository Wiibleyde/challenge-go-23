package student

func Capitalize(s string) string {
	nCapt := ""
	for i := 0; i < StrLen(s); i++ {
		if rune(s[i]) == 32 {
			nCapt = nCapt + string(rune(32))
			nCapt = nCapt + ToUpper(string(s[i+1]))
			i = i + 1
		} else if rune(s[i]) == 43 {
			nCapt = nCapt + string(rune(43))
			nCapt = nCapt + ToUpper(string(s[i+1]))
			i = i + 1
		} else if rune(s[i]) == 42 {
			nCapt = nCapt + string(rune(42))
			nCapt = nCapt + ToUpper(string(s[i+1]))
			i = i + 1
		} else if rune(s[i]) == 58 {
			nCapt = nCapt + string(rune(58))
			nCapt = nCapt + ToUpper(string(s[i+1]))
			i = i + 1
		} else {
			nCapt = nCapt + string(s[i])
		}
	}
	return nCapt
}
