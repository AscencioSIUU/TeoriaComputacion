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
		if i == 0 {
			continue // header
		}
		if len(rec) < 3 {
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

func ensureDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// PlotCSVWithOpts: agrega ymin (>0) para evitar “pared” en 0 y opción sin líneas.
func PlotCSVWithOpts(inCSV, outPNG, title, xLabel, yLabel string,
	logX, logY bool, yMin float64, drawLine bool) error {

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

	// Si se define un mínimo de Y, úsalo para evitar el “manchón” en 0
	if yMin > 0 {
		p.Y.Min = yMin
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
		if yMin > 0 && y < yMin {
			y = yMin
		}
		pts = append(pts, plotter.XY{X: x, Y: y})
	}

	if drawLine {
		line, err := plotter.NewLine(pts)
		if err != nil {
			return err
		}
		line.Color = color.Black
		p.Add(line)
	}

	scat, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	scat.Radius = vg.Points(2)
	p.Add(scat)

	if err := ensureDir(outPNG); err != nil {
		return err
	}
	return p.Save(6*vg.Inch, 4*vg.Inch, outPNG)
}
