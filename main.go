package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

var white *ebiten.Image

func init() {
	white, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	white.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(20, 20)
	op.GeoM.Translate(100, 100)
	screen.DrawImage(white, op)
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
