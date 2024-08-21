package main

import (
	"image/color"
	"log"

	//"math/rand"
	//"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	timer     int
	particle  Particle
	direction int
}

type Particle struct {
	x_pos         float32
	y_pos         float32
	radius        float32
	color         color.Color
	anti_aliasing bool
}

func (g *Game) Update() error {
	g.timer++
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	width, _ := ebiten.WindowSize()
	if g.particle.x_pos+g.particle.radius >= float32(width) {
		g.direction = -1
	}
	if g.particle.x_pos-g.particle.radius <= 0 {
		g.direction = 0
	}
	if g.timer%1 == 0 {
		if g.direction == 0 {
			g.particle.x_pos += 10
		}
		if g.direction < 0 {
			g.particle.x_pos -= 10
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.particle.x_pos = 0
	g.particle.y_pos = float32(screen.Bounds().Size().Y) / 2.0
	g.particle.radius = 50
	g.particle.color = color.RGBA{0, 255, 0, 0}
	g.particle.anti_aliasing = true
	vector.DrawFilledCircle(screen, g.particle.x_pos, g.particle.y_pos, g.particle.radius, g.particle.color, g.particle.anti_aliasing)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(2000, 1000)
	ebiten.SetWindowTitle("moving rectangle")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
