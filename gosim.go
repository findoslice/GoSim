package GoSim

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//
//TODO: Add a Variable expression for G
//

type Particle struct {
	//group particle variables together
	InitialHeight, InitialVelocity, FinalHeight, FinalVelocity, MaxHeight, HorizontalRange, FlightTime, Theta, G float64
}

type VecParticle struct {
	Xpos, Ypos, Xvel, Yvel, Xacc, Yacc float64
}

func (p *Particle) SetDefaults() {
	if p.G == 0.0 {
		p.G = 9.8
	}
	p.ThetaDegrees()
}

func (p *Particle) ThetaDegrees() {
	p.Theta = (p.Theta / 360) * 2 * math.Pi
	if math.Abs(p.Theta) > 90.0 {
		p.Theta = float64(int(p.Theta) % 90)
	}
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
	if p.FlightTime == 0.0 {
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
	n := 1.0
	if p.HorizontalRange < 1 {
		n = 100.0
	} else if p.HorizontalRange < 2 {
		n = 40.0
	} else if p.HorizontalRange < 5 {
		n = 10.0
	} else if p.HorizontalRange < 10 {
		n = 5.0
	} else if p.HorizontalRange < 20 {
		n = 3.0
	} else if p.HorizontalRange > 10000 {
		n = 0.1
	} else if p.HorizontalRange > 100000 {
		n = 0.00000001
	} else {
		n = 1.0
	}

	pts := make(plotter.XYs, int(n*p.HorizontalRange)+2)
	for i := 0; float64(i) < n*p.HorizontalRange; i++ {
		pts[i].X = float64(i) / float64(n)
		pts[i].Y = p.Position(float64(i) / float64(n))
	}
	pts[int(n*p.HorizontalRange)+1].X = p.HorizontalRange
	pts[int(n*p.HorizontalRange)+1].Y = 0.0
	pl, err := plot.New()
	if err != nil {
		panic(err)
	}
	pl.Title.Text = "Trajectory of Kristel's Chalk"
	pl.X.Label.Text = "Horizontal Displacement"
	pl.Y.Label.Text = "Height"

	err = plotutil.AddLinePoints(pl, "Chalk", pts)
	if err != nil {
		panic(err)
	}
	pl.X.Min = 0
	pl.Y.Min = 0
	if p.MaxHeight > p.HorizontalRange {
		pl.X.Max = pl.Y.Max
	} else {
		pl.Y.Max = pl.X.Max
	}
	pl.Add(plotter.NewGrid())
	if err := pl.Save(8*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
