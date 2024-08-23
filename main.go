package main

import (
	"image/color"
	"log"
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
}

type Particle struct {
	x_pos         float32
	y_pos         float32
	radius        float32
	color         color.Color
	anti_aliasing bool
	direction     [2]float32
}

func (g *Game) gameinitconfig() {
	g.width, g.height = ebiten.WindowSize()
    random_x := float32(g.width) * rand.Float32()
    random_y := float32(g.height) * rand.Float32()
    g.addparticle(float32(g.width)/2.0, float32(g.width)/2.0, 50, [4]uint8{0, 255, 0, 0}, true, [2]float32{10*rand.Float32(),10*rand.Float32()})
    g.addparticle(random_x, random_y, 50, [4]uint8{0, 0, 255, 0}, true, [2]float32{10*rand.Float32(),10*rand.Float32()})
    g.addparticle(random_x, random_y, 50, [4]uint8{255, 0, 0, 0}, true, [2]float32{10*rand.Float32(),10*rand.Float32()})
	g.isinit = true
}

func (g *Game) addparticle(x_pos, y_pos, radius float32, c [4]uint8, anti_aliasing bool, direction [2]float32){
    particle := Particle{}
    particle.x_pos = x_pos
    particle.y_pos = y_pos
    particle.radius = radius
    particle.color = color.RGBA{c[0], c[1], c[2], c[3]}
    particle.anti_aliasing = anti_aliasing
    particle.direction = direction
    g.particles = append(g.particles, &particle)
}

func bounce(particle *Particle) {
	width, height := ebiten.WindowSize()
	if particle.x_pos+particle.radius >= float32(width) || particle.x_pos-particle.radius <= 0 {
		particle.direction[0] = -1 * particle.direction[0]
	}
	if particle.y_pos+particle.radius >= float32(height) || particle.y_pos-particle.radius <= 0 {
		particle.direction[1] = -1 * particle.direction[1]
	}
}

func (g *Game) Update() error {
	g.timer++
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

    for _, particle := range g.particles{
        bounce(particle)
        if g.timer%1 == 0 {
            i, j := particle.direction[0], particle.direction[1]
            inc_p0 := 10
            particle.x_pos += i * float32(inc_p0)
            particle.y_pos += j * float32(inc_p0)
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
        for _, particle := range g.particles{
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
