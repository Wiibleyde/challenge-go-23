package student

func AppendRange(min, max int) []int {
	finalList:=[]int(nil)
	for i:=min;i<max;i++ {
		finalList = append(finalList, i)
	}
	return finalList
}