package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const G = 6.674e-11

type Satellite struct { //Each satellite variable represents one particle
	Mass, radius float64
	Coords, Force, Velocity Vector
	Name string 
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
	points := make([]plotter.XYs, len(ps.Satellites))
	for i := range points{
		points[i] = make(plotter.XYs, 500000)
	}

	pl, err := plot.New()
	if err != nil {
		panic(err)
	}

	pl.X.Min = 0
	pl.Y.Min = 0
	pl.X.Max = 0
	pl.Y.Max = 0

	for i := 0; float64(i) < 500000; i++ {
		ps.Update()

		for j,el := range ps.Satellites{
			points[j][i].X = el.Coords.X
			points[j][i].Y = el.Coords.Y
			if (points[j][i].X > pl.X.Max){
				pl.X.Max = 1.2*points[j][i].X
			}
			if (points[j][i].X < pl.X.Min){
				pl.X.Min = 1.2*points[j][i].X
			}
			if (points[j][i].Y > pl.Y.Max){
				pl.Y.Max = 1.2*points[j][i].Y
			}
			if (points[j][i].Y < pl.Y.Min){
				pl.Y.Min = 1.2*points[j][i].Y
			}
		}
	}
	pl.Title.Text = "Satellites and shit"
	pl.X.Label.Text = "X"
	pl.Y.Label.Text = "Y"

	for i,el := range ps.Satellites{
		err = plotutil.AddLinePoints(pl, el.Name, points[i])
		if err != nil {
			panic(err)
		}
	}
	pl.Add(plotter.NewGrid())
	if err := pl.Save(20*vg.Inch, 20*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func main() {
	sat1 := Satellite{Mass: 1.9891e+30, Coords: Vector{X: -149.6e+9, Y: 0}, Velocity: Vector{X:0, Y:0}, Name: "Sun"} //sun
	sat2 := Satellite{Mass: 5.972e+24, Coords: Vector{X: 0, Y: 0}, Velocity: Vector{X:0,Y:30000}, Name: "Earth"} //earth
	sat3 := Satellite{Mass:7.3476731e+22, Coords: Vector{X:384403000, Y:0}, Velocity: Vector{X:0,Y:31000}, Name: "Moon"} //moon
	sys := System{Satellites: []Satellite{sat1, sat2, sat3}}
	sys.Plot()

}
