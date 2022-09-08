package student

func StrLen(s string) int {
	counter := 0
	for _, char := range s {
		_ = char
		counter = counter + 1
	}
	return counter
}
