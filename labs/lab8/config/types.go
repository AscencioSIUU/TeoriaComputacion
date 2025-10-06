package config

// Runner is the function signature for an exercise implementation.
// It must return a counter to avoid dead-code elimination by the compiler.
type Runner func(n int) uint64

// Result is the timing summary for a single input size.
type Result struct {
	N     int
	AvgMs float64
	Runs  int
	Note  string
}
