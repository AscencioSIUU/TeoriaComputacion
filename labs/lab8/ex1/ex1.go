package ex1

func Ex1(n int) uint64 {
	var counter uint64
	for i := n / 2; i <= n; i++ {
		for j := 1; j+n/2 <= n; j++ {
			for k := 1; k <= n; k = k * 2 {
				// Work: simple increment to keep the loop non-trivial
				counter++
			}
		}
	}
	return counter
}
