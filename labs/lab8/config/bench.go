package config

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

// TimeN runs r 'runs' times for input n and returns the average millis.
func TimeN(r Runner, n, runs int) Result {
	if runs < 1 {
		runs = 1
	}
	var total time.Duration
	var sink uint64
	for i := 0; i < runs; i++ {
		start := time.Now()
		sink += r(n)
		total += time.Since(start)
	}
	// prevent compiler from optimizing away the work
	if sink == math.MaxUint64 {
		fmt.Fprintln(os.Stderr, "ignore:", sink)
	}
	return Result{
		N:     n,
		AvgMs: (float64(total.Nanoseconds()) / 1e6) / float64(runs),
		Runs:  runs,
		Note:  "",
	}
}

func EnsureDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// AppendCSV appends rows to CSV and writes header if the file is new.
func AppendCSV(path string, header []string, rows [][]string) error {
	_, statErr := os.Stat(path)
	fileExists := statErr == nil

	if err := ensureDir(path); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if !fileExists && len(header) > 0 {
		if err := w.Write(header); err != nil {
			return err
		}
	}
	for _, r := range rows {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}
