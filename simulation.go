package main

import (
	"fmt"
	"math"
)

type Particle struct {
	initialHeight, initialVelocity, finalHeight, finalVelocity, maxHeight, horizontalRange, flightTime, theta float64
}

func (p *Particle) max_height() {
	g := 9.8
	p.theta = (p.theta / 360) * 2 * math.Pi
	p.maxHeight = ((p.initialVelocity*math.Sin(p.theta))*(p.initialVelocity*math.Sin(p.theta)))/(2*g) + p.initialHeight
}

func main() {
	yerMaw := Particle{initialVelocity: 10.0, theta: 45.0, initialHeight: 0.0}
	yerMaw.max_height()
	fmt.Println(yerMaw.maxHeight)
}
