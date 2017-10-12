package main

import (
	"fmt"
	"math"
)

type Particle struct {
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

func main() {
	yerMaw := Particle{initialVelocity: 10.0, theta: 45.0, g: 9.8}
	yerMaw.thetaDegrees()
	yerMaw.maximumHeight()
	yerMaw.timeOfFlight()
	yerMaw.maxRange()
	fmt.Println(yerMaw.maxHeight, yerMaw.flightTime, yerMaw.horizontalRange)
}
