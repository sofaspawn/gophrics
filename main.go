package main

import (
	"image/color"
	"log"

	//"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	isinit    bool
	timer     int
	particles []*Particle
	width     int
	height    int
	g         float32
}

type Particle struct {
	x_pos         float32
	y_pos         float32
	radius        float32
	mass          float32
	velocity      float64
	color         color.Color
	anti_aliasing bool
	direction     [2]float32
	potential     float32
	total         float32
	//kinetic       float32
}

func (g *Game) gameinitconfig() {
	g.g = 10
	g.width, g.height = ebiten.WindowSize()
	random_x := float32(g.width) * rand.Float32()
	random_y := float32(g.height) * rand.Float32()
	var radius float32 = 100
	var velocity float64 = 1
	var mass float32 = 5

	g.addparticle(random_x, random_y, radius, mass, [4]uint8{0, 255, 0, 0}, true, [2]float32{10 * rand.Float32(), 10 * rand.Float32()}, velocity)
	g.addparticle(random_x, random_y, radius, mass, [4]uint8{0, 0, 255, 0}, true, [2]float32{10 * rand.Float32(), 10 * rand.Float32()}, velocity)
	g.addparticle(random_x, random_y, radius, mass, [4]uint8{255, 0, 0, 0}, true, [2]float32{10 * rand.Float32(), 10 * rand.Float32()}, velocity)

	g.isinit = true
}

func (g *Game) addparticle(x_pos, y_pos, radius, mass float32, c [4]uint8, anti_aliasing bool, direction [2]float32, velocity float64) {
	particle := Particle{}

	particle.x_pos = x_pos
	particle.y_pos = y_pos
	particle.radius = radius
	particle.mass = mass
	particle.color = color.RGBA{c[0], c[1], c[2], c[3]}
	particle.anti_aliasing = anti_aliasing
	particle.direction = direction
	particle.velocity = velocity

	g.particles = append(g.particles, &particle)
}

// FIXME: fix gravity
/*
func (g *Game) gravity(particle *Particle) {
	particle.direction[1] = particle.direction[1] - 0.1
}

// FIXME: fix conservation of energy
func (g *Game) consvOfEnergy(particle *Particle) {
	particle.potential = (particle.mass) * (g.g) * (float32(g.height) - particle.y_pos)
	particle.total = particle.mass * (g.g) * float32(g.height)
	particle.velocity = math.Sqrt(float64(2/particle.mass) * float64(particle.total-particle.potential))
}
*/

func (g *Game) change_radius(particle *Particle) {
	if g.timer%1 == 0 {
	}
}

func bounce(particle *Particle) {
	width, height := ebiten.WindowSize()
	if particle.x_pos+particle.radius >= float32(width) || particle.x_pos-particle.radius <= 0 {
		particle.direction[0] *= -1
	}
	if particle.y_pos+particle.radius >= float32(height) || particle.y_pos-particle.radius <= 0 {
		particle.direction[1] *= -1
	}
}

func (g *Game) Update() error {
	g.timer++
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	for _, particle := range g.particles {
		//g.gravity(particle)
		//g.consvOfEnergy(particle)
		bounce(particle)
		if g.timer%1 == 0 {
			i, j := particle.direction[0], particle.direction[1]

			particle.x_pos += i * float32(particle.velocity)
			particle.y_pos += j * float32(particle.velocity)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.isinit {
		g.particles = []*Particle{}
		g.gameinitconfig()
	}
	if len(g.particles) > 0 && g.particles[0] != nil {
		for _, particle := range g.particles {
			vector.DrawFilledCircle(screen, particle.x_pos, particle.y_pos, particle.radius, particle.color, particle.anti_aliasing)
		}
	}
	//fmt.Println(g.particles[0])
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("bouncing ball")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
