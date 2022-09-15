package student

func Index(s string, toFind string) int {
	tailleStr := StrLen(s)
	tailleFind := StrLen(toFind)
	for i := 0; i < tailleStr-tailleFind; i++ {
		if s[i:i+tailleFind] == toFind {
			return i
		}
	}
	return -1
}
