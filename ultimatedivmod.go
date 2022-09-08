package student

func UltimateDivMod(a *int, b *int) {
	solution := *a / *b
	*b = *a % *b
	*a = solution
}
