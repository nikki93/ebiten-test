package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"

	"github.com/nikki93/ebiten-test/game"
)

func main() {
	ebiten.SetWindowSize(game.W, game.H)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
