package student

func IsNegative(nb int) {
	if nb < 0 {
		print(true)
	} else {
		print(false)
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
