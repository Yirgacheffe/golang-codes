package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	// github.com/vdobler/chart
	// github.com/wcharczuk/go-chart
)

func main() {
	var values plotter.Values
	for i := 0; i < 1000; i++ {
		values = append(values, rand.NormFloat64())
	}
	histPlot(values)
	barPlot(values[:4])
	boxPlot(values)
}

func histPlot(values plotter.Values) {
	p := plot.New()
	p.Title.Text = "histogram plot"

	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}

	p.Add(hist)
	if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}

func barPlot(values plotter.Values) {
	p := plot.New()
	p.Title.Text = "bar plot"

	bar, err := plotter.NewBarChart(values, 15)
	if err != nil {
		panic(err)
	}

	p.Add(bar)
	if err := p.Save(3*vg.Inch, 3*vg.Inch, "bar.png"); err != nil {
		panic(err)
	}
}

func boxPlot(values plotter.Values) {
	p := plot.New()
	p.Title.Text = "box plot"

	box, err := plotter.NewBoxPlot(vg.Length(15), 0.0, values)
	if err != nil {
		panic(err)
	}

	p.Add(box)
	if err := p.Save(3*vg.Inch, 3*vg.Inch, "box.png"); err != nil {
		panic(err)
	}
}
