package student

func Join(strs []string, sep string) string {
	nStr := ""
	for ind, ele := range strs {
		if ind == len(strs)-1 {
			nStr = nStr + ele
		} else {
			nStr = nStr + ele + sep
		}
	}
	return nStr
}
