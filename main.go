package main

import (
	"image/color"
	"log"
	"math"

	"github.com/BoB1Edition/gobrain/neuron"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	n := neuron.NewNeuron(2)
	variables := [][]float64{{0.0, 0.0}, {1.0, 0.0}, {0.0, 1.0}, {1.0, 1.0}}
	answer := []float64{0, 0, 0, 1}
	p := plot.New()
	p.Title.Text = "histogram plot"
	values := make(plotter.XYs, 0)
	values2 := make(plotter.XYs, 0)
	for j := 0; j < 250; j++ {
		for i := range variables {
			ans, err := n.Forward(variables[i])
			if err != nil {
				log.Fatal(err)
			}
			if ans != answer[i] {
				n.Adjustment(answer[i])
			}
			//values := plotter.XYs{{0, 0}, {0, 1}, {0.5, 1}, {0.5, 0.6}, {0, 0.6}}
			values = append(values, plotter.XY{
				X: float64((j * 4) + i),
				Y: math.Abs(answer[i] - ans),
			})
			values2 = append(values, plotter.XY{
				X: float64((j * 4) + i),
				Y: answer[i] - ans,
			})
			//log.Println()
		}
		if j%1000 == 0 {
			log.Printf("j: %d", j)
		}
	}
	hist, err := plotter.NewLine(values)
	if err != nil {
		panic(err)
	}
	hist.Color = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 127,
	}
	hist2, err := plotter.NewLine(values2)
	if err != nil {
		panic(err)
	}
	hist2.Color = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 127,
	}
	p.Add(hist)
	p.Add(hist2)
	p.Save(30*vg.Inch, 5*vg.Inch, "hist.png")

}
