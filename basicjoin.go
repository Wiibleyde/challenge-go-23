package student

func BasicJoin(elems []string) string {
	nString := ""
	for _, ele := range elems {
		nString = nString + ele
	}
	return nString
}
