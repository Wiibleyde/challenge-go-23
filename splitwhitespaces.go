package student

func SplitWhiteSpaces(str string) []string {
	ind1 := 0
	noWS := false
	for index := range str {
		if noWS && index != 0 && (str[index-1] == '\n' || str[index-1] == '\t' || str[index-1] == ' ') && str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			ind1++
		}
		if str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			noWS = true
		}
	}
	ind1++
	x := 0
	ans := make([]string, ind1)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		ans[x] = ans[x] + string(c)
		ok = false
	}
	return ans
}
