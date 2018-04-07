package main

import (
	"fmt"
	"math"
	//"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const G = 6.674e-11

type Satellite struct { //Each satellite variable represents one particle
	Mass, radius float64
	Coords, Force, Velocity Vector
}

type System struct { //only one System struct should exist as this contains information about entire system
	TotalMass  float64
	MassCentre Vector
	Satellites []Satellite
}

type Vector struct {
	X, Y float64
}

func (ps *System) Initialise() { //calculates system data at the start of each iteration
	ps.CentreOfMass()
}

func (ps *System) CentreOfMass() {
	xmass := 0.0
	ymass := 0.0
	for i := 0; i < len(ps.Satellites); i++ {
		ps.TotalMass = ps.TotalMass + ps.Satellites[i].Mass
		xmass = xmass + ps.Satellites[i].Mass*ps.Satellites[i].Coords.X
		ymass = ymass + ps.Satellites[i].Mass*ps.Satellites[i].Coords.Y
	}
	ps.MassCentre = Vector{X: xmass / ps.TotalMass, Y: ymass / ps.TotalMass}
}

func (vec Vector) Magnitude() float64{
	return math.Sqrt(math.Pow(vec.X,2)+math.Pow(vec.Y,2))
}

func RelUnitVector(vec1 Vector, vec2 Vector) Vector{
	rel := Vector{X: vec2.X-vec1.X, Y: vec2.Y - vec1.Y}
	mag := rel.Magnitude()
	return Vector{X:rel.X/mag, Y:rel.Y/mag}
}


func (ps *System) GravitationalForces(){
	for i, el := range ps.Satellites{
		el.Force = Vector{X:0, Y:0}
		for j,obj := range ps.Satellites{
			if (i != j){
				force := (G*el.Mass*obj.Mass)/(math.Pow((obj.Coords.X-el.Coords.X),2)+math.Pow((obj.Coords.Y-el.Coords.Y),2))
				rel := RelUnitVector(el.Coords,obj.Coords)
				el.Force.X += force*rel.X
				el.Force.Y += force*rel.Y
				ps.Satellites[i] = el
			}
		}
	}
}

func (ps *System) Accelerate(){
	for i,el := range ps.Satellites{
		el.Velocity.X += el.Force.X/el.Mass
		el.Velocity.Y += el.Force.Y/el.Mass
		ps.Satellites[i] = el
	}
}

func (ps *System) Move(){
	for i,el := range ps.Satellites{
		el.Coords.X += el.Velocity.X
		el.Coords.Y += el.Velocity.Y
		ps.Satellites[i] = el
	}
}

func (ps System) TotalMomentum() Vector{
	momentum := Vector{X:0,Y:0}
	for _,el := range ps.Satellites{
		momentum.X += el.Velocity.X*el.Mass
		momentum.Y += el.Velocity.Y*el.Mass
	}
	return momentum
}

func (ps *System) Update(){
	ps.GravitationalForces()
	ps.Accelerate()
	ps.Move()
	fmt.Println(ps.TotalMomentum(), ps.Satellites[1].Velocity, ps.Satellites[1].Force)
}

func (ps System) Plot() {
	//n := 1.0
	/*if p.HorizontalRange < 1 {
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
	}*/

	pts := make(plotter.XYs, 500000)
	pts2 := make(plotter.XYs, 500000)
	for i := 0; float64(i) < 500000; i++ {
		ps.Update()
		pts[i].X = ps.Satellites[1].Coords.X
		pts[i].Y = ps.Satellites[1].Coords.Y

		pts2[i].X = ps.Satellites[0].Coords.X
		pts2[i].Y = ps.Satellites[0].Coords.Y
	}
	pl, err := plot.New()
	if err != nil {
		panic(err)
	}
	pl.Title.Text = "ISS trajectory"
	pl.X.Label.Text = "X"
	pl.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(pl, "ISS", pts)
	if err != nil {
		panic(err)
	}
	err = plotutil.AddLinePoints(pl, "Earth", pts2)
	if err != nil {
		panic(err)
	}
	pl.X.Min = -10e+3
	pl.Y.Min = -10e+3
	pl.X.Max = 10e+3
	pl.Y.Max = 10e+3
	/*if p.MaxHeight > p.HorizontalRange {
		pl.X.Max = pl.Y.Max
	} else {
		pl.Y.Max = pl.X.Max
	}*/
	pl.Add(plotter.NewGrid())
	if err := pl.Save(50*vg.Inch, 50*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func main() {
	sat1 := Satellite{Mass: 417289000000000000000, Coords: Vector{X: 0, Y: 0}, Velocity: Vector{X:0, Y:0}}
	sat2 := Satellite{Mass: 417289000000000000000, Coords: Vector{X: 6e3, Y: 0}, Velocity: Vector{X:0,Y:30}}
	sys := System{Satellites: []Satellite{sat1, sat2}}
	sys.Plot()

}
