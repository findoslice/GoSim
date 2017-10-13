package GoSim

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Particle struct {
	//group particle variables together
	initialHeight, initialVelocity, finalHeight, finalVelocity, maxHeight, horizontalRange, flightTime, theta, g float64
}

func (p *Particle) thetaDegrees() {
	p.theta = (p.theta / 360) * 2 * math.Pi
}

func (p *Particle) maximumHeight() {
	p.maxHeight = (math.Pow((p.initialVelocity*math.Sin(p.theta)), 2))/(2*p.g) + p.initialHeight
}

func (p *Particle) timeOfFlight() {
	a := math.Pow((p.initialVelocity * math.Sin(p.theta)), 2)
	b := 2 * p.g * p.initialHeight
	c := a + b
	p.flightTime = (math.Sqrt(c) + (p.initialVelocity * math.Sin(p.theta))) / p.g
}

func (p *Particle) maxRange() {
	if p.flightTime == 0 {
		p.timeOfFlight()
	}
	p.horizontalRange = (p.initialVelocity * math.Cos(p.theta)) * p.flightTime
}

func (p Particle) position(x float64) float64 {
	y := (x * math.Tan(p.theta)) - ((p.g * math.Pow(x, 2)) / (2 * math.Pow(p.initialVelocity, 2) * math.Pow((math.Cos(p.theta)), 2))) + p.initialHeight
	return y
}

func (p Particle) path() ([]float64, []float64) {
	var xs []float64
	var ys []float64
	for x := 0.0; x < p.horizontalRange; x += 0.1 {
		xs = append(xs, x)
		ys = append(ys, p.position(x))
	}
	return xs, ys
}

func (p Particle) pathPlot() {

	pts := make(plotter.XYs, int(p.horizontalRange)+2)
	for i := 0; float64(i) < p.horizontalRange; i++ {
		pts[i].X = float64(i)
		pts[i].Y = p.position(float64(i))
	}
	pts[int(p.horizontalRange)+1].X = p.horizontalRange
	pts[int(p.horizontalRange)+1].Y = 0.0
	pl, err := plot.New()
	if err != nil {
		panic(err)
	}
	pl.Title.Text = "Bleh"
	pl.X.Label.Text = "bluh"
	pl.Y.Label.Text = "blah"

	err = plotutil.AddLinePoints(pl, "First", pts)
	fmt.Println(pts)
	if err != nil {
		panic(err)
	}
	pl.X.Min = 0
	pl.Y.Min = 0
	if p.maxHeight > p.horizontalRange {
		pl.X.Max = pl.Y.Max
		fmt.Println(pl.X.Max, pl.Y.Max)
	} else {
		pl.Y.Max = pl.X.Max
		fmt.Println(pl.X.Max, pl.Y.Max)
	}
	if err := pl.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
