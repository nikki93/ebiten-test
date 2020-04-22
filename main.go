package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

var white *ebiten.Image

func init() {
	white, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	white.Fill(color.White)
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(20, 20)
	op.GeoM.Translate(100, 100)
	op.ColorM.Scale(1, 0, 0, 1)
	screen.DrawImage(white, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %0.2f", ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
