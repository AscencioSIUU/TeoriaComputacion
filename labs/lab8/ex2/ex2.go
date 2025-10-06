package ex2

// Ex2 implements Exercise 2 with an early break in the inner loop.
// Time complexity: O(n) for n > 1; O(1) when n <= 1.
func Ex2(n int) uint64 {
	if n <= 1 {
		return 0
	}
	var counter uint64
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			counter++ // simulate "printf" + break
			break
		}
	}
	return counter
}
