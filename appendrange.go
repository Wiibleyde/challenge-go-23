package student

func AppendRange(min, max int) []int {
	finalListe := []int(nil)
	for i := min; i < max; i++ {
		finalListe = append(finalListe, i)
	}
	return finalListe
}
