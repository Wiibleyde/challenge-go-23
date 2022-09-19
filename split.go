package student

func Split(s, sep string) []string {
	char_len := StrLen(sep)
	str_len := StrLen(s)
	size := 0
	for ind := 0; ind <= str_len-char_len; ind++ {
		if string(s[ind:ind+char_len]) == sep {
			size++
		}
	}
	resArr := make([]string, size+1)
	i := 0
	start := 0
	ind := 0
	for ; ind <= str_len-char_len; ind++ {
		if string(s[ind:ind+char_len]) == sep {
			resArr[i] = string(s[start:ind])
			i++
			start = ind + char_len
		}
		if ind == str_len-char_len {
			resArr[i] = string(s[start:])
		}
	}
	return resArr
}
