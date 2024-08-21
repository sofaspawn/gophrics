package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{
    timer int
    particle Particle
}

type Particle struct{
    x_pos float32
    y_pos float32
    inc float32
}

func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyEscape){
        return ebiten.Termination
    }
    g.timer++
    g.particle.inc++;
    width, height := ebiten.WindowSize()
    g.particle.y_pos = float32(rand.Intn(height))
    if g.particle.x_pos>=float32(width){
        g.particle.x_pos=0
    }
    if g.timer%1==0{
        g.particle.x_pos+=g.particle.inc
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image){
    //g.particle.y_pos = float32(screen.Bounds().Size().Y)/2.0
    vector.DrawFilledCircle(screen, g.particle.x_pos, g.particle.y_pos, 20, color.RGBA{0,255,0,0}, false)
    vector.DrawFilledCircle(screen, g.particle.x_pos, float32(rand.Intn(screen.Bounds().Size().Y)), 20, color.RGBA{255,0,0,0}, false)
    vector.DrawFilledCircle(screen, g.particle.x_pos, float32(rand.Intn(screen.Bounds().Size().Y)), 20, color.RGBA{0,0,255,0}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int){
    return outsideWidth, outsideHeight
}

func main() {
    rand.Seed(time.Now().UnixNano())
    ebiten.SetWindowSize(2000, 1000)
    ebiten.SetWindowTitle("moving rectangle")
    if err := ebiten.RunGame(&Game{}); err!=nil{
        log.Fatal(err)
    }
}
