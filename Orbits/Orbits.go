package main

import (
	"fmt"
	"math"
)

const Gravitation = 6.673e-11

type Satellite struct { //Each satellite variable represents one particle
	Mass, InitialVelocity, radius, centripetalforce float64
	Coords                                          Vector
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

func (ps *System) CentreForce() {
	for i := 0; i < len(ps.Satellites); i++ {
		planet := ps.Satellites[i]
		planet.centripetalforce = ((Gravitation * ps.TotalMass * planet.Mass) / math.Pow(planet.radius, 2))
	}
}

func main() {
	sat1 := Satellite{Mass: 378, Coords: Vector{X: 240, Y: 10}}
	sat2 := Satellite{Mass: 1, Coords: Vector{X: -110, Y: -10}}
	sys := System{Satellites: []Satellite{sat1, sat2}}
	sys.CentreOfMass()
	fmt.Println(sys.MassCentre, sys.TotalMass)
}
