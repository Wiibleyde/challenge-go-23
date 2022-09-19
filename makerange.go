package student

func MakeRange(min, max int) []int {
	if max-min < 1 {
		return []int(nil)
	}
	finalList := make([]int, max-min)
	for i := 0; i < len(finalList); i++ {
		finalList[i] = i + min
	}
	return finalList
}
