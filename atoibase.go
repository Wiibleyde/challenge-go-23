package student

func AtoiBase(s string, t string) int {
	resultat := 0
	taille := 0
	liste := map[rune]bool{}
	for _, c := range t {
		if liste[c] || c == '-' || c == '+' {
			return resultat
		}
		liste[c] = true
		taille++
	}
	if taille > 1 {
		for _, v := range s {
			count := 0
			if liste[v] {
				for _, j := range t {
					if j == v {
						break
					}
					count++
				}
				resultat = resultat*taille + count
			}
		}
	}
	return resultat
}
