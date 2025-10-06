package ex4

// LinearSearch returns the index of x in a, or -1 if not found.
// It also returns the number of equality comparisons performed.
func LinearSearch(a []int, x int) (idx int, comparisons int) {
	for i, v := range a { // scans left-to-right
		comparisons++
		if v == x {
			return i, comparisons
		}
	}
	return -1, comparisons
}
