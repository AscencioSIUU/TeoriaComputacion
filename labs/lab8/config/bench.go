package config

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

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
	// Evita que el compilador elimine el trabajo
	if sink == math.MaxUint64 {
		fmt.Fprintln(os.Stderr, "ignore:", sink)
	}
	// PRECISIÓN: usa ns -> ms con decimales (no Duration.Milliseconds())
	avgMs := (float64(total.Nanoseconds()) / 1e6) / float64(runs)

	return Result{
		N:     n,
		AvgMs: avgMs,
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
