package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"lab8/config" // <- usa tu module path real
	"lab8/ex1"
	"lab8/ex2"
	"lab8/ex3"
)

func main() {
	var (
		// RUN mode flags
		exercise int
		n        int
		nsRaw    string
		runs     int
		out      string

		// PLOT mode flags
		mode            string
		inPlot, outPlot string
		title           string
		logx, logy      bool
	)

	// -------- Flags (RUN) --------
	flag.IntVar(&exercise, "exercise", 1, "exercise to run (1|2|3)")
	flag.IntVar(&n, "n", 1000, "single input size n")
	flag.StringVar(&nsRaw, "ns", "", "comma-separated list of n values (e.g. 1,10,100)")
	flag.IntVar(&runs, "runs", 3, "repetitions for averaging")
	flag.StringVar(&out, "out", "", "CSV output (defaults to results/ex0X.csv)")

	// -------- Flags (PLOT) -------
	flag.StringVar(&mode, "mode", "run", "run | plot")
	flag.StringVar(&inPlot, "inplot", "", "input CSV (when -mode=plot)")
	flag.StringVar(&outPlot, "outplot", "plots/out.png", "output PNG (when -mode=plot)")
	flag.StringVar(&title, "title", "", "plot title (when -mode=plot)")
	flag.BoolVar(&logx, "logx", false, "log scale on X (when -mode=plot)")
	flag.BoolVar(&logy, "logy", false, "log scale on Y (when -mode=plot)")

	flag.Parse()

	// ---------- PLOT mode ----------
	if mode == "plot" {
		if strings.TrimSpace(inPlot) == "" {
			log.Fatal("missing -inplot=<csv>")
		}
		if strings.TrimSpace(title) == "" {
			title = "n vs time (ms)"
		}
		if err := config.PlotCSV(inPlot, outPlot, title, "n", "avg_ms", logx, logy); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Plot saved to", outPlot)
		return
	}

	// ---------- RUN mode ----------
	// Default CSV per exercise if user didn't pass -out
	if strings.TrimSpace(out) == "" {
		switch exercise {
		case 1:
			out = "results/ex01.csv"
		case 2:
			out = "results/ex02.csv"
		case 3:
			out = "results/ex03.csv"
		default:
			out = "results/unknown.csv"
		}
	}

	// Pick runner
	var runner config.Runner
	switch exercise {
	case 1:
		runner = ex1.Ex1
	case 2:
		runner = ex2.Ex2
	case 3:
		runner = ex3.Ex3
	default:
		fmt.Fprintln(os.Stderr, "use -exercise=1, -exercise=2, or -exercise=3")
		os.Exit(2)
	}

	// Parse ns or fallback to single n
	ns := []int{n}
	if strings.TrimSpace(nsRaw) != "" {
		parts := strings.Split(nsRaw, ",")
		ns = ns[:0]
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			if strings.ContainsAny(p, "eE") {
				f, err := strconv.ParseFloat(p, 64)
				if err != nil {
					log.Fatal(err)
				}
				ns = append(ns, int(f))
			} else {
				v, err := strconv.Atoi(p)
				if err != nil {
					log.Fatal(err)
				}
				ns = append(ns, v)
			}
		}
	}

	// Measure and write CSV incrementally
	exLabel := fmt.Sprintf("ex%02d", exercise)
	header := []string{"exercise", "n", "avg_ms", "runs", "note"}

	for idx, size := range ns {
		res := config.TimeN(runner, size, runs)
		fmt.Printf("[%s] n=%d avg=%.3fms runs=%d\n", exLabel, res.N, res.AvgMs, res.Runs)

		row := []string{
			exLabel,
			strconv.Itoa(res.N),
			fmt.Sprintf("%.3f", res.AvgMs),
			strconv.Itoa(res.Runs),
			"",
		}
		if err := config.AppendCSV(out, header, [][]string{row}); err != nil {
			log.Fatal(err)
		}
		if idx == 0 {
			header = nil
		}
	}
}
