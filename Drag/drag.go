package Drag

import (
	//"fmt"
	"GoSim"
	"math"
	//"gonum.org/v1/plot"
	//"gonum.org/v1/plot/plotter"
	//"gonum.org/v1/plot/plotutil"
	//"gonum.org/v1/plot/vg"
	"math"
)

type Drag struct {
	Particle                                             GoSim.Particle
	TerminalVelocity, xpos, xvel, xacc, ypos, yvel, yacc float64 //these variables are unexported as they should not be edited by the user
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
	p.xvel = p.Particle.InitialHeight * math.Cos(p.Particle.Theta) * math.Pow(math.E, (-1*p.Particle.G*t)/p.TerminalVelocity)
}

func (p *Drag) YVelocity() {
	continue
}

func (p *Drag) XAcceleration() {
	p.xacc = -1 * p.Particle.G * (p.xvel / p.TerminalVelocity)
}

func (p *Drag) YAcceleration() {
	p.yacc = -1 * p.Particle.G * (1 - (p.xvel / p.TerminalVelocity))
}
