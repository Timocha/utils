package utils

//Contains ... return true when slice a contains string x
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//ContainsInt ... return true when slice a contains int x
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
