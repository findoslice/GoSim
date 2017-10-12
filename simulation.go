package main

import (
	"fmt"
	"math"
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

func main() {
	yerMaw := Particle{initialVelocity: 10.0, theta: 45.0, g: 9.8}
	yerMaw.thetaDegrees()
	yerMaw.maximumHeight()
	yerMaw.timeOfFlight()
	yerMaw.maxRange()
	x, y := yerMaw.path()
	fmt.Println(x)
	fmt.Println("\n\n\n\n")
	fmt.Println(y)
}
