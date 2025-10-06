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
	"lab8/ex4"
)

func main() {
	var (
		// RUN flags (ya los tienes)
		exercise int
		n        int
		nsRaw    string
		runs     int
		out      string

		// PLOT flags (nuevos/ajustados)
		mode            string
		inPlot, outPlot string
		title           string
		logx, logy      bool
		ymin            float64
		nolines         bool
		pSuccess        float64
		outDir          string
	)

	// -------- Flags (RUN) --------
	flag.IntVar(&exercise, "exercise", 1, "exercise to run (1|2|3)")
	flag.IntVar(&n, "n", 1000, "single input size n")
	flag.StringVar(&nsRaw, "ns", "", "comma-separated list of n values (e.g. 1,10,100)")
	flag.IntVar(&runs, "runs", 3, "repetitions for averaging")
	flag.StringVar(&out, "out", "", "CSV output (defaults to results/ex0X.csv)")

	flag.StringVar(&mode, "mode", "run", "run | plot")
	flag.StringVar(&inPlot, "inplot", "", "input CSV (when -mode=plot)")
	flag.StringVar(&outPlot, "outplot", "plots/out.png", "output PNG (when -mode=plot)")
	flag.StringVar(&title, "title", "", "plot title (when -mode=plot)")
	flag.BoolVar(&logx, "logx", false, "log scale on X (when -mode=plot)")
	flag.BoolVar(&logy, "logy", false, "log scale on Y (when -mode=plot)")
	flag.Float64Var(&ymin, "ymin", 0, "minimum Y value (>0) for plotting (when -mode=plot)")
	flag.BoolVar(&nolines, "nolines", false, "plot points only (no connecting line)")
	flag.Float64Var(&pSuccess, "p", 1.0, "success probability for average mixed (0..1) (mode=gen-linear)")
	flag.StringVar(&outDir, "outdir", "results/ex04", "output directory for linear search CSVs (mode=gen-linear)")
	flag.Parse()

	// ---- en tu manejo de modos, antes del modo "run" ----
	if mode == "gen-linear" {
		// parse ns (reuse your existing ns parsing)
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
		if err := ex4.GenerateCSVs(ns, pSuccess, outDir); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Linear search CSVs written to:", outDir)
		return
	}

	// ---------- PLOT mode ----------
	if mode == "plot" {
		if strings.TrimSpace(inPlot) == "" {
			log.Fatal("missing -inplot=<csv>")
		}
		if strings.TrimSpace(title) == "" {
			title = "n vs time (ms)"
		}
		// drawLine es !nolines
		if err := config.PlotCSVWithOpts(inPlot, outPlot, title, "n", "avg_ms", logx, logy, ymin, !nolines); err != nil {
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
