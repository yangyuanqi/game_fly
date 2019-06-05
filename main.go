package main

import (
	"github.com/hajimehoshi/ebiten"
	"test/play/plays"
)

func main() {
	game := plays.NewGame()
	game.Onload()
	ebiten.Run(game.Update, 360, 640, 1, "灰机大战")
}
