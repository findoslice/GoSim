package main

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
	InitialHeight, InitialVelocity, FinalHeight, FinalVelocity, MaxHeight, HorizontalRange, FlightTime, Theta, G float64
	TerminalVelocity                                                                                             float64
}

func (p *Particle) ThetaDegrees() {
	p.Theta = float64(int((p.Theta/360)*2*math.Pi) % 90.0)
}

func (p *Particle) MaximumHeight() {
	p.MaxHeight = (math.Pow((p.InitialVelocity*math.Sin(p.Theta)), 2))/(2*p.G) + p.InitialHeight
}

func (p *Particle) TimeOfFlight() {
	a := math.Pow((p.InitialVelocity * math.Sin(p.Theta)), 2)
	b := 2 * p.G * p.InitialHeight
	c := a + b
	p.FlightTime = (math.Sqrt(c) + (p.InitialVelocity * math.Sin(p.Theta))) / p.G
}

func (p *Particle) MaxRange() {
	if p.FlightTime == 0 {
		p.TimeOfFlight()
	}
	p.HorizontalRange = (p.InitialVelocity * math.Cos(p.Theta)) * p.FlightTime
}

func (p Particle) Position(x float64) float64 {
	y := (x * math.Tan(p.Theta)) - ((p.G * math.Pow(x, 2)) / (2 * math.Pow(p.InitialVelocity, 2) * math.Pow((math.Cos(p.Theta)), 2))) + p.InitialHeight
	return y
}

func (p Particle) Path() ([]float64, []float64) {
	var xs []float64
	var ys []float64
	for x := 0.0; x < p.HorizontalRange; x += 0.1 {
		xs = append(xs, x)
		ys = append(ys, p.Position(x))
	}
	return xs, ys
}

func (p Particle) PathPlot() {

	pts := make(plotter.XYs, int(p.HorizontalRange)+2)
	for i := 0; float64(i) < p.HorizontalRange; i++ {
		pts[i].X = float64(i)
		pts[i].Y = p.Position(float64(i))
	}
	pts[int(p.HorizontalRange)+1].X = p.HorizontalRange
	pts[int(p.HorizontalRange)+1].Y = 0.0
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
	if p.MaxHeight > p.HorizontalRange {
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

func (p Particle) XVelocityDrag(t int) float64 {
	a := ((p.InitialVelocity) * math.Cos(p.Theta)) * math.Pow(math.E, (-1*(p.G*float64(t))/p.TerminalVelocity))
	return (a)
}

func (p Particle) XPositionDrag(t int) float64 {
	a := ((p.InitialVelocity * p.TerminalVelocity * math.Cos(p.Theta)) / p.G) * (1 - math.Pow(math.E, ((-1*(p.G*float64(t)))/p.TerminalVelocity)))
	return a
}

func (p Particle) YVelocityDrag(t int) float64 {
	a := (p.InitialVelocity*math.Sin(p.Theta))*math.Pow(math.E, ((-1*(p.G*float64(t)))/p.TerminalVelocity)) - p.TerminalVelocity*(1-math.Pow(math.E, ((-1*(p.G*float64(t)))/p.TerminalVelocity)))
	return a
}

func (p Particle) YPositionDrag(t int) float64 {
	a := (p.TerminalVelocity/p.G)*(p.InitialVelocity*math.Sin(p.Theta)+p.TerminalVelocity)*(1-math.Pow(math.E, ((-1*(p.G*float64(t)))/p.TerminalVelocity))) - (float64(t) * p.TerminalVelocity)
	return a
}

func (p Particle) PathPlotDrag() {

	pts := make(plotter.XYs, int(p.HorizontalRange)+2)
	for i := 0; float64(i) < p.HorizontalRange; i++ {
		pts[i].X = p.XPositionDrag(i)
		pts[i].Y = p.YPositionDrag(i)
		if p.YPositionDrag(i) < 0 {
			break
		}
	}
	//pts[int(p.HorizontalRange)+1].X = p.HorizontalRange
	//pts[int(p.HorizontalRange)+1].Y = 0.0
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
	if p.MaxHeight > p.HorizontalRange {
		pl.X.Max = pl.Y.Max
		fmt.Println(pl.X.Max, pl.Y.Max)
	} else {
		pl.Y.Max = pl.X.Max
		fmt.Println(pl.X.Max, pl.Y.Max)
	}
	if err := pl.Save(4*vg.Inch, 4*vg.Inch, "pointsdrag.png"); err != nil {
		panic(err)
	}
}

func main() {
	tps := Particle{InitialHeight: 10.0, Theta: 45.0, G: 9.8, InitialVelocity: 50.0, TerminalVelocity: 100.0}
	tps.ThetaDegrees()
	tps.MaximumHeight()
	tps.TimeOfFlight()
	tps.MaxRange()
	tps.PathPlotDrag()
	fmt.Println(tps)
}
