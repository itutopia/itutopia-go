package calc

//func Dec(a, b int) int {
//	return a - b
//}

func Dec(a, b int) int {
	res := a - b
	if res < 0 {
		return 0
	}
	return res
}
