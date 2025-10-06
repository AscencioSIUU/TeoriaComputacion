package ex4

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"lab8/config" // <- ajusta al module path real si es necesario
)

// Deterministic formulas (comparisons, not time)
func BestComparisons(n int) float64       { return 1 }
func WorstComparisons(n int) float64      { return float64(n) }
func AvgComparisonsSuccess(n int) float64 { return float64(n+1) / 2 } // x está y es uniforme
func AvgComparisonsMixed(n int, p float64) float64 {
	// Éxito con prob. p (uniforme en 1..n), fracaso con prob. (1-p)
	return p*float64(n+1)/2 + (1-p)*float64(n)
}

func ensureDir(dir string) error {
	if dir == "" || dir == "." {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// GenerateCSVs writes four CSVs into outDir: best, avg_success, avg_mixed_pXX, worst.
func GenerateCSVs(ns []int, p float64, outDir string) error {
	if err := ensureDir(outDir); err != nil {
		return err
	}
	header := []string{"exercise", "n", "avg_ms", "runs", "note"} // keep same format; "avg_ms" holds comparisons

	// helpers
	write := func(name string, label string, val func(int) float64) error {
		path := filepath.Join(outDir, name)
		if err := os.RemoveAll(path); err != nil { /* ignore */
		}
		for idx, n := range ns {
			row := []string{
				label,
				strconv.Itoa(n),
				fmt.Sprintf("%.6f", val(n)), // comparisons
				"1",
				"",
			}
			h := header
			if idx > 0 {
				h = nil // header already written by AppendCSV
			}
			if err := config.AppendCSV(path, h, [][]string{row}); err != nil {
				return err
			}
		}
		return nil
	}

	// files
	if err := write("ex04_best.csv", "ex04-best", BestComparisons); err != nil {
		return err
	}
	if err := write("ex04_avg_success.csv", "ex04-avg-success", AvgComparisonsSuccess); err != nil {
		return err
	}
	nameMixed := "ex04_avg_mixed_p" + strconv.FormatFloat(p, 'f', 2, 64) + ".csv"
	if err := write(nameMixed, "ex04-avg-mixed", func(n int) float64 { return AvgComparisonsMixed(n, p) }); err != nil {
		return err
	}
	if err := write("ex04_worst.csv", "ex04-worst", WorstComparisons); err != nil {
		return err
	}

	return nil
}
