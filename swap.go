package student

func Swap(a *int, b *int) {
	tempB := *b
	*b = *a
	*a = tempB
}
