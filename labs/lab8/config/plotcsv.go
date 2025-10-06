package config

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type PlotRow struct {
	N     float64
	AvgMs float64
}

func ensureDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// LoadCSV expects header: exercise,n,avg_ms,runs,note
func LoadCSV(path string) ([]PlotRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) <= 1 {
		return nil, fmt.Errorf("no data rows in %s", path)
	}
	out := make([]PlotRow, 0, len(records)-1)
	for i, rec := range records {
		if i == 0 { // skip header
			continue
		}
		if len(rec) < 3 { // need n and avg_ms
			continue
		}
		nVal, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			continue
		}
		tVal, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			continue
		}
		out = append(out, PlotRow{N: nVal, AvgMs: tVal})
	}
	return out, nil
}

// PlotCSV builds a PNG line+scatter chart from the CSV.
func PlotCSV(inCSV, outPNG, title, xLabel, yLabel string, logX, logY bool) error {
	rows, err := LoadCSV(inCSV)
	if err != nil {
		return err
	}

	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	if logX {
		p.X.Scale = plot.LogScale{}
		p.X.Tick.Marker = plot.LogTicks{}
	}
	if logY {
		p.Y.Scale = plot.LogScale{}
		p.Y.Tick.Marker = plot.LogTicks{}
	}

	pts := make(plotter.XYs, 0, len(rows))
	for _, r := range rows {
		x := r.N
		y := r.AvgMs
		if logX && x <= 0 {
			x = 1
		}
		if logY && y <= 0 {
			y = math.SmallestNonzeroFloat64
		}
		pts = append(pts, plotter.XY{X: x, Y: y})
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	scat, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	line.Color = color.Black
	scat.Radius = vg.Points(2)

	p.Add(line, scat)

	if err := ensureDir(outPNG); err != nil {
		return err
	}
	return p.Save(6*vg.Inch, 4*vg.Inch, outPNG)
}
