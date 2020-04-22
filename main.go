package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var white *ebiten.Image

const N = 30000

type rect struct {
	x, y    float64
	r, g, b float64
	speed   float64
	w, h    float64
	phase   float64
}

var rects = [N]*rect{}

type game struct{}

var startTime time.Time
var lastFrameTime time.Time

func (g *game) Load() {
	startTime = time.Now()
	lastFrameTime = startTime

	white, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	white.Fill(color.White)

	for i := 0; i < N; i++ {
		rects[i] = &rect{
			x:     320 * rand.Float64(),
			y:     240 * rand.Float64(),
			r:     rand.Float64(),
			g:     rand.Float64(),
			b:     rand.Float64(),
			speed: 0.5 * 240 * rand.Float64(),
			w:     320 * rand.Float64() / 16.0,
			h:     320 * rand.Float64() / 16.0,
			phase: 2 * math.Pi * rand.Float64(),
		}
	}
}

func (g *game) Update(screen *ebiten.Image) error {
	now := time.Now()
	t := now.Sub(startTime).Seconds()
	dt := now.Sub(lastFrameTime).Seconds()

	for i := 0; i < N; i++ {
		rect := rects[i]
		rect.x += rect.speed * math.Sin(t+rect.phase) * dt
	}

	lastFrameTime = now
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < N; i++ {
		rect := rects[i]
		op.GeoM.Reset()
		op.ColorM.Reset()
		op.GeoM.Scale(rect.w, rect.h)
		op.GeoM.Translate(rect.x, rect.y)
		op.ColorM.Scale(rect.r, rect.g, rect.b, 1)
		screen.DrawImage(white, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %0.2f", ebiten.CurrentFPS()))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := &game{}
	g.Load()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
