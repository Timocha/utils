package utils

//InSlice ... return true when slice a contains string x
func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//InSliceInt ... return true when slice a contains int x
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
