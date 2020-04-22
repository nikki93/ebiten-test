package game

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const W = 947
const H = 781

const N = 1000

type rect struct {
	x, y    float64
	r, g, b float64
	speed   float64
	w, h    float64
	phase   float64
}

type Game struct {
	loaded bool

	white *ebiten.Image

	rects [N]*rect
}

var startTime time.Time
var lastFrameTime time.Time

func (g *Game) Load() {
	startTime = time.Now()
	lastFrameTime = startTime

	g.white, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	g.white.Fill(color.White)

	for i := 0; i < N; i++ {
		g.rects[i] = &rect{
			x:     W * rand.Float64(),
			y:     H * rand.Float64(),
			r:     rand.Float64(),
			g:     rand.Float64(),
			b:     rand.Float64(),
			speed: 0.5 * H * rand.Float64(),
			w:     W * rand.Float64() / 16.0,
			h:     W * rand.Float64() / 16.0,
			phase: 2 * math.Pi * rand.Float64(),
		}
	}

	g.loaded = true
}

func (g *Game) Update(screen *ebiten.Image) error {
	if !g.loaded {
		g.Load()
	}

	now := time.Now()
	t := now.Sub(startTime).Seconds()
	dt := now.Sub(lastFrameTime).Seconds()

	for i := 0; i < N; i++ {
		rect := g.rects[i]
		rect.x += rect.speed * math.Sin(t+rect.phase) * dt
	}

	lastFrameTime = now
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < N; i++ {
		rect := g.rects[i]
		op.GeoM.Reset()
		op.ColorM.Reset()
		op.GeoM.Scale(rect.w, rect.h)
		op.GeoM.Translate(rect.x, rect.y)
		op.ColorM.Scale(rect.r, rect.g, rect.b, 1)
		screen.DrawImage(g.white, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("fps: %0.2f", ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
}
