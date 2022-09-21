package student

func SplitWhiteSpaces(s string) []string {
	tab := []string{}
	str := ""
	for _, j := range s {
		if j == ' ' || j == '\n' || j == '\t' {
			if str != "" {
				tab = append(tab, str)
			}
			str = ""
		} else {
			str = str + string(j)
		}
	}
	tab = append(tab, str)
	return tab
}
