package ex3

// Ex3 implements Exercise 3: i in [1..n/3], j in [1..n] step 4.
// Time complexity: O(n^2), space O(1).
func Ex3(n int) uint64 {
	var counter uint64
	for i := 1; i <= n/3; i++ {
		for j := 1; j <= n; j += 4 {
			counter++
		}
	}
	return counter
}
