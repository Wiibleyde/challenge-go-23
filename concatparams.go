package student

func ConcatParams(args []string) string {
	finalStr := ""
	for ind, ele := range args {
		if ind == len(args)-1 {
			finalStr = finalStr + ele
		} else {
			finalStr = finalStr + ele + "\n"
		}
	}
	return finalStr
}
