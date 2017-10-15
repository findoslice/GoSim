package Drag

import (
	//"fmt"
	"GoSim"
	"math"
	//"gonum.org/v1/plot"
	//"gonum.org/v1/plot/plotter"
	//"gonum.org/v1/plot/plotutil"
	//"gonum.org/v1/plot/vg"
)

type Drag struct {
	Particle         GoSim.Particle
	vector           GoSim.VecParticle //this vector is unexported as it should not be referenced by the user
	TerminalVelocity float64
}

func (p *Drag) SetDefaults() {
	if p.Particle.G == 0 {
		p.Particle.G = 9.8
	}
	if p.TerminalVelocity == 0 {
		p.TerminalVelocity = 100.0
	}
}

func (p *Drag) XPosition() {
	continue
}

func (p *Drag) YPosition() {
	continue
}

func (p *Drag) XVelocity(t float64) {
	p.vector.Xvel = p.Particle.InitialHeight * math.Cos(p.Particle.Theta) * math.Pow(math.E, (-1*p.Particle.G*t)/p.TerminalVelocity)
}

func (p *Drag) YVelocity() {
	continue
}

func (p *Drag) XAcceleration() {
	p.vector.Xacc = -1 * p.Particle.G * (p.vector.Xvel / p.TerminalVelocity)
}

func (p *Drag) YAcceleration() {
	p.vector.Yacc = -1 * p.Particle.G * (1 - (p.vector.Yvel / p.TerminalVelocity))
}
