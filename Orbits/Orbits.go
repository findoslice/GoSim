package main

import (
	"fmt"
	"math"
	"time"
)

const G = 6.673e-11

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
				force := (G*el.Mass*obj.Mass)/  (math.Pow((obj.Coords.X-el.Coords.X),2)+math.Pow((obj.Coords.Y-el.Coords.Y),2))
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

func (ps *System) Update(){
	ps.GravitationalForces()
	ps.Accelerate()
	ps.Move()
	fmt.Println(ps.Satellites[1].Coords,ps.Satellites[1].Velocity, ps.Satellites[1].Force)
}

func main() {
	sat1 := Satellite{Mass: 9e+20, Coords: Vector{X: 0, Y: 0}, Velocity: Vector{X:0, Y:0}}
	sat2 := Satellite{Mass: 10, Coords: Vector{X: 1, Y: 0}, Velocity: Vector{X:100,Y:0}}
	sys := System{Satellites: []Satellite{sat1, sat2}}
	i :=0
	for (i<360000){
		sys.Update()
		time.Sleep(50*time.Millisecond)
		i++
	}
	//fmt.Println(sys.Satellites[1].Coords)
}
