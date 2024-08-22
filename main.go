package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	//temporary
	isinit bool
	//----
	timer    int
	particle Particle
	width    int
	height   int
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
	g.particle.x_pos = float32(g.width) / 2.0
	g.particle.y_pos = float32(g.height) / 2.0
	g.particle.radius = 50
	g.particle.color = color.RGBA{0, 255, 0, 0}
	g.particle.anti_aliasing = true
	g.isinit = true
	g.particle.direction = [2]float32{1, 2}
}

func (g *Game) bounce() {
	if g.particle.x_pos+g.particle.radius >= float32(g.width) || g.particle.x_pos-g.particle.radius <= 0 {
		g.particle.direction[0] = -1 * g.particle.direction[0]
	}
	if g.particle.y_pos+g.particle.radius >= float32(g.height) || g.particle.y_pos-g.particle.radius <= 0 {
		g.particle.direction[1] = -1 * g.particle.direction[1]
	}
}

func (g *Game) Update() error {
	g.timer++
	//g.particle.direction = [2]float32{1, 2}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	g.bounce()
	if g.timer%1 == 0 {
		i, j := g.particle.direction[0], g.particle.direction[1]
		//fmt.Println(i, j)
		g.particle.x_pos += i * 20
		g.particle.y_pos += j * 20
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.isinit {
		g.gameinitconfig()
	}
	vector.DrawFilledCircle(screen, g.particle.x_pos, g.particle.y_pos, g.particle.radius, g.particle.color, g.particle.anti_aliasing)
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
