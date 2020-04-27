package main

import (
	"github.com/hajimehoshi/ebiten/mobile"

	"github.com/nikki93/ebiten-test/game"
)

func init() {
	mobile.SetGame(&game.Game{})
}

func Dummy() {}
